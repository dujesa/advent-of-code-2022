package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	items        []int
	operator     string
	operand      int
	testOperand  int
	positiveNext int
	negativeNext int
	inspected    int
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	var ms []monkey
	monkeysData := strings.Split(strings.TrimSpace(string(file)), "\n\n")
	for _, monkeyData := range monkeysData {
		ms = append(ms, createMonkey(monkeyData))
	}

	for round := 0; round < 10000; round++ {
		for i, _ := range ms {
			playRound(i, ms, round)
		}
		if round == 19 {
			fmt.Printf("%v", ms[0])
		}
		if round == 0 || round == 19 || round == 999 || round == 1999 || round == 2999 || round == 3999 || round == 4999 || round == 5999 || round == 6999 || round == 7999 || round == 8999 || round == 9999 {
			fmt.Printf("\nRound %v: ", round+1)
			for _, m := range ms {
				fmt.Printf("%v ", m.inspected)
			}
		}
	}

	fmt.Printf("\nSolution 2: %v", productOfTopTwoInspects(ms))
}

func productOfTopTwoInspects(ms []monkey) int {
	topA, topB := ms[0].inspected, ms[1].inspected
	curr := 0

	for i := 2; i < len(ms); i++ {
		curr = ms[i].inspected
		if curr > topA {
			topB = topA
			topA = curr

			continue
		}

		if curr > topB {
			topB = curr
		}
	}

	return topA * topB
}

func playRound(i int, ms []monkey, round int) {
	worryLevel, next := 0, 0

	divider := 1
	for _, m := range ms {
		divider *= m.testOperand
	}

	for _, item := range ms[i].items {
		worryLevel = int(calc(item, ms[i].operand, ms[i].operator) % divider)

		next = test(worryLevel, ms[i].testOperand, ms[i].positiveNext, ms[i].negativeNext)
		ms[next].items = append(ms[next].items, worryLevel)

		ms[i].inspected++
	}

	ms[i].items = ms[i].items[:0]
}

func test(x int, y int, neg int, pos int) int {
	if x%y == 0 {
		return neg
	}

	return pos
}

func calc(x int, y int, op string) int {
	switch op {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	case "square":
		return x * x
	}
	return 0
}

func createMonkey(monkeyData string) monkey {
	lines := strings.Split(monkeyData, "\n")
	var m monkey

	for i, line := range lines {
		switch i {
		case 1:
			rawItems := strings.Split(strings.TrimSpace(line), ": ")
			items := strings.Split(rawItems[1], ", ")

			var item int
			for _, rawItem := range items {
				item, _ = strconv.Atoi(rawItem)
				m.items = append(m.items, item)
			}
		case 2:
			rawOperation := strings.Split(strings.TrimSpace(line), "old ")
			operation := strings.Split(rawOperation[1], " ")

			if operation[1] == "old" {
				m.operator = "square"
				break
			}

			m.operator = operation[0]
			m.operand, _ = strconv.Atoi(operation[1])
		case 3:
			rawTest := strings.Split(line, "by ")

			m.testOperand, _ = strconv.Atoi(rawTest[1])
		case 4:
			rawPositive := strings.Split(line, "monkey ")

			m.positiveNext, _ = strconv.Atoi(rawPositive[1])
		case 5:
			rawNegative := strings.Split(line, "monkey ")

			m.negativeNext, _ = strconv.Atoi(rawNegative[1])
		}
	}

	return m
}
