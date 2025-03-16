package handlers

import (
	"fmt"
	"net/http"
	"sena-crud/db"
)

type Producto struct {
	IdProd     int
	NombreProd string
	Categoria  string
	Precio     string
	Stock      int
}

func ProductHome(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hola")
	conexion := db.ConexionDB()
	registros, err := conexion.Query("SELECT * FROM productos")

	if err != nil {
		panic(err.Error())
	}

	producto := Producto{}
	arrayProductos := []Producto{}

	for registros.Next() {
		var codigo, stock int
		var nombre, categoria string
		var precio float64

		err = registros.Scan(&codigo, &nombre, &categoria, &precio, &stock)
		if err != nil {
			panic(err.Error())
		}

		producto.IdProd = codigo
		producto.NombreProd = nombre
		producto.Categoria = categoria
		producto.Precio = fmt.Sprintf("%.2f", precio)
		producto.Stock = stock

		arrayProductos = append(arrayProductos, producto)
	}

	plantillas.ExecuteTemplate(w, "productos", arrayProductos)
}

func AgregarProducto(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "agregarProducto", nil)
}

func Productos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		categoria := r.FormValue("categoria")
		precio := r.FormValue("precio")
		stock := r.FormValue("stock")

		conexion := db.ConexionDB()
		insertar, err := conexion.Prepare("INSERT INTO productos(nombre, categoria, precio, stock) VALUES(?, ?, ?, ?)")

		if err != nil {
			panic(err.Error())
		}

		insertar.Exec(nombre, categoria, precio, stock)
		http.Redirect(w, r, "/lista-productos", 301)
	}
}

func BorrarProducto(w http.ResponseWriter, r *http.Request) {
	idProducto := r.URL.Query().Get("codigo")

	conexion := db.ConexionDB()
	borrar, err := conexion.Prepare("DELETE FROM productos WHERE codigo=?")

	if err != nil {
		panic(err.Error())
	}

	borrar.Exec(idProducto)
	http.Redirect(w, r, "/lista-productos", 301)

}

func EditarProducto(w http.ResponseWriter, r *http.Request) {
	idProducto := r.URL.Query().Get("codigo")

	conexion := db.ConexionDB()
	registro, err := conexion.Query("SELECT * FROM productos WHERE codigo=?", idProducto)

	producto := Producto{}

	for registro.Next() {
		var codigo, stock int
		var nombre, categoria string
		var precio float64

		err = registro.Scan(&codigo, &nombre, &categoria, &precio, &stock)

		if err != nil {
			panic(err.Error())
		}

		producto.IdProd = codigo
		producto.NombreProd = nombre
		producto.Categoria = categoria
		producto.Precio = fmt.Sprintf("%.2f", precio)
		producto.Stock = stock
	}

	plantillas.ExecuteTemplate(w, "editarProducto", producto)
}

func ActualizarProducto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		codigo := r.FormValue("codigo")
		nombre := r.FormValue("nombre")
		categoria := r.FormValue("categoria")
		precio := r.FormValue("precio")
		stock := r.FormValue("stock")

		conexion := db.ConexionDB()
		modificar, err := conexion.Prepare("UPDATE productos SET nombre=?, categoria=?, precio=?, stock=? WHERE codigo=?")

		if err != nil {
			panic(err.Error())
		}

		modificar.Exec(nombre, categoria, precio, stock, codigo)
		http.Redirect(w, r, "/lista-productos", 301)
	}
}
