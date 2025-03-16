package main

import (
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"sena-crud/handlers"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/agregar-cliente", handlers.AgregarCliente)
	http.HandleFunc("/clientes", handlers.Clientes)
	http.HandleFunc("/borrar-cliente", handlers.BorrarCliente)
	http.HandleFunc("/editar-cliente", handlers.EditarCliente)
	http.HandleFunc("/actualizar-cliente", handlers.ActualizarCliente)

	http.HandleFunc("/lista-productos", handlers.ProductHome)
	http.HandleFunc("/agregar-producto", handlers.AgregarProducto)
	http.HandleFunc("/productos", handlers.Productos)

	log.Println("Server Running...")
	http.ListenAndServe(":8080", nil)
}
