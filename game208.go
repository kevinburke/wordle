//go:build ignore

package main

var gone = map[byte]bool{
	'c': true,
	'r': true,
	's': true,
	'd': true,
	'l': true,
	'n': true,
	'i': true,
}

var guessed = [5]byte{'a', 0, 0, 'e', 'y'}
var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	{'a': true},
	nil,
	nil,
	nil,
}
