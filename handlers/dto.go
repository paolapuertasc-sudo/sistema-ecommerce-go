package handlers

import (
	"ecommerce-go/models"
	"time"
)

type ProductoDTO struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock  int     `json:"stock"`
}
type ClienteDTO struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
}
type ItemPedidoDTO struct {
	Producto ProductoDTO `json:"producto"`
	Cantidad int         `json:"cantidad"`
	Subtotal float64     `json:"subtotal"`
}
type PedidoDTO struct {
	ID       int             `json:"id"`
	Cliente  ClienteDTO      `json:"cliente"`
	Items    []ItemPedidoDTO `json:"items"`
	Fecha    string          `json:"fecha"`
	Subtotal float64         `json:"subtotal"`
	IVA      float64         `json:"iva"`
	Total    float64         `json:"total"`
}

func productoToDTO(producto models.Producto) ProductoDTO {
	return ProductoDTO{ID: producto.ID(), Nombre: producto.Nombre(), Precio: producto.Precio(), Stock: producto.Stock()}
}
func clienteToDTO(cliente models.Cliente) ClienteDTO {
	return ClienteDTO{ID: cliente.ID(), Nombre: cliente.Nombre(), Correo: cliente.Correo()}
}
func pedidoToDTO(pedido models.Pedido) PedidoDTO {
	items := []ItemPedidoDTO{}
	for _, item := range pedido.Items() {
		items = append(items, ItemPedidoDTO{Producto: productoToDTO(item.Producto()), Cantidad: item.Cantidad(), Subtotal: item.Subtotal()})
	}
	return PedidoDTO{ID: pedido.ID(), Cliente: clienteToDTO(pedido.Cliente()), Items: items, Fecha: pedido.Fecha().Format(time.RFC3339), Subtotal: pedido.Subtotal(), IVA: pedido.IVA(), Total: pedido.Total()}
}
