package handlers

import (
	"net/http"
	"sena-crud/db"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

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

func AgregarCliente(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "agregarCliente", nil)
}

func Clientes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("name")
		correo := r.FormValue("correo")
		telefono := r.FormValue("telefono")
		direccion := r.FormValue("direccion")

		conexion := db.ConexionDB()
		insertar, err := conexion.Prepare("INSERT INTO clientes(nombre, correo, telefono, direccion) VALUES(?, ?, ?, ?)")

		if err != nil {
			panic(err.Error())
		}

		insertar.Exec(nombre, correo, telefono, direccion)
		http.Redirect(w, r, "/", 301)
	}
}

func BorrarCliente(w http.ResponseWriter, r *http.Request) {
	idCliente := r.URL.Query().Get("id")

	conexion := db.ConexionDB()
	borrar, err := conexion.Prepare("DELETE FROM clientes WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	borrar.Exec(idCliente)
	http.Redirect(w, r, "/", 301)

}

func EditarCliente(w http.ResponseWriter, r *http.Request) {
	idCliente := r.URL.Query().Get("id")

	conexion := db.ConexionDB()
	registro, err := conexion.Query("SELECT * FROM clientes WHERE id=?", idCliente)

	cliente := Cliente{}

	for registro.Next() {
		var id int
		var nombre, correo, telefono, direccion string
		err = registro.Scan(&id, &nombre, &correo, &telefono, &direccion)

		if err != nil {
			panic(err.Error())
		}

		cliente.Id = id
		cliente.Nombre = nombre
		cliente.Correo = correo
		cliente.Telefono = telefono
		cliente.Direccion = direccion
	}

	plantillas.ExecuteTemplate(w, "editarCliente", cliente)
}

func ActualizarCliente(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("name")
		correo := r.FormValue("correo")
		telefono := r.FormValue("telefono")
		direccion := r.FormValue("direccion")

		conexion := db.ConexionDB()
		modificar, err := conexion.Prepare("UPDATE clientes SET nombre=?, correo=?, telefono=?, direccion=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		modificar.Exec(nombre, correo, telefono, direccion, id)
		http.Redirect(w, r, "/", 301)
	}
}
