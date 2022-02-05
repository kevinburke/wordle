//go:build ignore

package main

var gone = map[byte]bool{
	's': true,
	'n': true,
	'r': true,
	'b': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{0, 'l', 0, 0, 0}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	{'a': true},
	{'a': true},
	{'e': true, 't': true},
	{'e': true},
}
