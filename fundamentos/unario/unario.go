package main

import "fmt"

func main() {
  x := 1
  y := 2

  // Apenas postfix
  x++
  fmt.Println(x)

  y--
  fmt.Println(y)

  // fmt.Println(x == y--)
}