package services

import (
	"fmt"
	"time"
	"sistema-ecommerce-go/models"
)

func CalcularSubtotal(productos []models.Producto) float64 {
	subtotal := 0.0
	for _, producto := range productos {
		subtotal += producto.Precio
	}
	return subtotal
}

func CalcularIVA(subtotal float64) float64 {
	return subtotal * 0.15
}

func CalcularTotal(subtotal float64, iva float64) float64 {
	return subtotal + iva
}

func CrearPedido(id int, cliente models.Cliente, productos []models.Producto) models.Pedido {
	subtotal := CalcularSubtotal(productos)
	iva := CalcularIVA(subtotal)
	total := CalcularTotal(subtotal, iva)

	return models.Pedido{
		ID: id, Cliente: cliente, Productos: productos,
		Fecha: time.Now(), Subtotal: subtotal, IVA: iva, Total: total,
	}
}

func MostrarPedido(pedido models.Pedido) {
	fmt.Println("\n===== RESUMEN DEL PEDIDO =====")
	fmt.Printf("Pedido N°: %d\n", pedido.ID)
	fmt.Printf("Cliente: %s\n", pedido.Cliente.Nombre)
	fmt.Printf("Correo: %s\n", pedido.Cliente.Correo)
	fmt.Printf("Fecha: %s\n", pedido.Fecha.Format("02/01/2006 15:04"))
	fmt.Println("\nProductos comprados:")
	for _, producto := range pedido.Productos {
		fmt.Printf("- %s: $%.2f\n", producto.Nombre, producto.Precio)
	}
	fmt.Printf("\nSubtotal: $%.2f\n", pedido.Subtotal)
	fmt.Printf("IVA 15%%: $%.2f\n", pedido.IVA)
	fmt.Printf("Total a pagar: $%.2f\n", pedido.Total)
}
