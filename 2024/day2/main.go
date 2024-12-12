package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== day 2 ===")

	lines := readFile("./input.txt")

	safeReports := 0
	for i := 0; i < len(lines); i++ {
		nums := strings.Fields(lines[i])
		f, _ := strconv.Atoi(nums[0])
		s, _ := strconv.Atoi(nums[1])
		isIncreasing := f < s

		//check if safe
		safeLevel := true
		for i := 1; i < len(nums); i++ {
			c, _ := strconv.Atoi(nums[i])   //c -> current
			l, _ := strconv.Atoi(nums[i-1]) //l -> last
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
		if safeLevel {
			safeReports++
		}
	}
	fmt.Println(safeReports)

}

// readFile read the file and return an array of strings
func readFile(path string) []string {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}
