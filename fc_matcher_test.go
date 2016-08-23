package main

import (
	"fmt"
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
	var tests = map[string]map[string]float64{
		"West Bromwich Albion Football Club": map[string]float64{
			"West Bromwich Albion Football Club": 1.0,
			"West Bromwich Albion":               0.5,
			"West Bromwich Albion FC":            0.5,
			"West Bromwich Albion F.C.":          0.5,
			"West Bromwich FC":                   0.5,
			"West Bromwich":                      0.5,
			"WBA FC":                             0.5,
			"WB Albion":                          0.5,
			"West Brom":                          0.5,
			"New Mills":                          0.0,
			"":                                   0.0,
		},
		"Arsenal Football Club": map[string]float64{
			"Arsenal Football Club": 1.0,
			"Arsenal":               0.5,
			"Arsenal FC":            0.5,
		},
		"": map[string]float64{
			"What?": 0.0,
		},
	}

	fc := &FCMatcher{}
	for original := range tests {
		for key := range tests[original] {
			result, err := fc.Match(original, key)
			assert.Nil(t, err)
			assert.Condition(t, func() bool {
				fmt.Println(key, "~", original)
				fmt.Println(result, ">=", tests[original][key])
				return result >= tests[original][key]
			})
		}
	}
}
