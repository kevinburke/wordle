//go:build ignore

package main

var gone = map[byte]bool{
	'a': true,
	'e': true,
	's': true,
	't': true,
	'o': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{0, 'r', 0, 'c', 'k'}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	{'c': true},
	nil,
	{'r': true},
	nil,
	nil,
}
