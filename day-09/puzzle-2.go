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
	tss := [9][][]int{{{0, 0}}, {{0, 0}}, {{0, 0}}, {{0, 0}}, {{0, 0}}, {{0, 0}}, {{0, 0}}, {{0, 0}}, {{0, 0}}}
	moves := strings.Split(strings.TrimSpace(string(file)), "\n")

	head := heads[0]
	var tail []int

	nextHead := head
	nextTail := tail

	tailMoves := make(map[string]int)

	for _, move := range moves {
		head = heads[len(heads)-1]
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
			for i, ts := range tss {
				tail = ts[len(ts)-1]

				if i == 0 {
					nextTail = moveTail(tail, head)
				} else {
					tails := tss[i-1]
					nextTail = moveTail(tail, tails[len(tails)-1])
				}

				ts = append(heads, nextTail)
				if i == len(tss)-1 {
					nextTailKey := strconv.Itoa(nextTail[0]) + "-" + strconv.Itoa(nextTail[1])
					val, exists := tailMoves[nextTailKey]
					if exists {
						tailMoves[nextTailKey] = val + 1
					} else {
						tailMoves[nextTailKey] = 1
					}
				}
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
