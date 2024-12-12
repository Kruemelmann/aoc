package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("=== day 3 ===")
	lines := readFile("./input.txt")

	r := regexp.MustCompile("mul" + regexp.QuoteMeta("(") + "[0-9]+,[0-9]+" + regexp.QuoteMeta(")"))

	re := regexp.MustCompile(regexp.QuoteMeta("don't()") + "*.*" + regexp.QuoteMeta("do()"))

	sum := 0
	//cut out all disabled sections in the memory
	split := re.Split(lines[0], -1)
	set := ""
	for i := range split {
		set += "\n" + split[i]
	}
	fmt.Println(set)

	instructions := r.FindAllString(set, -1)
	for _, in := range instructions {
		rs := regexp.MustCompile("[0-9]+")
		pair := rs.FindAllString(in, -1)
		fmt.Println(pair)
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		sum += x * y
	}

	fmt.Println(sum)
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
