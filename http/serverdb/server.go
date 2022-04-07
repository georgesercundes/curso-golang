package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("Executando...")
	port := fmt.Sprintf(":%s", os.Getenv("DB_PORT"))
	fmt.Println(port)
	log.Fatal(http.ListenAndServe(port, nil))
}
