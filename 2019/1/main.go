package main

import (
  "bufio"
  "flag"
  "fmt"
  "log"
  "math"
  "os"
  "strconv"
)

func calculateFuelCapacity(mass float64) float64 {
  value := (math.Floor(mass/3) - 2)
  // fmt.Println(mass, value)
  if value < 0 {
    return 0
  } else {
    return value + calculateFuelCapacity(value)
  }
}

func check(e error) {
  if e != nil {
    log.Fatal(e)
  }
}

func calculateSum() {
  sum := 0.0
  fptr := flag.String("fpath", "input.txt", "file path to read from")
  flag.Parse()

  f, err := os.Open(*fptr)
  check(err)
  defer func() {
    err = f.Close()
    check(err)
  }()
  s := bufio.NewScanner(f)
  for s.Scan() {
    n, err := strconv.ParseFloat(s.Text(), 64)
    check(err)
    sum += calculateFuelCapacity(n)
    // fmt.Println("mass", n, "|", "fuel", calculateFuelCapacity(n))
    // fmt.Println(s.Text())
  }
  err = s.Err()
  check(err)
  fmt.Println("total fuel capacity", int(sum))
}

func sanityCheck() {
  fmt.Println("mass = 12 and expect 2 ==", calculateFuelCapacity(12))
}

func main() {
	fmt.Println("Day 1 - Problem")
  calculateSum()
  // fmt.Println(calculateFuelCapacity(1969))
}
