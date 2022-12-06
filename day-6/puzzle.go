package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(rawInput))

	packetBuf, startOfPacket := getIndexOfNDifferentElements(input, 4)
	fmt.Printf("%v - first pckt mark after: %v chars\n", startOfPacket, packetBuf)

	messageBuf, startOfMessage := getIndexOfNDifferentElements(input, 14)
	fmt.Printf("%v - first msg mark after: %v chars\n", startOfMessage, messageBuf)
}

func allDifferentElements(array string) bool {
	for i := 0; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[i] == array[j] {
				return false
			}
		}
	}

	return true
}

func getIndexOfNDifferentElements(input string, n int) (string, int) {
	for i := n; i < len(input); i++ {
		buffer := input[i-n : i]

		if allDifferentElements(buffer) {
			return buffer, i
		}
	}

	return "", len(input)
}
