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

	heads := [][]int{{0, 0}}
	tails := [][]int{{0, 0}}
	moves := strings.Split(strings.TrimSpace(string(file)), "\n")

	head := heads[0]
	tail := tails[0]

	nextHead := head
	nextTail := tail

	tailMoves := make(map[string]int)

	for _, move := range moves {
		head = heads[len(heads)-1]
		tail = tails[len(tails)-1]

		nextHead = head

		direction := strings.Split(move, " ")
		steps, err := strconv.Atoi(direction[1])
		if err != nil {
			panic(err)
		}

		for steps >= 1 {
			if direction[0] == "R" {
				nextHead[0]++
			}
			if direction[0] == "L" {
				nextHead[0]--
			}
			if direction[0] == "U" {
				nextHead[1]++
			}
			if direction[0] == "D" {
				nextHead[1]--
			}
			heads = append(heads, nextHead)

			steps--

			nextTail = moveTail(tail, head)
			tails = append(heads, nextTail)

			nextTailKey := strconv.Itoa(nextTail[0]) + "-" + strconv.Itoa(nextTail[1])
			val, exists := tailMoves[nextTailKey]
			if exists {
				tailMoves[nextTailKey] = val + 1
			} else {
				tailMoves[nextTailKey] = 1
			}
		}
	}

	fmt.Printf("Visited coords: %v\n", len(tailMoves))
}

func moveTail(tail []int, head []int) []int {
	diffX := tail[0] - head[0]
	diffY := tail[1] - head[1]

	if tail[0] == head[0] {
		if diffY == -2 {
			tail[1]++
		} else if diffY == 2 {
			tail[1]--
		}
	} else if tail[1] == head[1] {
		if diffX == -2 {
			tail[0]++
		} else if diffX == 2 {
			tail[0]--
		}
	} else if diffX > 1 || diffX < -1 || diffY > 1 || diffY < -1 {
		if diffX > 0 {
			tail[0]--
		} else {
			tail[0]++
		}

		if diffY > 0 {
			tail[1]--
		} else {
			tail[1]++
		}
	}

	return tail
}
