//go:build ignore

package main

var gone = map[byte]bool{
	'a': true,
	'n': true,
	'r': true,
	'p': true,
}

// var guessed = [5]byte{0, 0, 0, 0, 0}
var guessed = [5]byte{0, 0, 0, 0, 0}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	{'s': true},
	{'o': true},
	{'e': true},
	{'e': true, 't': true},
	{'s': true},
}
