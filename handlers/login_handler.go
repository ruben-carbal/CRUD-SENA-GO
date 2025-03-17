package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"sena-crud/utils"
	"time"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

var users = map[string]Login{}

var AuthError = errors.New("No autorizado")

func authorize(r *http.Request) error {
	username := r.FormValue("username")
	user, ok := users[username]

	if !ok {
		return AuthError
	}

	// obtener el token de las cookies
	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		return AuthError
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken || csrf == "" {
		return AuthError
	}

	return nil
}

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

func LoginFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, ok := users[username]
		if !ok || utils.CheckPasswordHash(password, user.HashedPassword) {
			er := http.StatusBadRequest
			http.Error(w, "Usuario o contraseña incorrecto", er)
			return
		}

		sessionToken := utils.GenerateToken(32)
		csrfToken := utils.GenerateToken(32)

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    csrfToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: false,
		})

		// guardar el token en la base de datos
		user.SessionToken = sessionToken
		user.CSRFToken = csrfToken
		users[username] = user

		http.Redirect(w, r, "/lista-clientes", 301)
	}
}

func Protected(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := authorize(r); err != nil {
			er := http.StatusUnauthorized
			http.Error(w, "No autorizado", er)
			return
		}

		username := r.FormValue("username")
		fmt.Fprintf(w, "Validación exitosa! bienvenido, %s", username)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if err := authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "No autorizado", er)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	//borrar los tokens de la base de datos
	username := r.FormValue("username")
	user, _ := users[username]
	user.SessionToken = ""
	user.CSRFToken = ""
	users[username] = user

	fmt.Fprintln(w, "logout exitoso!")
}
