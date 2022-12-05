package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(file), "\n\n")
	cargoRawPlan, moveRawPlan := strings.Split(inputs[0], "\n"), strings.Split(strings.TrimSpace(inputs[1]), "\n")
	cargoPlan := make(map[int][]string)

	for _, cargoLayer := range cargoRawPlan[:len(cargoRawPlan)-1] {
		for i, crate := range cargoLayer {
			if (i-1)%4 != 0 || crate == ' ' {
				continue
			}

			mark := ((i - 1) / 4) + 1
			cargoPlan[mark] = append([]string{string(crate)}, cargoPlan[mark]...)
		}
	}

	for _, moveRaw := range moveRawPlan {
		instructions := strings.Split(moveRaw, " ")
		move, _ := strconv.Atoi(instructions[1])
		from, _ := strconv.Atoi(instructions[3])
		to, _ := strconv.Atoi(instructions[5])

		fromLen := len(cargoPlan[from])

		cargoPlan[to] = append(cargoPlan[to], reverse(cargoPlan[from][fromLen-move:])...)
		cargoPlan[from] = cargoPlan[from][:fromLen-move]
	}

	echoTopCrates(cargoPlan)
}

func echoTopCrates(cargo map[int][]string) {
	var solution string

	for i := 1; i < len(cargo)+1; i++ {
		stack := cargo[i]
		solution += stack[len(stack)-1]
	}

	fmt.Printf("Solution: %v\n", solution)
}

func reverse(array []string) []string {
	for i := 0; i < len(array)/2; i++ {
		j := len(array) - i - 1
		array[i], array[j] = array[j], array[i]
	}

	return array
}
