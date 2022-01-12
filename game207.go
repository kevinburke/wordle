//go:build ignore

package main

var gone = map[byte]bool{
	'c': true,
	'e': true,
	't': true,
	'l': true,
	'i': true,
	'y': true,
	's': true,
	//'i': true,
	//'y': true,
	//'d': true,
	//'n': true,
	//'s': true,
	//'u': true,
	//'p': true,
}

var guessed = [5]byte{0, 'a', 'v', 'o', 'r'}
var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	nil,
	{'r': true},
	{'r': true},
	nil,
}
