//go:build ignore

package main

var gone = map[byte]bool{
	'c': true,
	'a': true,
	'e': true,
	's': true,
	'd': true,
	'n': true,
	'y': true,
	'v': true,
	'i': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{0, 0, 0, 0, 0}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	{'r': true},
	{'r': true, 'o': true, 't': true},
	{'r': true},
	{'o': true},
}
