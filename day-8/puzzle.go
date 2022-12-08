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

	var trees [][]int
	rows := strings.Split(strings.TrimSpace(string(file)), "\n")

	for i, row := range rows {
		trees = append(trees, []int{})
		for _, col := range row {
			if n, err := strconv.Atoi(string(col)); err == nil {
				trees[i] = append(trees[i], n)
			}
		}
	}

	count := 4 + 2*(len(trees)-2) + 2*(len(trees[0])-2)
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			tree := trees[i][j]

			isVisibleX := isMaxVertical(tree, trees[i][:j]) || isMaxVertical(tree, trees[i][j+1:])
			isVisibleY := isMaxHorizontal(tree, j, trees[:i]) || isMaxHorizontal(tree, j, trees[i+1:])

			if isVisibleX || isVisibleY {
				count++
			}
		}
	}

	fmt.Printf("Puzzle 1. - Visible count: %v\n", count)

	var maxScenicScore, scenicScore int
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			scenicScore = countSmaller(i, j, trees)
			if maxScenicScore < scenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Printf("Puzzle 2. - Scenic score: %v\n", maxScenicScore)
}

func countSmaller(x int, y int, nss [][]int) int {
	var w, a, s, d int

	for i := x - 1; i >= 0; i-- {
		w++
		if nss[x][y] <= nss[i][y] {
			break
		}
	}

	for i := x + 1; i < len(nss); i++ {
		s++
		if nss[x][y] <= nss[i][y] {
			break
		}
	}

	for j := y - 1; j >= 0; j-- {
		a++
		if nss[x][y] <= nss[x][j] {
			break
		}
	}

	for j := y + 1; j < len(nss[0]); j++ {
		d++
		if nss[x][y] <= nss[x][j] {
			break
		}
	}

	return w * a * s * d
}

func isMaxVertical(checking int, ns []int) bool {
	for _, n := range ns {
		if checking <= n {
			return false
		}
	}

	return true
}

func isMaxHorizontal(checking int, col int, nss [][]int) bool {
	for _, ns := range nss {
		if checking <= ns[col] {
			return false
		}
	}

	return true
}
