package goaccessql

import (
	"database/sql"
	"fmt"

	//Para conexion con MariaDB
	_ "github.com/go-sql-driver/mysql"
)

//SetupBD Conexion BD y operaciones con la base de datos
type SetupBD struct {
	Servidor string `json:"servidor,omitempty"`
	Puerto   string `json:"puerto,omitempty"`
	BD       string `json:"bd,omitempty"`
	Usuario  string `json:"usuario,omitempty"`
	Clave    string `json:"clave,omitempty"`
}

//FilaSQL estructura para mapeo de datos obtenidos
type FilaSQL map[string]string

//Get Obtiene la conexion al servidor de base de datos
func (s SetupBD) get() (db *sql.DB, err error) {
	//user:password@tcp(server:port)/database?tls=false&autocommit=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", s.Usuario, s.Clave, s.Servidor, s.Puerto, s.BD)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	return db, err
}

//GetEstado Obtiene el estado de conexion con la base de datos
func (s SetupBD) GetEstado() bool {
	_, err := s.get()
	return err == nil
}

//Select Operaciones de consulta y obtencion de datos
func (s SetupBD) Select(query string, valores ...interface{}) (tabla []FilaSQL, err error) {
	bd, err := s.get()
	if err != nil {
		return nil, err
	}
	defer bd.Close()

	stmt, err := bd.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	filas, err := stmt.Query(valores...)

	if err != nil {
		return nil, err
	}
	defer filas.Close()

	columns, err := filas.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	for c := range values {
		scanArgs[c] = &values[c]
	}

	for filas.Next() {
		err = filas.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		fila := make(FilaSQL, len(columns))
		for i, val := range values {
			if val == nil {
				fila[columns[i]] = "--"
			} else {
				fila[columns[i]] = string(val)
			}
		}
		tabla = append(tabla, fila)
	}
	return tabla, err
}

//Query Operaciones de insercion, actualizacion y eliminacion de datos
func (s SetupBD) Query(query string, values ...interface{}) (err error) {
	bd, err := s.get()
	if err != nil {
		return err
	}
	defer bd.Close()

	stmt, err := bd.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}
	return err
}
