//go:build ignore

package main

var gone = map[byte]bool{
	'a': true,
	'r': true,
	'c': true,
	'e': true,
	's': true,
	'p': true,
	'i': true,
	't': true,
	'b': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{0, 0, 'o', 0, 0}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	{'o': true, 'l': true},
	nil,
	{'n': true},
	{'n': true},
}
