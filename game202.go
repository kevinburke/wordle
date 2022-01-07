//go:build ignore

package main

var gone = map[byte]bool{
	't': true,
	'o': true,
	'n': true,
	'e': true,
	'h': true,
	'a': true,
	'r': true,
	'd': true,
	'f': true,
	'i': true,
}

var guessed = [5]byte{
	's',
	0,
	'u',
	0,
	'p',
}

var inWord = [5]map[byte]bool{
	nil,
	{'u': true},
	{'s': true},
	nil,
	{'l': true},
}
