//go:build ignore

package main

var gone = map[byte]bool{'t': true, 'o': true, 'e': true, 's': true, 'h': true}

var guessed = [5]byte{0, 0, 'a', 'n', 0}

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	{'a': true},
	nil,
	nil,
	{'r': true},
	nil,
}

var lettersInWord = map[byte]bool{}
