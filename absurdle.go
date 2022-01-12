//go:build ignore

package main

var gone = map[byte]bool{
	'a': true,
	't': true,
	'o': true,
	'n': true,
	's': true,
	'h': true,
	'l': true,
	'f': true,
	'g': true,
	'v': true,
}

var guessed = [5]byte{0, 0, 'e', 0, 'y'}

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	{'r': true},
	nil,
	nil,
	nil,
}

var lettersInWord = map[byte]bool{}
