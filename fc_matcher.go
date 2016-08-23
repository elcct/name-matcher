package main

import (
	//	"fmt"
	"strings"
)

/*
Algorithm draft

To calculate similarity score of Football Team name, we will use couple of
strategies:

- We clean both input strings

- We assume that order of words in alternative name is the same as in original.
  For example alternative to West Bromwich will not be Bromwich West

- Next we expand alternative string by capital letter as if it was separate word
  For example: ABeC becomes A Be C

- We iterate over original words and checking how much of alternative word is
  in its original counterpart
    B - Bromwich, that is 1/8
    If we have partial match, we can assume it is abbreviation of the word
    We fit the weight to the range of 1.0>=weight>=0.5
    If there is no match then we put 0 weight

- We calculate the similarity as sum of weights divided by number
   of match attempts performed

- If resulting similarity is >=0.5 we can assume we have a likely match
    If similarity is 1.0 then it is a perfect match
    If similarity is <0.5 then match is unlikely
    If similarity is 0.0 then for sure there is no match
*/

// FCMatcher trying to match two football club names and tells their similarity
// It implements Matcher interface
type FCMatcher struct {
}

// Match matches original to alternative and returns the similarity factor
// similarity = 1.0 means total match
// similarity >= means likely match
// similarity = 0.0 means no match at all
func (fcm *FCMatcher) Match(original string, alternative string) (similarity float64, err error) {
	o := cleanString(original)
	a := expandString(alternative)
	a = cleanString(a)

	if o == a {
		return 1.0, nil
	}

	po := strings.Fields(o)
	pa := strings.Fields(a)

	matches := 0
	weights := 0.0

	for i, token := range po {
		if i == len(pa) {
			break
		}
		if strings.HasPrefix(token, pa[i]) {
			weights += 0.5 + (float64(len(pa[i]))/float64(len(token)))*0.5
		}
		matches++
	}

	if matches == 0 {
		return 0.0, nil
	}

	similarity = weights / float64(matches)

	return
}
