package main

import (
	"sistema-ecommerce-go/models"
	"sistema-ecommerce-go/services"
)

func main() {
	productos := []models.Producto{
		services.CrearProducto(1, "Laptop HP", 650.00, 5),
		services.CrearProducto(2, "Mouse inalámbrico", 15.00, 20),
		services.CrearProducto(3, "Teclado mecánico", 45.00, 10),
		services.CrearProducto(4, "Audífonos Bluetooth", 30.00, 8),
	}

	services.ListarProductos(productos)

	cliente := services.CrearCliente(1, "Paola Puertas", "paola@example.com")
	productosSeleccionados := []models.Producto{productos[0], productos[1], productos[3]}
	pedido := services.CrearPedido(1, cliente, productosSeleccionados)

	services.MostrarPedido(pedido)
}
