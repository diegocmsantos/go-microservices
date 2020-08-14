package github

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestBodyAsJSON(t *testing.T) {
	requestBody := CreateRepoRequestBody{
		Name:        "golang instructions",
		Description: "a golang introduction repository",
		Homepage:    "http://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(requestBody)

	assert.NotNil(t, bytes)
	assert.Nil(t, err)
}

func TestCreateRepoRequestBodyFromJsonToStruct(t *testing.T) {

	// Arrange
	jsonAsString := "{\"name\":\"golang instructions\",\"description\":\"a golang introduction repository\"," +
		"\"homepage\":\"http://github.com\",\"private\":true,\"has_issues\":true,\"has_projects\":true,\"has_wiki\":true}"
	var target CreateRepoRequestBody

	// Act
	err := json.Unmarshal([]byte(jsonAsString), &target)

	// Arrange
	assert.Nil(t, err)
	assert.EqualValues(t, "golang instructions", target.Name)
	assert.EqualValues(t, "a golang introduction repository", target.Description)
	assert.EqualValues(t, "http://github.com", target.Homepage)
	assert.True(t, target.Private)
	assert.True(t, target.HasIssues)
	assert.True(t, target.HasProjects)
	assert.True(t, target.HasWiki)
}
