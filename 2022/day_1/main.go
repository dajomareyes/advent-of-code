package main

import (
  "bufio"
  "log"
  "os"
  "strconv"
)

func solution() {
  file, err := os.Open("final.txt")
  // file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  sum := 0
  largestElfSum := 0

  for scanner.Scan() {
    line := scanner.Text()

    if line != "" {
      val, err := strconv.Atoi(line)
      if err != nil {
        log.Fatal(err)
      }
      sum += val
    }

    if (line == "") {
      log.Println(sum, largestElfSum)
      if sum > largestElfSum {
        largestElfSum = sum
      }
      sum = 0
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  log.Println("Largest sum is", largestElfSum)
}


func main() {
  log.Println("Day 1 - Start Problem")

  solution()
}
