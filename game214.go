//go:build ignore

package main

var gone = map[byte]bool{
	'c': true,
	'a': true,
	'r': true,
	'e': true,
	's': true,
	'y': true,
}

var guessed = [5]byte{'p', 0, 0, 'n', 0}

// var guessed = [5]byte{0, 0, 0, 0, 0}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	{'i': true},
	{'o': true},
	nil,
	nil,
}
