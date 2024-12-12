package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== day 4 ===")
	lines := readFile("./input_medium.txt")

	grid := [][]rune{}

	for _, v := range lines {
		grid = append(grid, []rune(v))
	}

	// Call the function
	result := countXmas(grid)
	fmt.Println("xmas count: ", result) // Output the result

	result = countXmas2(grid)
	fmt.Println("Xmas count: ", result) // Output the result
}

// Check if the word "XMAS" can be formed starting from (r, c) in the direction (dr, dc)
func checkDirection(grid [][]rune, r, c, dr, dc int) bool {
	word := "XMAS"
	for i := 0; i < len(word); i++ {
		nr, nc := r+i*dr, c+i*dc
		if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) || grid[nr][nc] != rune(word[i]) {
			return false
		}
	}
	return true
}

// Count how many times "XMAS" appears in the grid
func countXmas(grid [][]rune) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Check all 8 directions
	directions := [][2]int{
		{0, 1},   // Horizontal right
		{0, -1},  // Horizontal left
		{1, 0},   // Vertical down
		{-1, 0},  // Vertical up
		{1, 1},   // Diagonal down-right
		{-1, -1}, // Diagonal up-left
		{1, -1},  // Diagonal down-left
		{-1, 1},  // Diagonal up-right
	}

	// Traverse the grid and check all directions
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, direction := range directions {
				dr, dc := direction[0], direction[1]
				if checkDirection(grid, r, c, dr, dc) {
					count++
				}
			}
		}
	}

	return count
}

// Count how many X-MASes appear in the grid
func countXmas2(grid [][]rune) int {
	//TODO
	return -1
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
