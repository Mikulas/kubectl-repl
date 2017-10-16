package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClosestString(t *testing.T) {
	assert.Equal(t, "alpha", ClosestString("al", []string{"aX", "alpha"}),
		"should prefer insertion over replacement and deletion")

	assert.Equal(t, "alpha", ClosestString("aXpha", []string{"aX", "alpha"}),
		"should correct typos")
}
