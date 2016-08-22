package main

// FCMatcher trying to match two football club names and tells their similarity
// It implements Matcher interface
type FCMatcher struct {
}

// Match matches original to alternative and returns the similarity factor
// similar = 1.0 means total match
// similar = 0.0 means no match at all
func (fcm *FCMatcher) Match(original string, alternative string) (similarity float64, err error) {
	if original == alternative {
		return 1.0, nil
	}

	return
}
