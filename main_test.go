package main

import (
	"testing"

	"github.com/bfontaine/stargazer/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestGotDM(t *testing.T) {
	assert.Equal(t, enablingMessage, gotDM("d123", "u1", "enable"))
	assert.Equal(t, enablingMessage, gotDM("d123", "u2", "enable"))
	assert.Equal(t, enabledMessage, gotDM("d123", "u1", "enable"))

	assert.Equal(t, disabledMessage, gotDM("d123", "u3", "disable"))
	assert.Equal(t, disablingMessage, gotDM("d123", "u1", "disable"))
	assert.Equal(t, disabledMessage, gotDM("d123", "u1", "disable"))

	assert.Equal(t, "", gotDM("d123", "u4", "hello who are you"))
}
