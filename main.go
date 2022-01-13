package main

import (
	"fmt"
	"math"
	"sort"
)

var gone = map[byte]bool{
	// 'c': true,
}

var guessed = [5]byte{0, 0, 0, 0, 0}
var mustGuess int

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
	mustGuess = 0
	for i := range guessed {
		if guessed[i] == 0 {
			mustGuess++
			continue
		}
		lettersInWord[guessed[i]] = true
	}
	for i := range inWord {
		for j := range inWord[i] {
			lettersInWord[j] = true
		}
	}
	for letter := range lettersInWord {
		possiblePositions := [5]bool{}
		count := 0
		for j := range possiblePositions {
			if guessed[j] == letter {
				// we found where the letter is
				count = 0
				break
			}
			if guessed[j] != 0 {
				// can't put here
				continue
			}
			if inWord[j][letter] {
				// can't put here because we got yellow at this spot
				continue
			}
			possiblePositions[j] = true
			count++
		}
		if count == 1 {
			for j := range possiblePositions {
				if possiblePositions[j] && guessed[j] == 0 {
					fmt.Println("assigning", string(letter), "to only remaining position", j+1)
					guessed[j] = letter
					mustGuess--
				}
			}
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

func informationGained(word []byte, eligibleWord bool, wordsLeft int, debug bool) float64 {
	sum := [5]float64{}
	repeats := make(map[byte]bool)
	for i, letter := range word {
		if gone[letter] {
			continue
		}
		if _, ok := frequency[letter]; !ok {
			// there are some words with capital letters
			return 0
		}
		// However, we _can_ gain information about the other positions
		otherPositionSum := float64(0)
		for j := 0; j < 5; j++ {
			if j != i && guessed[j] == 0 {
				otherPositionSum += frequencyV2[j][letter] * 1 / float64(mustGuess)
			}
		}
		otherPositionSum = otherPositionSum / 4
		otherPositionInfo := float64(0)
		if otherPositionSum > 0 {
			otherPositionInfo = otherPositionSum * -1 * math.Log(otherPositionSum)
			if eligibleWord {
				//fmt.Println(string(word), "other position", otherPositionSum, otherPositionInfo)
			}
		}
		if guessed[i] != 0 && guessed[i] == letter {
			// we gain 0 information from guessing a letter in a position we
			// already know
			sum[i] = otherPositionInfo / 3
			continue
		}
		if inWord[i][guessed[i]] {
			// we've already guessed this letter in this position, we gain
			// 0 from guessing it again
			sum[i] = otherPositionInfo
			continue
		}
		if val, ok := frequencyV2[i][letter]; !ok || val == 0 {
			// we gain zero information at this position by guessing a letter
			// that doesn't fit with any of the current words
			sum[i] = otherPositionInfo
			continue
		}
		fraction := float64(1)
		if guessed[i] != 0 {
			// we've already guessed the letter in this position, but if we
			// guess a different letter, we can detect a position somewhere else
			fraction = fraction * (float64(mustGuess) * 1.5) / float64(10)
		}
		for j := range guessed {
			if i != j && guessed[j] == letter {
				// a letter that was guessed elsewhere in the word can still be
				// in a different position in the word, but assign it a lower
				// probability - more likely we are stepping on some other
				// letter that can go here.
				fraction = fraction * 1 / float64(mustGuess)
			}
		}
		if repeats[letter] {
			// We can still guess it exactly but we won't get more info about
			// whether this letter is somewhere else in the word.
			//
			// However the utility of repeats declines as there are fewer words
			// left to guess
			fraction = fraction * 1 / float64(mustGuess)
		}
		repeats[letter] = true
		freq, ok := frequencyV2[i][letter]
		if !ok {
			return 0
		}
		sum[i] += fraction*freq*-1*math.Log(freq) + otherPositionInfo/4
	}
	rval := float64(0)
	for i := range sum {
		rval += sum[i]
	}
	if eligibleWord {
		// We gain a lot of information if we guess the word correctly!
		wordGuessBonus := 1 / float64(wordsLeft) * -1 * math.Log(1/float64(wordsLeft))
		rval += wordGuessBonus
	}
	if debug {
		fmt.Printf("%s", string(word))
		for i := range sum {
			pval := sum[i]
			if pval != 0 {
				pval = pval
			}
			fmt.Printf(" %0.2f", pval)
			rval += sum[i]
		}
		fmt.Printf("\n")
	}
	return rval
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
	scores := make([]score, 0)
	eligibleWords := make([]string, 0)
	eligibleMap := make(map[string]bool, 0)
	for _, word := range words {
		if len(word) != 5 {
			continue
		}
		bword := []byte(word)
		if eligible(bword) {
			eligibleWords = append(eligibleWords, word)
			eligibleMap[word] = true
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
	for _, word := range words {
		bword := []byte(word)
		entropy := informationGained(bword, eligibleMap[word], len(eligibleMap), false)
		scores = append(scores, score{word: word, entropy: entropy})
		sort.Slice(scores, func(i, j int) bool {
			return scores[i].entropy > scores[j].entropy
		})
	}
	for i := range scores {
		fmt.Printf("%s %0.4f ", scores[i].word, scores[i].entropy)
		informationGained([]byte(scores[i].word), eligibleMap[scores[i].word], len(eligibleMap), true)
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
	//for i := 0; i < 5; i++ {
	//for j := 'a'; j <= 'z'; j++ {
	//fmt.Printf("%s %0.4f\n", string(j), frequencyV2[i][byte(j)])
	//}
	//fmt.Println("")
	//}
}
