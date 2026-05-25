package models

import "time"

type Producto struct {
	ID     int
	Nombre string
	Precio float64
	Stock  int
}

type Cliente struct {
	ID     int
	Nombre string
	Correo string
}

type Pedido struct {
	ID        int
	Cliente   Cliente
	Productos []Producto
	Fecha     time.Time
	Subtotal  float64
	IVA       float64
	Total     float64
}
