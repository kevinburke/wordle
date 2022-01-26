//go:build ignore

package main

var gone = map[byte]bool{
	'c': true,
	'e': true,
	't': true,
	'w': true,
	'i': true,
	'm': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{'s', 0, 0, 'a', 'r'}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	{'a': true},
	{'r': true, 'a': true},
	{'r': true},
	{'s': true},
}
