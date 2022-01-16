//go:build ignore

package main

var gone = map[byte]bool{
	'r': true,
	'e': true,
	's': true,
	'm': true,
	'h': true,
	't': true,
	'j': true,
	'k': true,
	'y': true,
}

var guessed = [5]byte{0, 'a', 0, 0, 0}
var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	{'c': true},
	nil,
	nil,
	{'c': true},
	nil,
}
