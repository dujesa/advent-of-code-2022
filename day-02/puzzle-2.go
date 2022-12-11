package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	games := [][]string{}
	games = parseStrategy(games)
	points := 0

	for _, game := range games {
		points += calculatePoints(game)
	}

	fmt.Printf("Total score: %v", points)
}

func calculatePoints(game []string) int {
	gamePoints := 0
	oponentMove := game[0]
	playerMove := game[2]

	switch playerMove {
	case "X":
		gamePoints = getScoreForLoseMove(oponentMove)
	case "Y":
		gamePoints = getScoreForDrawMove(oponentMove)
	case "Z":
		gamePoints = getScoreForWinMove(oponentMove)
	}

	fmt.Printf("Oponent: %s, Me: %v, Pts: %v \n", oponentMove, playerMove, gamePoints)
	return gamePoints
}

func getScoreForLoseMove(oponentMove string) int {
	gamePoints := 0

	switch oponentMove {
	case "A":
		return 3 + gamePoints
	case "B":
		return 1 + gamePoints
	case "C":
		return 2 + gamePoints
	}

	return gamePoints
}

func getScoreForDrawMove(oponentMove string) int {
	gamePoints := 3

	switch oponentMove {
	case "A":
		return 1 + gamePoints
	case "B":
		return 2 + gamePoints
	case "C":
		return 3 + gamePoints
	}

	return gamePoints
}

func getScoreForWinMove(oponentMove string) int {
	gamePoints := 6

	switch oponentMove {
	case "A":
		return 2 + gamePoints
	case "B":
		return 3 + gamePoints
	case "C":
		return 1 + gamePoints
	}

	return gamePoints
}
func parseStrategy(games [][]string) [][]string {
	var input string

	file, err := os.Open("rock-paper-scissors-games")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = scanner.Text()
		lineInputs := strings.Split(input, "")
		games = append(games, lineInputs)
	}

	return games
}
