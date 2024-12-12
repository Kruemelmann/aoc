package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Main function to process the input and count safe reports
func main() {
	// Parse the input
	lines := readFile("./input.txt")

	count := 0
	for _, line := range lines {
		// Convert the line into an array of integers
		report := make([]int, 0)
		for _, str := range strings.Fields(line) {
			num, _ := strconv.Atoi(str)
			report = append(report, num)
		}

		// Check if the report is safe (either directly or after removal of one level)
		if isSafe(report) || isSafeAfterRemoval(report) {
			count++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", count)
}

// Function to check if a report is safe (both strictly increasing or strictly decreasing)
func isSafe(report []int) bool {
	isIncreasing := report[0] < report[1]

	//check if safe
	safeLevel := true
	for i := 1; i < len(report); i++ {
		c := report[i]   //c -> current
		l := report[i-1] //l -> last
		if isIncreasing {
			if l > c {
				safeLevel = false
				break
			}
			if (c-l) < 1 || (c-l) > 3 {
				safeLevel = false
				break
			}
		} else {
			if l < c {
				safeLevel = false
				break
			}
			if (l-c) < 1 || (l-c) > 3 {
				safeLevel = false
				break
			}
		}
	}

	return safeLevel
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// Function to check if a report is safe after removing one element
func isSafeAfterRemoval(report []int) bool {
	for i := 0; i < len(report); i++ {
		cpy := make([]int, len(report))
		copy(cpy, report)

		newReport := RemoveIndex(cpy, i)
		if isSafe(newReport) {
			return true
		}
	}
	return false
}

// Reads the file and returns the lines as an array
func readFile(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}
