//go:build ignore

package main

var gone = map[byte]bool{
	'c': true,
	'a': true,
	't': true,
}

var guessed = [5]byte{0, 0, 'i', 0, 0}

// var guessed = [5]byte{0, 0, 0, 0, 0}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	{'r': true},
	{'e': true},
	{'r': true},
	{'e': true, 's': true},
	{'s': true},
}
