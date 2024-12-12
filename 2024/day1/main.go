package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== day 1 ===")
	lines := readFile("./input.txt")

	//==================== part 1
	distance := 0
	arrL, arrR := splitAndSortInput(lines)
	for i := 0; i < len(lines); i++ {
		distance += distanceBetween(float64(arrL[i]), float64(arrR[i]))
	}
	fmt.Println("dist: ", distance)

	//==================== part 2

	//calculate frequency of every number in arrR
	freq := make(map[int]int)
	for _, num := range arrR {
		freq[num] = freq[num] + 1
	}

	//filter duplicates in left slice
	slices.Sort(arrL)
	arrL = slices.Compact(arrL)

	similarity_score := 0
	for _, v := range arrL {
		similarity_score += (v * freq[v])
	}
	fmt.Println("similarity_score: ", similarity_score)
}

func splitAndSortInput(lines []string) ([]int, []int) {
	arrleft := []int{}
	arrright := []int{}

	//split lines into arrays
	for _, v := range lines {
		vals := strings.Fields(v)

		d, _ := strconv.Atoi(vals[0])
		arrleft = append(arrleft, d)

		e, _ := strconv.Atoi(vals[1])
		arrright = append(arrright, e)
	}

	//sort
	sort.Ints(arrleft)
	sort.Ints(arrright)
	return arrleft, arrright
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

// read one line
func distanceBetween(x, y float64) int {
	if x > y {
		return int(math.Abs(x - y))
	} else {
		return int(math.Abs(y - x))
	}
}
