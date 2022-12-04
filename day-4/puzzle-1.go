package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var pairs []string
	var countFully, countOverlap int

	pairs = parseFile("assignment-pairs")

	for _, pair := range pairs {
		xMin, xMax, yMin, yMax := getRanges(pair)
		isFullyContained := (xMin <= yMin && xMax >= yMax) || (xMin >= yMin && xMax <= yMax)
		isOverlapping := (xMin >= yMin && xMin <= yMax) || (xMax >= yMin && xMax <= yMax) || isFullyContained

		if isFullyContained {
			countFully++
		}

		if isOverlapping {
			countOverlap++
		}
	}

	fmt.Printf("Fully contained: %v\n", countFully)
	fmt.Printf("Overlap: %v\n", countOverlap)
}

func getRanges(pair string) (int, int, int, int) {
	pairData := strings.FieldsFunc(pair, func(delimiter rune) bool {
		return delimiter == '-' || delimiter == ','
	})

	a, _ := strconv.Atoi(pairData[0])
	b, _ := strconv.Atoi(pairData[1])
	c, _ := strconv.Atoi(pairData[2])
	d, _ := strconv.Atoi(pairData[3])

	return a, b, c, d
}

func parseFile(filename string) []string {
	var lines []string

	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines = strings.Split(strings.TrimSpace(string(file)), "\n")

	return lines
}
