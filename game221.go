//go:build ignore

package main

var gone = map[byte]bool{
	'r': true,
	'e': true,
	's': true,
	'p': true,
	'l': true,
	'i': true,
	't': true,
	'o': true,
	'n': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{0, 0, 0, 0, 0}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	{'c': true, 'a': true},
	{'a': true, 'c': true},
	{'c': true},
	{'a': true},
	nil,
}
