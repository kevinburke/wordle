//go:build ignore

package main

var gone = map[byte]bool{
	'a': true,
	'e': true,
	't': true,
	'c': true,
	'b': true,
	'o': true,
	's': true,
	'y': true,
	'u': true,
}

var guessed = [5]byte{'d', 'r', 0, 'n', 'k'}

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	nil,
	{'r': true},
	nil,
	nil,
}
