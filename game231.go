//go:build ignore

package main

var gone = map[byte]bool{
	's': true,
	'n': true,
	'e': true,
	'r': true,
	'c': true,
	'y': true,
	'p': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{0, 'l', 'o', 0, 't'}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	{'a': true, 'o': true},
	{'a': true},
	{'l': true, 'a': true},
	nil,
}
