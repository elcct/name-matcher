package main

import (
	"flag"
	"fmt"
)

func main() {
	original := flag.String("original", "", "Original Football Club Name")
	alternative := flag.String("alternative", "", "Alternative Football Club Name")
	flag.Parse()

	fc := &FCMatcher{}
	result, err := fc.Match(*original, *alternative)
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}

	fmt.Println(result)

	return
}
