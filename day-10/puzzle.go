package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	instructions := strings.Split(strings.TrimSpace(string(file)), "\n")
	register, cycleCount, signalSum := 1, 0, 0
	instructionValue := 0

	var instructionData []string
	var image [6][40]string

	for _, instruction := range instructions {
		cycleCount++
		signalSum += calculateSignal(register, cycleCount)
		crtDraw(register, cycleCount, &image)

		instructionData = strings.Split(instruction, " ")
		if instructionData[0] == "noop" {
			continue
		}

		cycleCount++
		signalSum += calculateSignal(register, cycleCount)
		crtDraw(register, cycleCount, &image)

		instructionValue, _ = strconv.Atoi(instructionData[1])
		register += instructionValue
	}

	fmt.Printf("Solution 1 is: %v\n\n", signalSum)
	fmt.Printf("Solution 2 is: \n")
	for _, row := range image {
		fmt.Printf("%v\n", row)
	}
}

func crtDraw(register int, cycleCount int, image *[6][40]string) {
	h := (cycleCount - 1) / 40
	w := (cycleCount - 1) % 40

	if cycleCount%40 != register && cycleCount%40 != register+1 && cycleCount%40 != register+2 {
		(*image)[h][w] = "."
	} else {
		(*image)[h][w] = "#"
	}
}

func calculateSignal(register int, cycleCount int) int {
	if (cycleCount-20)%40 != 0 {
		return 0
	}

	return register * cycleCount
}
