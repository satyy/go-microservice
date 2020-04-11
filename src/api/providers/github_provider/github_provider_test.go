package github_provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthorizationHeader(t *testing.T)  {
	header := getAuthorizationHeader("test123")
	assert.EqualValues(t, "token test123", header)
}
