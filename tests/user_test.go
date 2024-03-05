package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	assert := assert.New(t)

	err := truncate("Users", "Bookmarks")
	assert.Nil(err)
}
