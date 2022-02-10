//go:build ignore

package main

var gone = map[byte]bool{
	's': true,
	'a': true,
	'n': true,
	'p': true,
	'i': true,
	't': true,
	'o': true,
	'f': true,
	'u': true,
	'g': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{0, 0, 0, 'e', 'r'}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	nil,
	nil,
	{'l': true},
	nil,
}
