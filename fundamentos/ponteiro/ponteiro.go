package main

import "fmt"

func main() {
  i := 1

  var p *int = nil

  p = &i // Pegando o endereço da variável
  *p++ // Desreferenciando (pegando o valor)
  i++

  // Go não tem aritmética de ponteiros
  // p++

  fmt.Println(p, *p, i, &i)

}