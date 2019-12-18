package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

var ADD int = 1
var MULTI int = 2
var EXIT int = 99

func performAddition(input []int, i int) {
  positionOne := input[i+1]
  positionTwo := input[i+2]
  replace := input[i+3]
  input[replace] = input[positionOne] + input[positionTwo]
}

func performMultiplication(input []int, i int) {
  positionOne := input[i+1]
  positionTwo := input[i+2]
  replace := input[i+3]
  input[replace] = input[positionOne] * input[positionTwo]
}

func parseIntCode(input []int) []int {
  for i := 0; i < len(input);  i += 4 {
    opcode := input[i]
    if opcode == ADD {
//      fmt.Println(opcode)
      performAddition(input, i)
    } else if opcode == MULTI {
//      fmt.Println(opcode)
      performMultiplication(input, i)
    } else if opcode == EXIT {
//      fmt.Println(opcode)
      fmt.Println("Exiting program!")
      break
    }
  }
  return input
}

func getInput() []int {
  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    fmt.Println("File reading error", err)
    return nil
  }
  raw := strings.Split(string(data), ",")
  arr := make([]int, len(raw))
  for i := range arr {
    arr[i], _ = strconv.Atoi(raw[i])
  }
  fmt.Println(arr)
  return arr
}

func main() {
  for x := 1; x < 100; x++ {
    for y := 1; y < 100; y++ {
      temp := getInput()
      temp[1] = x
      temp[2] = y
      temp = parseIntCode(temp)
      fmt.Println(x, y, temp[0])
      if temp[0] == 19690720 {
        fmt.Println("noun", x)
        fmt.Println("verb", y)

        fmt.Println("Answer is", 100 * x + y)
        return
      }
    }
  }
}
