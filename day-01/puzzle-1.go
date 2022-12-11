package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
  file, err := os.Open("calories-by-elf")

  if err != nil {
    fmt.Println(err)
  }

  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)

  maxCalories := 0
  var calories int

  for fileScanner.Scan() {
    lineInput := fileScanner.Text()

    if lineInput != ""  {
      inputCalories, err := strconv.Atoi(lineInput)
      if err == nil {
        calories += inputCalories
      }
    } else  {
      if maxCalories < calories {
        maxCalories = calories
      }

      calories = 0
    }
  }

  if maxCalories < calories {
    maxCalories = calories
  }
  calories = 0

  file.Close()

  fmt.Printf("%v", maxCalories)
}

