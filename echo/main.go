package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/", saludar)
	e.GET("/dividir", dividir)

	// Grupo de rutas
	// e.POST("/personas/crear", crear)
	// e.GET("/personas/consultar", consultar)
	// e.PUT("/personas/actualizar", actualizar)
	// e.DELETE("/personas/borrar", borrar)
	persons := e.Group("/personas")
	persons.Use(middlewareLogPersonas)
	persons.POST("", crear)
	persons.GET("/:id", consultar)
	persons.PUT("/:id", actualizar)
	persons.DELETE("/:id", borrar)

	e.Start(":8080")
}

func saludar(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"saludo": "Hola Mundo"})
}

func dividir(c echo.Context) error {
	d := c.QueryParam("d")
	f, _ := strconv.Atoi(d)
	if f == 0 {
		return c.String(http.StatusBadRequest, "El valor no puede ser cero")
	}

	r := 3000 / f

	return c.String(http.StatusOK, strconv.Itoa(r))
}

func crear(c echo.Context) error {
	return c.String(http.StatusOK, "creado")
}

func actualizar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "actualizado "+id)
}

func borrar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "borrado "+id)
}

func consultar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "consultado "+id)
}

func middlewareLogPersonas(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Petición hecha a /personas")
		return f(c)
	}
}
