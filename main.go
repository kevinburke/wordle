package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

var ineligible = map[string]bool{
	// Words in /usr/share/dict/words that aren't accepted by Wordle.
	"soary": true,
	"barie": true,
	"solay": true,
	"seary": true,
	"sairy": true,
	"saily": true,
	"sorty": true,
	"tarie": true,
}

var gone = map[byte]bool{
	'c': true,
	'e': true,
	't': true,
	'l': true,
	'i': true,
	'y': true,
	'd': true,
	'n': true,
	's': true,
	'u': true,
	'p': true,
}

var guessed = [5]byte{0, 'a', 0, 'o', 'r'}

// In the word but in the wrong position
var inWord = [5]map[byte]bool{
	{'r': true},
	nil,
	{'r': true},
	{'r': true},
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

// Combines:
// the actual frequency at that letter based on the remaining words
// At 1/4 value, the frequency of all remaining letters in any position
// Letters we know are in the word but in different positions
var frequencyV2 = [5]map[byte]float64{}

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
		if gone[letter] {
			continue
		}
		if val, ok := frequencyV2[i][letter]; !ok || val == 0 {
			return 0
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
			// we've already guessed the letter in this position, but if we
			// guess a different letter, we can detect a position somewhere else
			fraction = fraction * 1 / float64(4)
		}
		for j := range guessed {
			if i != j && guessed[j] == letter {
				// a letter that was guessed elsewhere can still be in the
				// word, but assign it a lower probability - more likely we are
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
		freq, ok := frequencyV2[i][letter]
		if !ok {
			return 0
		}
		sum += fraction * freq * math.Log(freq)
	}
	return -1 * sum
}

func eligible(word []byte) bool {
	if ineligible[string(word)] {
		return false
	}
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
	scores := make([]score, 0)
	eligibleWords := make([]string, 0)
	for _, word := range wordsSplit {
		if len(word) != 5 {
			continue
		}
		sword := string(word)
		if eligible(word) {
			eligibleWords = append(eligibleWords, sword)
		}
	}
	if len(eligibleWords) == 1 {
		fmt.Println("the word is:", eligibleWords[0])
		return
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
	// Build frequency tables for each position from the remaining eligible
	// words
	frequencyV2 = [5]map[byte]float64{}
	for _, word := range eligibleWords {
		for i, letter := range []byte(word) {
			if frequencyV2[i] == nil {
				frequencyV2[i] = make(map[byte]float64)
			}
			if _, ok := frequencyV2[i][letter]; !ok {
				frequencyV2[i][letter] = 0
			}
			frequencyV2[i][letter] += 1
		}
	}
	for i := range frequencyV2 {
		for letter, count := range frequencyV2[i] {
			frequencyV2[i][letter] = count / float64(len(eligibleWords))
		}
	}
	for _, word := range wordsSplit {
		if len(word) != 5 {
			continue
		}
		sword := string(word)
		if ineligible[sword] {
			continue
		}
		entropy := informationGained(word)
		scores = append(scores, score{word: sword, entropy: entropy})
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
