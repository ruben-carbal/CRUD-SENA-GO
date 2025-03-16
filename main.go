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
	http.HandleFunc("/form-registro", handlers.FormRegistro)
	http.HandleFunc("/registro", handlers.Register)

	http.HandleFunc("/lista-clientes", handlers.ClienteHome)
	http.HandleFunc("/agregar-cliente", handlers.AgregarCliente)
	http.HandleFunc("/clientes", handlers.Clientes)
	http.HandleFunc("/borrar-cliente", handlers.BorrarCliente)
	http.HandleFunc("/editar-cliente", handlers.EditarCliente)
	http.HandleFunc("/actualizar-cliente", handlers.ActualizarCliente)

	http.HandleFunc("/lista-productos", handlers.ProductHome)
	http.HandleFunc("/agregar-producto", handlers.AgregarProducto)
	http.HandleFunc("/productos", handlers.Productos)
	http.HandleFunc("/borrar-producto", handlers.BorrarProducto)
	http.HandleFunc("/editar-producto", handlers.EditarProducto)
	http.HandleFunc("/actualizar-producto", handlers.ActualizarProducto)

	log.Println("Server Running...")
	http.ListenAndServe(":8080", nil)
}
