package main

import (
	"fmt"
)

// FCMatcher trying to match two football club names and tells their similarity
// It implements Matcher interface
type FCMatcher struct {
}

// Match matches original to alternative and returns the similarity factor
// similar = 1.0 means total match
// similar = 0.0 means no match at all
func (fcm *FCMatcher) Match(original string, alternative string) (similarity float64, err error) {
	o := cleanString(original)
	a := expandString(alternative)
	a = cleanString(a)

	fmt.Println(a)

	if o == a {
		return 1.0, nil
	}

	return
}

/*

// Weights are just arbitrary, will change once implementation is ready

original := "West Bromwich Albion Football Club"

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

To calculate similarity score of Football Team name, we will use couple of
strategies:
1) We clean both input strings
2) We drop "Football Club", "FC", "F.C." (???)
3) Baseline will be number of words in football team name
4) Next we expand alternative string by capital letter as if it was separate word
5) We match unmatched words by how much of alternative word is in original
    B - Bromwich, that is 1/8
    we need to shift the result a bit, so that 1/8 will be 0.5 and 8/8 will be 1.0
    0.5 + (1/8)/2
    if there is no match then we put 0

*/
