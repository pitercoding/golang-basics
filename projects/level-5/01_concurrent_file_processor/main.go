package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type FileStats struct {
	FileName   string
	Characters int
	Words      int
	Lines      int
	Err        error
}

func main() {
	files := []string{
		"projects/level-5/01_concurrent_file_processor/sample1.txt",
		"projects/level-5/01_concurrent_file_processor/sample2.txt",
		"projects/level-5/01_concurrent_file_processor/sample3.txt",
	}

	results := make(chan FileStats)
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go worker(file, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var allStats []FileStats

	for stats := range results {
		allStats = append(allStats, stats)
	}

	printReport(allStats)
}

func worker(filename string, results chan<- FileStats, wg *sync.WaitGroup) {
	defer wg.Done()

	stats := processFile(filename)
	results <- stats
}

func processFile(filename string) FileStats {
	file, err := os.Open(filename)
	if err != nil {
		return FileStats{
			FileName: filename,
			Err:      err,
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		lines   int
		words   int
		chars   int
		content strings.Builder
	)

	for scanner.Scan() {
		line := scanner.Text()

		lines++
		words += len(strings.Fields(line))
		chars += len(line)

		content.WriteString(line)
		content.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		return FileStats{
			FileName: filename,
			Err:      err,
		}
	}

	return FileStats{
		FileName:   filename,
		Characters: chars,
		Words:      words,
		Lines:      lines,
		Err:        nil,
	}
}

func printReport(allStats []FileStats) {
	fmt.Println("\n=== FILE REPORT ===")

	totalLines, totalWords, totalChars := 0, 0, 0

	for _, stats := range allStats {
		if stats.Err != nil {
			fmt.Printf("File: %s -> ERROR: %v\n", stats.FileName, stats.Err)
			continue
		}

		fmt.Printf(
			"File: %-15s | Characters: %-5d | Words: %-5d | Lines: %-5d\n",
			stats.FileName,
			stats.Characters,
			stats.Words,
			stats.Lines,
		)

		totalChars += stats.Characters
		totalWords += stats.Words
		totalLines += stats.Lines
	}

	fmt.Println("--------------------------------------")
	fmt.Printf("TOTAL -> Files: %d | Lines: %d | Words: %d | Characters: %d\n",
		len(allStats),
		totalLines,
		totalWords,
		totalChars,
	)
}
