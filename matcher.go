package main

// Matcher interface for different types of matchers
type Matcher interface {
	Match(string, string) (float64, error)
}
