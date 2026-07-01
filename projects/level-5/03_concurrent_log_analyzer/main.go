package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type LogStats struct {
	FileName    string
	TotalLines  int
	LevelCounts map[string]int // ex: {"ERROR": 5, "INFO": 120}
	ErrorLines  []string       // errors
	ProcessTime time.Duration
}

func extractLevel(line string) string {
	levels := []string{"ERROR", "WARN", "INFO"}

	for _, level := range levels {
		if strings.HasPrefix(line, "["+level+"]") {
			return level
		}
	}

	return ""
}

func analyzeFile(filename string) LogStats {
	start := time.Now()
	stats := LogStats{
		FileName:    filename,
		LevelCounts: make(map[string]int),
	}

	file, err := os.Open(filename)
	if err != nil {
		stats.ErrorLines = append(stats.ErrorLines, fmt.Sprintf("error while opening: %v", err))
		return stats
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stats.TotalLines++

		level := extractLevel(line)
		if level != "" {
			stats.LevelCounts[level]++
		}
		if level == "ERROR" {
			stats.ErrorLines = append(stats.ErrorLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		stats.ErrorLines = append(stats.ErrorLines, fmt.Sprintf("error while scanning: %v", err))
	}

	stats.ProcessTime = time.Since(start)
	return stats
}

func processFilesConcurrently(filenames []string) []LogStats {
	const workerCount = 4

	jobs := make(chan string)
	results := make(chan LogStats)

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// Send jobs
	go func() {
		for _, filename := range filenames {
			jobs <- filename
		}
		close(jobs)
	}()

	// Close results when workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	var allStats []LogStats

	for stats := range results {
		allStats = append(allStats, stats)
	}

	return allStats
}

func printReport(allStats []LogStats) {
	fmt.Println("==========================")
	fmt.Println("       REPORTS")
	fmt.Println("==========================")

	totalLevelCounts := make(map[string]int)
	totalLines := 0
	totalErrors := 0

	for _, stats := range allStats {
		fmt.Printf("\nFile: %s\n", stats.FileName)
		fmt.Printf("  Lines processed: %d\n", stats.TotalLines)
		fmt.Printf("  Time: %v\n", stats.ProcessTime)

		for level, count := range stats.LevelCounts {
			fmt.Printf("  %s: %d\n", level, count)
			totalLevelCounts[level] += count
		}

		if len(stats.ErrorLines) > 0 {
			fmt.Printf("  Errors found: %d\n", len(stats.ErrorLines))
			totalErrors += len(stats.ErrorLines)
		}

		totalLines += stats.TotalLines
	}

	fmt.Println("\n----------------------------------------")
	fmt.Println("               SUMMARY")
	fmt.Println("----------------------------------------")
	fmt.Printf("Total of lines: %d\n", totalLines)
	fmt.Printf("Total de errors: %d\n", totalErrors)
	fmt.Println("Log count per level:")
	for level, count := range totalLevelCounts {
		fmt.Printf("  %s: %d\n", level, count)
	}
	fmt.Println("========================================")

}

func worker(jobs <-chan string, results chan<- LogStats, wg *sync.WaitGroup) {
	defer wg.Done()

	for filename := range jobs {
		stats := analyzeFile(filename)
		results <- stats
	}
}

func main() {
	filenames := []string{
		"projects/level-5/03_concurrent_log_analyzer/logs.txt",
	}

	fmt.Println("\n**************************")
	fmt.Println("       LOG ANALYZER")
	fmt.Println("**************************")
	fmt.Printf("Processing %d files concurrently...\n\n", len(filenames))

	allStats := processFilesConcurrently(filenames)
	printReport(allStats)
}
