package services

import (
	"errors"

	"ecommerce-go/models"
	"ecommerce-go/utils"
)

const ivaPorcentaje = 0.15

type PedidoService struct{}

func (s PedidoService) CrearItem(producto models.Producto, cantidad int) (models.ItemPedido, models.Producto, error) {
	if err := utils.ValidarCantidad(cantidad); err != nil {
		return models.ItemPedido{}, models.Producto{}, err
	}
	// Se descuenta el stock antes de confirmar el item en el pedido.
	productoActualizado, err := producto.DescontarStock(cantidad)
	if err != nil {
		return models.ItemPedido{}, models.Producto{}, err
	}
	return models.NuevoItemPedido(producto, cantidad), productoActualizado, nil
}

func (s PedidoService) CrearPedido(id int, cliente models.Cliente, items []models.ItemPedido) (models.Pedido, error) {
	if err := utils.ValidarID(id); err != nil {
		return models.Pedido{}, err
	}
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
	// Suma el subtotal de cada item: precio del producto por cantidad.
	subtotal := 0.0
	for _, item := range items {
		subtotal += item.Subtotal()
	}
	return subtotal
}
func (s PedidoService) CalcularIVA(subtotal float64) float64                { return subtotal * ivaPorcentaje }
func (s PedidoService) CalcularTotal(subtotal float64, iva float64) float64 { return subtotal + iva }
