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

//func BorrarProducto(w http.ResponseWriter, r *http.Request) {
//	idproducto := r.URL.Query().Get("id")
//
//	conexion := db.ConexionDB()
//	borrar, err := conexion.Prepare("DELETE FROM productos WHERE id=?")
//
//	if err != nil {
//		panic(err.Error())
//	}
//
//	borrar.Exec(idproducto)
//	http.Redirect(w, r, "/", 301)
//
//}
//
//func EditarProducto(w http.ResponseWriter, r *http.Request) {
//	idproducto := r.URL.Query().Get("id")
//
//	conexion := db.ConexionDB()
//	registro, err := conexion.Query("SELECT * FROM productos WHERE id=?", idCliente)
//
//	producto := Cliente{}
//
//	for registro.Next() {
//		var id int
//		var nombre, correo, telefono, direccion string
//		err = registro.Scan(&id, &nombre, &correo, &telefono, &direccion)
//
//		if err != nil {
//			panic(err.Error())
//		}
//
//		producto.Id = id
//		producto.Nombre = nombre
//		producto.Correo = correo
//		producto.Telefono = telefono
//		producto.Direccion = direccion
//	}
//
//	plantillas.ExecuteTemplate(w, "editarproducto", cliente)
//}
//
//func ActualizarProducto(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "POST" {
//		id := r.FormValue("id")
//		nombre := r.FormValue("name")
//		correo := r.FormValue("correo")
//		telefono := r.FormValue("telefono")
//		direccion := r.FormValue("direccion")
//
//		conexion := db.ConexionDB()
//		modificar, err := conexion.Prepare("UPDATE productos SET nombre=?, correo=?, telefono=?, direccion=? WHERE id=?")
//
//		if err != nil {
//			panic(err.Error())
//		}
//
//		modificar.Exec(nombre, correo, telefono, direccion, id)
//		http.Redirect(w, r, "/", 301)
//	}
//}
