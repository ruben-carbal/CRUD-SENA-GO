package handlers

import (
	//"fmt"
	"net/http"
	"sena-crud/utils"
	//"time"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

var users = map[string]Login{}

func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Home")
	plantillas.ExecuteTemplate(w, "userLogin", nil)
}

func FormRegistro(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "registro", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if len(password) < 4 {
			er := http.StatusNotAcceptable
			http.Error(w, "Tu contraseña es muy corta", er)
			return
		}

		if _, ok := users[username]; ok {
			// en go podemos obtener dos valores de un map valor, ok : users["clave"]. El segundo valor es un booleano
			// en este caso el guión bajo indica que ignoramos el valor
			er := http.StatusConflict
			http.Error(w, "El usuario ya existe", er)
			return
		}

		hashedPassword, _ := utils.HashPassword(password)
		users[username] = Login{
			HashedPassword: hashedPassword,
		}

		//fmt.Fprintln(w, "User registered succesfully!")
		http.Redirect(w, r, "/lista-clientes", 301)
	}
}

func LoginFunc() (w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, ok := users[username]
		if !ok {
			er := http.StatusBadRequest
			http.Error(w, "Usuario o contraseña incorrecto", er)
			return
		}
	}
}

//func Logout() (w http.ResponseWriter, r *http.Request) {}
//
//func Protected() (w http.ResponseWriter, r *http.Request) {}
