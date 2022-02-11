//go:build ignore

package main

var gone = map[byte]bool{
	'n': true,
	'r': true,
	'h': true,
	't': true,
	'm': true,
	'l': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{'p', 'a', 'u', 0, 0}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	{'s': true},
	nil,
	{'e': true},
	{'e': true},
	{'s': true},
}
