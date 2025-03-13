package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	log.Println("Server Running...")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola")
}
