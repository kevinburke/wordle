//go:build ignore
package main

import (
	"fmt"
	"math"
	"sort"
)

var gone = map[byte]bool{
	// 'c': true,
	'c': true,
	'a': true,
	'e': true,
	's': true,
	'd': true,
	'n': true,
}

var guessed = [5]byte{0, 'r', 'o', 0, 'y'}

// var guessed = [5]byte{0, 0, 0, 0, 0}

var mustGuess int

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	nil,
	{'r': true},
	nil,
	nil,
}
package main
