package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFCMatcherIsDefined(t *testing.T) {
	fc := &FCMatcher{}
	assert.NotNil(t, fc)
}

func TestFCMatcherSameNames(t *testing.T) {
	fc := &FCMatcher{}

	original := "West Bromwich Albion Football Club"

	result, err := fc.Match(original, original)
	assert.Nil(t, err)
	assert.Equal(t, result, 1.0)
}
