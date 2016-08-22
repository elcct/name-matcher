package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanString(t *testing.T) {
	source := " west  bromwich f. c. "
	expected := "West Bromwich F C"

	result := cleanString(source)
	assert.Equal(t, expected, result)
}

func TestExpandString(t *testing.T) {
	source := "ABeC"
	expected := "A Be C"

	result := expandString(source)
	assert.Equal(t, expected, result)
}
