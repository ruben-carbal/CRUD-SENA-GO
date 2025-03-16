package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConexionDB() (conexion *sql.DB) {
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
