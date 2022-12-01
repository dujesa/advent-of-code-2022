package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "sort"
)

func main() {
  file, err := os.Open("calories-by-elf")

  if err != nil {
    fmt.Println(err)
  }

  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)

  var calories, topThreeSum int
  var elvesCalories []int

  for fileScanner.Scan() {
    lineInput := fileScanner.Text()

    if lineInput != ""  {
      inputCalories, err := strconv.Atoi(lineInput)
      if err == nil {
        calories += inputCalories
      }
    } else  {
      elvesCalories = append(elvesCalories, calories)
      calories = 0
    }
  }

  elvesCalories = append(elvesCalories, calories)
  calories = 0

  file.Close()

  sort.Slice(elvesCalories, func(i, j int) bool {
    return elvesCalories[i] > elvesCalories[j]
  })

  for i := 0; i < 3; i++ {
    topThreeSum += elvesCalories[i]
  }

  fmt.Printf("%v", topThreeSum)
}

