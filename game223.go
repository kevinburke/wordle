//go:build ignore

package main

var gone = map[byte]bool{
	'a': true,
	'c': true,
	'e': true,
	's': true,
	'd': true,
	'o': true,
	'y': true,
	'b': true,
	'i': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{0, 'r', 0, 'n', 'g'}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	nil,
	{'r': true},
	nil,
	nil,
}
