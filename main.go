package main

import (
	//"fmt"
	"log"
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Home)
	log.Println("Server Running...")
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hola")
	plantillas.ExecuteTemplate(w, "clientes", nil)
}
