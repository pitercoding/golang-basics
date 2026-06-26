package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Analysis struct {
	Text            string
	CharacterCount  int
	LetterCount     int
	WordCount       int
	WordFrequency   map[string]int
	LetterFrequency map[rune]int
	MostCommonWord  string
	MostCommonRune  rune
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	text := readString(scanner, "Enter text: ")

	analytics := analyzeText(text)

	fmt.Println("\n=== Text Analyzer ===")
	fmt.Printf("Text: %s\n", analytics.Text)
	fmt.Printf("Characters: %d\n", analytics.CharacterCount)
	fmt.Printf("Letters: %d\n", analytics.LetterCount)
	fmt.Printf("Words: %d\n", analytics.WordCount)
	fmt.Printf("Most Common Word: %s\n", analytics.MostCommonWord)
	fmt.Printf("Most Common Letter: %c\n", analytics.MostCommonRune)

	fmt.Println("\n=== Word Frequency ===")
	for word, count := range analytics.WordFrequency {
		fmt.Printf("%s: %d\n", word, count)
	}

	fmt.Println("\n=== Letter Frequency ===")
	for letter, count := range analytics.LetterFrequency {
		fmt.Printf("%c: %d\n", letter, count)
	}
}

func analyzeText(text string) Analysis {
	return Analysis{
		Text:            text,
		CharacterCount:  countCharacters(text),
		WordCount:       countWords(text),
		LetterCount:     countLetters(text),
		WordFrequency:   wordFrequency(text),
		LetterFrequency: letterFrequency(text),
		MostCommonWord:  mostCommonWord(text),
		MostCommonRune:  mostCommonRune(text),
	}
}

func countCharacters(text string) int {
	return utf8.RuneCountInString(text)
}

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func countLetters(text string) int {
	count := 0

	for _, char := range text {
		if unicode.IsLetter(char) {
			count++
		}
	}

	return count
}

func wordFrequency(text string) map[string]int {
	freq := make(map[string]int)

	text = strings.ToLower(text)
	words := strings.Fields(text)

	for _, word := range words {
		word = strings.Trim(word, ".,!?;:")
		freq[word]++
	}

	return freq
}

func letterFrequency(text string) map[rune]int {
	freq := make(map[rune]int)

	text = strings.ToLower(text)

	for _, char := range text {
		if unicode.IsLetter(char) {
			freq[char]++
		}
	}

	return freq
}

func mostCommonWord(text string) string {
	if len(strings.TrimSpace(text)) == 0 {
		return ""
	}

	freq := wordFrequency(text)

	maxCount := 0
	mostCommon := ""

	for word, count := range freq {
		if count > maxCount {
			maxCount = count
			mostCommon = word
		}
	}

	return mostCommon
}

func mostCommonRune(text string) rune {
	freq := letterFrequency(text)

	maxCount := 0
	var mostCommon rune

	for char, count := range freq {
		if count > maxCount {
			maxCount = count
			mostCommon = char
		}
	}

	return mostCommon
}

func readString(scanner *bufio.Scanner, message string) string {
	for {
		fmt.Print(message)

		scanner.Scan()

		value := strings.TrimSpace(scanner.Text())

		if value != "" {
			return value
		}

		fmt.Println("Text cannot be empty! Try again.")
	}
}
