// Creates action files for Go code
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/goccy/go-yaml"
	"github.com/posener/goaction"
	"github.com/posener/goaction/actionutil"
	"github.com/posener/goaction/internal/metadata"
	"github.com/posener/goaction/log"
	"github.com/posener/script"
)

var (
	//goaction:required
	path    = flag.String("path", "", "Path to main Go main package.")
	name    = flag.String("name", "", "Override action name, the default name is the package name.")
	desc    = flag.String("desc", "", "Override action description, the default description is the package synopsis.")
	image   = flag.String("image", "golang:1.14.2-alpine3.11", "Override Docker image to run the action with (See https://hub.docker.com/_/golang?tab=tags).")
	install = flag.String("install", "", "Comma separated list of requirements to 'apk add'.")
	icon    = flag.String("icon", "", "Set branding icon. (See options at https://feathericons.com).")
	color   = flag.String("color", "", "Set branding color. (white, yellow, blue, green, orange, red, purple or gray-dark).")

	email       = goaction.Getenv("email", "posener@gmail.com", "Email for commit message.")
	githubToken = goaction.Getenv("github-token", "", "Github token for PR comments. Optional.")
)

const (
	action      = "action.yml"
	dockerfile  = "Dockerfile"
	autoComment = "# File generated by github.com/posener/goaction. DO NOT EDIT.\n\n"
)

func main() {
	flag.Parse()

	// Load go code.
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, *path, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// Get main package
	var mainPkg *ast.Package
	for name, pkg := range pkgs {
		if name == "main" {
			mainPkg = pkg
			break
		}
	}
	if mainPkg == nil {
		log.Fatalf("No main package in path %q", *path)
	}

	// Parse Go code to Github actions metadata.
	m, err := metadata.New(mainPkg)
	if err != nil {
		// For parsing error, log the file location.
		var pe metadata.ErrParse
		if errors.As(err, &pe) {
			log.FatalFile(fset.Position(pe.Pos), err)
		}
		log.Fatal(err)
	}
	if *name != "" {
		m.Name = *name
	}
	if *desc != "" {
		m.Desc = *desc
	}

	m.Branding.Icon = *icon
	m.Branding.Color = *color

	// Applying changes.

	// Create action file.
	log.Printf("Writing %s\n", action)
	err = script.Writer("yml", func(w io.Writer) error {
		w.Write([]byte(autoComment))
		return yaml.NewEncoder(w).Encode(m)
	}).ToFile(action)
	if err != nil {
		log.Fatal(err)
	}

	// Create dockerfile
	log.Printf("Writing %s\n", dockerfile)
	dir, err := pathRelDir(*path)
	data := tmplData{
		Dir:     dir,
		Image:   *image,
		Install: strings.ReplaceAll(*install, ",", " "),
	}
	err = script.Writer("template", func(w io.Writer) error {
		w.Write([]byte(autoComment))
		return tmpl.Execute(w, data)
	}).ToFile(dockerfile)
	if err != nil {
		log.Fatal(err)
	}

	diff := gitDiff()

	if diff != "" {
		log.Printf("Applied changes:\n\n%s\n\n", diff)
	} else {
		log.Printf("No changes were made.")
	}

	if !goaction.CI {
		return
	}

	err = actionutil.GitConfig("goaction", email)
	if err != nil {
		log.Fatal(err)
	}

	switch goaction.Event {
	case goaction.EventPush:
		log.Printf("Push mode.")
		if diff == "" {
			log.Printf("Skipping commit stage.")
			break
		}
		push()
	case goaction.EventPullRequest:
		log.Printf("Pull request mode.")
		pr(diff)
	default:
		log.Printf("Unsupported action mode: %s", goaction.Event)
	}
}

func gitDiff() string {
	var diff strings.Builder
	for _, path := range []string{action, dockerfile} {
		// Add files to git, in case it does not exists
		d, err := actionutil.GitDiff(path)
		if err != nil {
			log.Fatal(err)
		}
		if d != "" {
			diff.WriteString(fmt.Sprintf("Path `%s`:\n\n", path))
			diff.WriteString(fmt.Sprintf("```diff\n%s\n```\n\n", d))
		}
	}
	return diff.String()
}

// Commit and push chnages to upstream branch.
func push() {
	err := actionutil.GitCommitPush(
		[]string{action, dockerfile},
		"Update action files")
	if err != nil {
		log.Fatal(err)
	}
}

// Post a pull request comment with the expected diff.
func pr(diff string) {
	if githubToken == "" {
		log.Printf("In order to add request comment, set the GITHUB_TOKEN input.")
		return
	}

	body := "[Goaction](https://github.com/posener/goaction) will apply the following changes after PR is merged.\n\n" + diff
	if diff == "" {
		body = "[Goaction](https://github.com/posener/goaction) detected no required changes to Github action files."
	}

	ctx := context.Background()
	err := actionutil.PRComment(ctx, githubToken, body)
	if err != nil {
		log.Fatal(err)
	}
}

// pathRelDir returns the containing directory of a given path in a relative form, relative to the
// working directory prefixed with "./"
func pathRelDir(path string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path, err = filepath.Rel(wd, path)
	if err != nil {
		return "", err
	}
	// If Rel returned ".", fix it to empty string which will eventually mutate to "./".
	if path == "." {
		path = ""
	}
	// Add a "./" prefix.
	if !strings.HasPrefix(path, "./") {
		path = "./" + path
	}
	return path, nil
}

type tmplData struct {
	Dir     string
	Image   string
	Install string
}

var tmpl = template.Must(template.New("dockerfile").Parse(`
FROM {{ .Image }}
RUN apk add git {{ .Install }}

COPY . /home/src
WORKDIR /home/src
RUN go build -o /bin/action {{ .Dir }}

ENTRYPOINT [ "/bin/action" ]
`))
