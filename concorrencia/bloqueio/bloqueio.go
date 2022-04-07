package main

import (
	"fmt"
	"time"	
)

func rotina (c chan int) {
	time.Sleep(time.Second)
	c <- 1 // Operação bloqueante
	fmt.Println("Lido")
}

func main() {
	c := make(chan int)

	go rotina(c)

	fmt.Println(<-c) // Operação bloqueante
	fmt.Println("Só depois que foi lido")
	fmt.Println(<-c) // Deadlock
	fmt.Println("Fim")
	
}