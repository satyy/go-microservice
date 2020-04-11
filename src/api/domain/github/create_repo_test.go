package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T)  {
	request := CreateRepoRequest{
		Name:        "Test-Repo",
		Description: "Test repo",
		Homepage:    "http://github.com",
		Private:     false,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     false,
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target CreateRepoRequest

	unmarshalError := json.Unmarshal(bytes, &target)
	assert.Nil(t, unmarshalError)

	assert.EqualValues(t,request.Name, target.Name)
	assert.EqualValues(t, request.HasWiki, target.HasWiki)
	fmt.Println(string(bytes))
}
