package main

import (
	"fmt"
	"curso_golang/html"
)


func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar (entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.uol.com.br/", "https://www.youtube.com"),
	)
	fmt.Println("1º)", <-c)
	fmt.Println("2º)", <-c)
	fmt.Println("3º)", <-c)
	fmt.Println("4º)", <-c)
}