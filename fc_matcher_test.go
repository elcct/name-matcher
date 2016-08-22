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
	assert.Equal(t, 1.0, result)
}

func TestFCMatcherSimilarNames(t *testing.T) {
	var tests = map[string]float64{
		"West Bromwich Albion":      0.9,
		"West Bromwich Albion FC":   0.9,
		"West Bromwich Albion F.C.": 0.9,
		"West Bromwich FC":          0.9,
		"West Bromwich":             0.9,
		"WBA FC":                    0.6,
		"WB Albion":                 0.6,
		"West Brom":                 0.6,
	}

	fc := &FCMatcher{}

	original := "West Bromwich Albion Football Club"
	for key := range tests {
		result, err := fc.Match(original, key)
		assert.Nil(t, err)
		assert.Equal(t, tests[key], result)
	}
}
