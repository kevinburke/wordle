//go:build ignore
// +build ignore

package main

var gone = map[byte]bool{
	'a': true,
	't': true,
	'n': true,
	's': true,
	'c': true,
	'i': true,
	'h': true,
	'p': true,
	'd': true,
}

var guessed = [5]byte{0, 'o', 'r', 0, 'e'}

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	nil,
	{'o': true},
	{'o': true},
	nil,
}
