package main

import (
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"sena-crud/db"
	"sena-crud/handlers"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/agregar-cliente", handlers.AgregarCliente)
	http.HandleFunc("/clientes", handlers.Clientes)
	http.HandleFunc("/borrar-cliente", handlers.BorrarCliente)
	http.HandleFunc("/editar-cliente", handlers.EditarCliente)
	http.HandleFunc("/actualizar-cliente", handlers.ActualizarCliente)

	log.Println("Server Running...")
	http.ListenAndServe(":8080", nil)
}

type Cliente struct {
	Id        int
	Nombre    string
	Correo    string
	Telefono  string
	Direccion string
}

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hola")
	conexion := db.ConexionDB()
	registros, err := conexion.Query("SELECT * FROM clientes")

	if err != nil {
		panic(err.Error())
	}

	cliente := Cliente{}
	arrayCliente := []Cliente{}

	for registros.Next() {
		var id int
		var nombre, correo, telefono, direccion string

		err = registros.Scan(&id, &nombre, &correo, &telefono, &direccion)
		if err != nil {
			panic(err.Error())
		}

		cliente.Id = id
		cliente.Nombre = nombre
		cliente.Correo = correo
		cliente.Telefono = telefono
		cliente.Direccion = direccion

		arrayCliente = append(arrayCliente, cliente)
	}

	plantillas.ExecuteTemplate(w, "clientes", arrayCliente)
}
