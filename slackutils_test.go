package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemoizeGetter(t *testing.T) {
	var calls int

	f := func(s string) (string, error) {
		calls++
		if s == "error" {
			return "", errors.New("oops")
		}
		return "called:" + s, nil
	}

	memoized := memoizeGetter(f)

	require.NotNil(t, memoized)
	assert.Equal(t, 0, calls)

	s, err := memoized("yo")
	assert.Nil(t, err)
	assert.Equal(t, "called:yo", s)
	assert.Equal(t, 1, calls)

	s, err = memoized("yo")
	assert.Nil(t, err)
	assert.Equal(t, "called:yo", s)
	assert.Equal(t, 1, calls)

	s, err = memoized("error")
	assert.NotNil(t, err)
	assert.Equal(t, 2, calls)

	s, err = memoized("error")
	assert.NotNil(t, err)
	assert.Equal(t, 3, calls)
}
