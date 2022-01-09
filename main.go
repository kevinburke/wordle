package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

var gone = map[byte]bool{
	//'t': true,
}

var guessed = [5]byte{0, 0, 0, 0, 0}

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	nil,
	nil,
	nil,
	nil,
	nil,
}

var lettersInWord = map[byte]bool{}

func init() {
	for i := range guessed {
		if guessed[i] == 0 {
			continue
		}
		lettersInWord[guessed[i]] = true
	}
	for i := range inWord {
		for j := range inWord[i] {
			lettersInWord[j] = true
		}
	}
}

// https://www.sttmedia.com/characterfrequency-english
var frequency = map[byte]float64{
	'e': 0.1260,
	't': 0.0937,
	'a': 0.0834,
	'o': 0.0770,
	'n': 0.0680,
	'i': 0.0671,
	'h': 0.0611,
	's': 0.0611,
	'r': 0.0568,
	'l': 0.0424,
	'd': 0.0414,
	'u': 0.0285,
	'c': 0.0273,
	'm': 0.0253,
	'w': 0.0234,
	'y': 0.0204,
	'f': 0.0203,
	'g': 0.0192,
	'p': 0.0166,
	'b': 0.0154,
	'v': 0.0106,
	'k': 0.0087,
	'j': 0.0023,
	'x': 0.0020,
	'q': 0.0009,
	'z': 0.0006,
}

func validWord(word []byte) bool {
	if len(word) != 5 {
		return false
	}
	for i, letter := range word {
		if gone[letter] {
			return false
		}
		if guessed[i] != 0 && letter != guessed[i] {
			return false
		}
	}
	return true
}

func informationGained(word []byte) float64 {
	sum := float64(0)
	repeats := make(map[byte]bool)
	for i, letter := range word {
		if _, ok := frequency[letter]; !ok {
			return 0
		}
		if gone[letter] {
			continue
		}
		if guessed[i] != 0 && guessed[i] == letter {
			// we gain 0 information from guessing a letter in a position we
			// already know
			continue
		}
		if inWord[i][guessed[i]] {
			// we've already guessed this letter in this position, we gain
			// 0 from guessing it again
			continue
		}
		fraction := float64(1)
		if guessed[i] != 0 {
			// we can't guess the word, but we can detect a position somewhere
			// else
			fraction = fraction * 1 / float64(3)
		}
		for j := range guessed {
			if i != j && guessed[j] == letter {
				// letter that was guessed elsewhere can still be in the word,
				// but assign it a lower probability - more likely we are
				// stepping on some other letter that can go here.
				fraction = fraction * 1 / float64(3)
			}
		}
		if repeats[letter] {
			// We can still guess it exactly but we won't get more info about
			// whether this letter is somewhere else in the word.
			fraction = fraction * 1 / float64(4)
		}
		repeats[letter] = true
		freq, ok := frequency[letter]
		if !ok {
			return 0
		}
		sum += fraction * freq * math.Log(freq)
	}
	return -1 * sum
}

func eligible(word []byte) bool {
	wordCopy := make(map[byte]bool, len(lettersInWord))
	for i := range lettersInWord {
		wordCopy[i] = true
	}
	for i, letter := range word {
		if gone[letter] {
			return false
		}
		if _, ok := frequency[letter]; !ok {
			return false
		}
		if inWord[i][letter] {
			// we know that this letter is in the word but _not_ in this
			// position
			return false
		}
		if guessed[i] != 0 && letter != guessed[i] {
			return false
		}
		delete(wordCopy, letter)
	}
	if len(wordCopy) > 0 {
		return false
	}
	return true
}

type score struct {
	word    string
	entropy float64
}

func main() {
	wordsAll, err := os.ReadFile("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	wordsSplit := bytes.Split(wordsAll, []byte{'\n'})
	fmt.Println(len(wordsSplit))
	scores := make([]score, 0)
	eligibleWords := make([]string, 0)
	for _, word := range wordsSplit {
		if len(word) != 5 {
			continue
		}
		if eligible(word) {
			eligibleWords = append(eligibleWords, string(word))
		}
	}
	// find letters which don't appear in any of the eligible words, and add
	// them to the 'gone' list
	missingLetters := make(map[byte]bool)
	for i := range frequency {
		missingLetters[i] = true
	}
	for i := range gone {
		delete(missingLetters, i)
	}
	for _, word := range eligibleWords {
		for i := range word {
			delete(missingLetters, word[i])
		}
		if len(missingLetters) == 0 {
			break
		}
	}
	for i := range missingLetters {
		fmt.Println("adding", string(i), "to 'gone' list")
		gone[i] = true
	}
	for _, word := range wordsSplit {
		if len(word) != 5 {
			continue
		}
		entropy := informationGained(word)
		scores = append(scores, score{word: string(word), entropy: entropy})
		sort.Slice(scores, func(i, j int) bool {
			return scores[i].entropy > scores[j].entropy
		})
	}
	for i := range scores {
		fmt.Println(scores[i].word, scores[i].entropy)
		if i > 50 {
			break
		}
	}
	fmt.Println("")
	sort.Strings(eligibleWords)
	fmt.Printf("%d eligible words", len(eligibleWords))
	if len(eligibleWords) < 150 {
		fmt.Printf(":\n")
		for i := range eligibleWords {
			fmt.Println(eligibleWords[i])
		}
	} else {
		fmt.Printf("\n")
	}
}
