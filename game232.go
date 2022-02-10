//go:build ignore

package main

var gone = map[byte]bool{
	'a': true,
	'n': true,
	'e': true,
	'r': true,
	'p': true,
	'o': true,
	'u': true,
	't': true,
	'y': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{'s', 0, 0, 0, 0}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	nil,
	nil,
	nil,
	nil,
}
