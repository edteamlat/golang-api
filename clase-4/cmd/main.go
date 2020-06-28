package main

import "github.com/EDteam/golang-api/clase-4/funciones"

func execute(name string, f funciones.MyFunction) {
	f(name)
}

func main() {
	name := "Comunidad EDteam"
	execute(name, funciones.MiddlewareLog(funciones.Saludar))
	execute(name, funciones.MiddlewareLog(funciones.Despedirse))
}
