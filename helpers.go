package main

import (
	"bytes"
	"strings"
	"unicode"
)

// cleanString should clean string from white spaces and capitalize the input
// remove any punctuation marks
// TODO: optimize this :)
func cleanString(s string) (res string) {
	remover := func(r rune) bool {
		switch r {
		case '.', ',':
			return true
		}
		return false
	}

	res = strings.Join(strings.FieldsFunc(s, remover), "")
	res = strings.Join(strings.Fields(res), " ")
	res = strings.Title(res)
	return
}

// expandString should expand any abbreviation ABeC -> A Be C
// TODO: optimize this :)
func expandString(s string) (res string) {
	var buffer bytes.Buffer

	for _, v := range s {
		if unicode.IsUpper(v) {
			buffer.WriteString(" ")
		}
		buffer.WriteString(string(v))
	}

	return strings.Trim(buffer.String(), " ")
}
