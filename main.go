package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce-go/data"
	"ecommerce-go/handlers"
	"ecommerce-go/interfaces"
	"ecommerce-go/services"
)

func main() {
	// Se declaran los servicios por medio de interfaces para mantener bajo acoplamiento.
	var productoService interfaces.ProductoManager = services.ProductoService{}
	var clienteService interfaces.ClienteManager = services.ClienteService{}
	var pedidoService interfaces.PedidoManager = services.PedidoService{}

	store := data.NewStore(productoService, clienteService)
	app := handlers.NewApp(store, productoService, clienteService, pedidoService)
	mux := http.NewServeMux()
	app.RegistrarRutas(mux)
	puerto := ":8080"
	fmt.Println("====================================")
	fmt.Println(" SISTEMA DE GESTIÓN E-COMMERCE WEB")
	fmt.Println("====================================")
	fmt.Println("Página web: http://localhost" + puerto)
	fmt.Println("API JSON:   http://localhost" + puerto + "/api")
	fmt.Println("Presione Ctrl + C para detener el servidor.")
	fmt.Println("====================================")
	log.Fatal(http.ListenAndServe(puerto, mux))
}
