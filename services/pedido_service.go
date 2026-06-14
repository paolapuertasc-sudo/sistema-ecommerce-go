package services

import (
	"errors"
	"fmt"

	"ecommerce-go/models"
	"ecommerce-go/utils"
)

const ivaPorcentaje = 0.15

type PedidoService struct{}

func (s PedidoService) CrearItem(producto models.Producto, cantidad int) (models.ItemPedido, models.Producto, error) {
	if err := utils.ValidarCantidad(cantidad); err != nil {
		return models.ItemPedido{}, models.Producto{}, err
	}

	// Antes de agregar el producto al pedido, se valida y descuenta el stock.
	// El método DescontarStock no modifica el producto original directamente,
	// sino que devuelve una nueva versión del producto con el stock actualizado.
	productoActualizado, err := producto.DescontarStock(cantidad)
	if err != nil {
		return models.ItemPedido{}, models.Producto{}, err
	}

	item := models.NuevoItemPedido(producto, cantidad)
	return item, productoActualizado, nil
}

func (s PedidoService) CrearPedido(id int, cliente models.Cliente, items []models.ItemPedido) (models.Pedido, error) {
	if !cliente.EsValido() {
		return models.Pedido{}, errors.New("el cliente no es válido para crear un pedido")
	}

	if len(items) == 0 {
		return models.Pedido{}, errors.New("no se puede crear un pedido sin productos")
	}

	subtotal := s.CalcularSubtotal(items)
	iva := s.CalcularIVA(subtotal)
	total := s.CalcularTotal(subtotal, iva)

	return models.NuevoPedido(id, cliente, items, subtotal, iva, total), nil
}

func (s PedidoService) CalcularSubtotal(items []models.ItemPedido) float64 {
	// Se recorre cada item del pedido y se suma su subtotal.
	// Cada subtotal se calcula multiplicando precio del producto por cantidad comprada.
	subtotal := 0.0

	for _, item := range items {
		subtotal += item.Subtotal()
	}

	return subtotal
}

func (s PedidoService) CalcularIVA(subtotal float64) float64 {
	return subtotal * ivaPorcentaje
}

func (s PedidoService) CalcularTotal(subtotal float64, iva float64) float64 {
	return subtotal + iva
}

func (s PedidoService) MostrarPedido(pedido models.Pedido) {
	fmt.Println("\n===== RESUMEN DEL PEDIDO =====")
	fmt.Printf("Pedido N°: %d\n", pedido.ID())
	fmt.Printf("Cliente: %s\n", pedido.Cliente().Nombre())
	fmt.Printf("Correo: %s\n", pedido.Cliente().Correo())
	fmt.Printf("Fecha: %s\n", pedido.Fecha().Format("02/01/2006 15:04"))

	fmt.Println("\nProductos comprados:")
	for _, item := range pedido.Items() {
		fmt.Printf("- %s | Cantidad: %d | Subtotal: $%.2f\n",
			item.Producto().Nombre(),
			item.Cantidad(),
			item.Subtotal(),
		)
	}

	fmt.Printf("\nSubtotal: $%.2f\n", pedido.Subtotal())
	fmt.Printf("IVA 15%%: $%.2f\n", pedido.IVA())
	fmt.Printf("Total a pagar: $%.2f\n", pedido.Total())
}

func (s PedidoService) ListarPedidos(pedidos []models.Pedido) {
	fmt.Println("\n===== HISTORIAL DE PEDIDOS =====")

	if len(pedidos) == 0 {
		fmt.Println("No existen pedidos registrados.")
		return
	}

	for _, pedido := range pedidos {
		fmt.Printf("Pedido N°: %d | Cliente: %s | Total: $%.2f | Fecha: %s\n",
			pedido.ID(),
			pedido.Cliente().Nombre(),
			pedido.Total(),
			pedido.Fecha().Format("02/01/2006 15:04"),
		)
	}
}
