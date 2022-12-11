package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var duplicates, badges string
	rucksacks := parseFile("input")

	for _, rucksack := range rucksacks {
		itemsPerCompartment := len(rucksack) / 2

		firstCompartment := rucksack[:itemsPerCompartment]
		secondCompartment := rucksack[itemsPerCompartment:]

		duplicates += findDuplicates(firstCompartment, secondCompartment)
	}

	prioritySum := sum(mapToCodes(duplicates))
	fmt.Printf("Priority sum: %v\n", prioritySum)

	for i := 0; i < len(rucksacks); i += 3 {
		x := rucksacks[i]
		y := rucksacks[i+1]
		z := rucksacks[i+2]

		badge := findDuplicates(findDuplicates(x, y), z)
		badges += badge
	}

	badgesPrioritySum := sum(mapToCodes(badges))
	fmt.Printf("Badges priority sum: %v\n", badgesPrioritySum)
}

func findDuplicates(first string, second string) string {
	var duplicates string
	var j int

	for _, i := range first {
		j = strings.IndexRune(second, i)

		if j >= 0 && (strings.IndexRune(duplicates, i) < 0) {

			duplicates += string(second[j])
		}
	}

	return duplicates
}

func sum(array []int) int {
	sum := 0

	for _, i := range array {
		sum += i
	}

	return sum
}

func mapToCodes(word string) []int {
	var codes []int

	for _, character := range word {
		asciiCode := int(character)
		if (asciiCode >= 65) && (asciiCode <= 96) {
			codes = append(codes, asciiCode-38)
		} else {
			codes = append(codes, asciiCode-96)
		}
	}

	return codes
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
