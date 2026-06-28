package models

import "time"

// ItemPedido representa un producto seleccionado dentro de un pedido.
type ItemPedido struct {
	producto Producto
	cantidad int
}

func NuevoItemPedido(producto Producto, cantidad int) ItemPedido {
	return ItemPedido{producto: producto, cantidad: cantidad}
}

func (i ItemPedido) Producto() Producto { return i.producto }
func (i ItemPedido) Cantidad() int      { return i.cantidad }
func (i ItemPedido) Subtotal() float64  { return i.producto.Precio() * float64(i.cantidad) }

// Pedido representa una venta realizada dentro del e-commerce.
type Pedido struct {
	id       int
	cliente  Cliente
	items    []ItemPedido
	fecha    time.Time
	subtotal float64
	iva      float64
	total    float64
}

func NuevoPedido(id int, cliente Cliente, items []ItemPedido, subtotal float64, iva float64, total float64) Pedido {
	return Pedido{id: id, cliente: cliente, items: items, fecha: time.Now(), subtotal: subtotal, iva: iva, total: total}
}

func (p Pedido) ID() int             { return p.id }
func (p Pedido) Cliente() Cliente    { return p.cliente }
func (p Pedido) Items() []ItemPedido { return p.items }
func (p Pedido) Fecha() time.Time    { return p.fecha }
func (p Pedido) Subtotal() float64   { return p.subtotal }
func (p Pedido) IVA() float64        { return p.iva }
func (p Pedido) Total() float64      { return p.total }
