//go:build ignore

package main

var gone = map[byte]bool{
	'a': true,
	'r': true,
	's': true,
	'd': true,
	'h': true,
	'p': true,
	'y': true,
	'u': true,
	'o': true,
}

var guessed = [5]byte{0, 0, 'n', 'c', 0}

// var guessed = [5]byte{0, 0, 0, 0, 0}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	{'c': true},
	{'e': true},
	nil,
	{'e': true},
	nil,
}
