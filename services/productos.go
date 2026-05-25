package services

import (
	"fmt"
	"sistema-ecommerce-go/models"
)

func CrearProducto(id int, nombre string, precio float64, stock int) models.Producto {
	return models.Producto{ID: id, Nombre: nombre, Precio: precio, Stock: stock}
}

func ListarProductos(productos []models.Producto) {
	fmt.Println("===== CATÁLOGO DE PRODUCTOS =====")
	for _, producto := range productos {
		fmt.Printf("ID: %d | Producto: %s | Precio: $%.2f | Stock: %d\n",
			producto.ID, producto.Nombre, producto.Precio, producto.Stock)
	}
}
