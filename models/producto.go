package models

import "errors"

// Producto representa un producto del e-commerce.
// Los campos son privados para aplicar encapsulación.
type Producto struct {
	id     int
	nombre string
	precio float64
	stock  int
}

func NuevoProducto(id int, nombre string, precio float64, stock int) Producto {
	return Producto{id: id, nombre: nombre, precio: precio, stock: stock}
}

func (p Producto) ID() int                      { return p.id }
func (p Producto) Nombre() string               { return p.nombre }
func (p Producto) Precio() float64              { return p.precio }
func (p Producto) Stock() int                   { return p.stock }
func (p Producto) Disponible(cantidad int) bool { return cantidad > 0 && cantidad <= p.stock }

// DescontarStock valida la cantidad solicitada y devuelve una nueva versión
// del producto con el stock actualizado. Si no existe stock suficiente,
// devuelve un error y no modifica el inventario real.
func (p Producto) DescontarStock(cantidad int) (Producto, error) {
	if cantidad <= 0 {
		return Producto{}, errors.New("la cantidad debe ser mayor a cero")
	}
	if cantidad > p.stock {
		return Producto{}, errors.New("stock insuficiente para realizar la compra")
	}
	p.stock -= cantidad
	return p, nil
}
