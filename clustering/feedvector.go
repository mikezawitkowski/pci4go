/*

feedvector.go

This is similar to the python file described in Chapter 03,
except that we will not be using the Feed Parser python module,
and we may change a few things about this to be more universal
in how it can be used.

TODO: finish this based on "generatepostsvector.py" in borg
TODO: find a stopwords library for Golang

*/
package main

import (
	"fmt"
	"strings"
)

// wordCounts is given a row and returns the title and dictionary of word counts for an RSS feed
func wordCounts(row string) {
	// TODO: change row to a map[string]string
	// TODO: return 2 items: the title of the 'row', and a map of words and the count of those words in the row
	fmt.Println("Finish this function!")
}

// words is given a string (html) and returns an array of words, broken by spaces
func words(html string) []string {
	// TODO: Strip html tags from the string before generating the word list
	return strings.Fields(strings.ToLower(html))
}

func main() {
	test := "A string for Testing words() function."
	out := words(test)
	for i := range out {
		fmt.Println(out[i])
	}
}
