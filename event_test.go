package goaction

// Code auto generated with `go run ./internal/genevents/main.go`. DO NOT EDIT

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCheckRun(t *testing.T) {
	if Event != EventCheckRun {
		t.Skipf("Only applicatble for 'check run'")
	}
	event, err := GetCheckRun()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetCheckSuite(t *testing.T) {
	if Event != EventCheckSuite {
		t.Skipf("Only applicatble for 'check suite'")
	}
	event, err := GetCheckSuite()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetCreate(t *testing.T) {
	if Event != EventCreate {
		t.Skipf("Only applicatble for 'create'")
	}
	event, err := GetCreate()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetDelete(t *testing.T) {
	if Event != EventDelete {
		t.Skipf("Only applicatble for 'delete'")
	}
	event, err := GetDelete()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetDeployment(t *testing.T) {
	if Event != EventDeployment {
		t.Skipf("Only applicatble for 'deployment'")
	}
	event, err := GetDeployment()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetFork(t *testing.T) {
	if Event != EventFork {
		t.Skipf("Only applicatble for 'fork'")
	}
	event, err := GetFork()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetGollum(t *testing.T) {
	if Event != EventGollum {
		t.Skipf("Only applicatble for 'gollum'")
	}
	event, err := GetGollum()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetIssueComment(t *testing.T) {
	if Event != EventIssueComment {
		t.Skipf("Only applicatble for 'issue comment'")
	}
	event, err := GetIssueComment()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetIssues(t *testing.T) {
	if Event != EventIssues {
		t.Skipf("Only applicatble for 'issues'")
	}
	event, err := GetIssues()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetLabel(t *testing.T) {
	if Event != EventLabel {
		t.Skipf("Only applicatble for 'label'")
	}
	event, err := GetLabel()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetMilestone(t *testing.T) {
	if Event != EventMilestone {
		t.Skipf("Only applicatble for 'milestone'")
	}
	event, err := GetMilestone()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetPageBuild(t *testing.T) {
	if Event != EventPageBuild {
		t.Skipf("Only applicatble for 'page build'")
	}
	event, err := GetPageBuild()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetProject(t *testing.T) {
	if Event != EventProject {
		t.Skipf("Only applicatble for 'project'")
	}
	event, err := GetProject()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetProjectCard(t *testing.T) {
	if Event != EventProjectCard {
		t.Skipf("Only applicatble for 'project card'")
	}
	event, err := GetProjectCard()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetPublic(t *testing.T) {
	if Event != EventPublic {
		t.Skipf("Only applicatble for 'public'")
	}
	event, err := GetPublic()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetPullRequest(t *testing.T) {
	if Event != EventPullRequest {
		t.Skipf("Only applicatble for 'pull request'")
	}
	event, err := GetPullRequest()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetPullRequestReview(t *testing.T) {
	if Event != EventPullRequestReview {
		t.Skipf("Only applicatble for 'pull request review'")
	}
	event, err := GetPullRequestReview()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetPullRequestReviewComment(t *testing.T) {
	if Event != EventPullRequestReviewComment {
		t.Skipf("Only applicatble for 'pull request review comment'")
	}
	event, err := GetPullRequestReviewComment()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetPush(t *testing.T) {
	if Event != EventPush {
		t.Skipf("Only applicatble for 'push'")
	}
	event, err := GetPush()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetRelease(t *testing.T) {
	if Event != EventRelease {
		t.Skipf("Only applicatble for 'release'")
	}
	event, err := GetRelease()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetStatus(t *testing.T) {
	if Event != EventStatus {
		t.Skipf("Only applicatble for 'status'")
	}
	event, err := GetStatus()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetWatch(t *testing.T) {
	if Event != EventWatch {
		t.Skipf("Only applicatble for 'watch'")
	}
	event, err := GetWatch()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}

func TestGetRepositoryDispatch(t *testing.T) {
	if Event != EventRepositoryDispatch {
		t.Skipf("Only applicatble for 'repository dispatch'")
	}
	event, err := GetRepositoryDispatch()
	assert.NoError(t, err)

	var out bytes.Buffer
	err = json.NewEncoder(&out).Encode(event)
	require.NoError(t, err)
	t.Log(out.String())
}
