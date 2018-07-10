# GoAccesSQL

Libreria para la conexion y operaciones con el motor de base de datos MariaDB y MySQL con GOLANG

## 1. Obtencion de libreria

```bash
go get github.com/MALDRU/GoAccesSQL
```

## 2. Conexion

```go
package main

import (
	"fmt"
	"github.com/MALDRU/GoAccesSQL"
)

func main() {
	context := goaccessql.SetupBD{
		Servidor: "servidor Base Datos",
		Puerto:   "puerto",
		BD:       "nombre base datos",
		Usuario:  "usuario Base de datos",
		Clave:    "contraseña del usuario",
	}

	if context.GetEstado() {
		fmt.Println("Conexion establecida correctamente")
	} else {
		fmt.Println("Error al establecer la conexion con la base de datos")
	}
}
```

## 3. Consultas SQL
### Consulta Basica
```go
	if !context.GetEstado() {
		panic("Error al establecer conexion con la base de datos")
	}

	tabla, err := context.Select("SELECT * FROM TABLA")
	if err != nil {
		panic(err)
	}

	fmt.Println(tabla)	
```
### Consulta con clausulas
```go
	if !context.GetEstado() {
		panic("Error al establecer conexion con la base de datos")
	}

	tabla, err := context.Select("SELECT * FROM TABLA WHERE id = ? AND nombre = ?", 25, "algun nombre")
	if err != nil {
		panic(err)
	}

	fmt.Println(tabla)
```

### Accediendo a valores retornados
```go
	if !context.GetEstado() {
		panic("Error al establecer conexion con la base de datos")
	}

	tabla, err := context.Select("SELECT * FROM TABLA")
	if err != nil {
		panic(err)
	}

	if len(tabla) > 0 {
		fmt.Println("Primer Registro: ", tabla[0]["nombre"])
	} else {
		fmt.Println("Sin registros")
	}
```
## 4. Inserciones, actualizaciones y eliminaciones SQL
```go
	if !context.GetEstado() {
		panic("Error al establecer conexion con la base de datos")
	}

	err := context.Query("INSERT INTO productos VALUES(null, ?, ?)", "PRODUCTO1", "330")
	if err != nil {
		panic(err)
	}
	fmt.Println("La operacion se realizo correctamente")
```
