package main

import (
	//"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"text/template"
)

func conexionDB() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contraseña := ""
	Nombre := "datos"

	db, err := sql.Open(Driver, Usuario+":"+Contraseña+"@tcp(127.0.0.1)/"+Nombre)

	if err != nil {
		panic(err.Error())
	}

	return db
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/agregar-cliente", AgregarCliente)
	http.HandleFunc("/clientes", Clientes)

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
	conexion := conexionDB()
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

	log.Println("clientes: ", arrayCliente)

	plantillas.ExecuteTemplate(w, "clientes", arrayCliente)
}

func AgregarCliente(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "agregarCliente", nil)
}

func Clientes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("name")
		correo := r.FormValue("correo")
		telefono := r.FormValue("telefono")
		direccion := r.FormValue("direccion")

		conexion := conexionDB()
		insertar, err := conexion.Prepare("INSERT INTO clientes(nombre, correo, telefono, direccion) VALUES(?, ?, ?, ?)")

		if err != nil {
			panic(err.Error())
		}

		insertar.Exec(nombre, correo, telefono, direccion)
		http.Redirect(w, r, "/", 301)
	}
}
