package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	workingDir := []string{}
	filesystem := make(map[string]int)

	cmdIOs := parseFile("test-input")

	for _, cmdIO := range cmdIOs {
		cmd := strings.Split(strings.TrimSpace(cmdIO), " ")

		fileSize, err := strconv.Atoi(cmd[0])
		if err == nil {
			recalcWorkingDirSize(fileSize, workingDir, filesystem)
		}

		if cmd[1] == "cd" {

			workingDir = changeDirectory(cmd[2], workingDir)
			if cmd[2] != ".." && cmd[2] != "/" {

				filesystem[printWorkingDir(workingDir)] = 0
			}
		}
	}

	sizeSum, sizeTotal := 0, 0
	for dir, size := range filesystem {
		sizeTotal += size
		if size < 100000 {
			sizeSum += size
		}
		fmt.Printf("%v - %v\n", dir, size)
	}

	fmt.Printf("Puzzle 1: %v\n", sizeSum)

	totalSpace := 70000000
	//requiredSpace := 30000000

	fmt.Printf("\n\nUsed space: %v\nFree space: %v\n\n", filesystem["/"], totalSpace-sizeTotal)
}

func parseFile(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSpace(string(file)), "\n")
}

func changeDirectory(directory string, workingDir []string) []string {
	if directory == ".." {
		return workingDir[:len(workingDir)-1]
	}

	if directory == "/" {
		return []string{"/"}
	}

	return append(workingDir, directory)
}

func recalcWorkingDirSize(addedSize int, currentWorkingDir []string, filesystem map[string]int) {
	for i := 1; i < len(currentWorkingDir); i++ {
		path := printWorkingDir(currentWorkingDir[:i+1])
		filesystem[path] += addedSize
	}
}

func printWorkingDir(workingDir []string) string {
	return workingDir[0] + strings.Join(workingDir[1:], "/")
}
