//go:build ignore

package main

var gone = map[byte]bool{
	's': true,
	'a': true,
	'n': true,
	'e': true,
	'c': true,
	'l': true,
	't': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{0, 'u', 'm', 'o', 'r'}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	nil,
	{'o': true},
	{'u': true},
	nil,
}
