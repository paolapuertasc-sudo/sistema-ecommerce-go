package interfaces

import "ecommerce-go/models"

// ProductoManager define el contrato para gestionar productos.
type ProductoManager interface {
	CrearProducto(id int, nombre string, precio float64, stock int) (models.Producto, error)
	RegistrarProducto(productos []models.Producto, producto models.Producto) []models.Producto
	BuscarProductoPorID(productos []models.Producto, id int) (models.Producto, int, error)
	ActualizarProducto(productos []models.Producto, index int, productoActualizado models.Producto) ([]models.Producto, error)
}

// ClienteManager define el contrato para gestionar clientes.
type ClienteManager interface {
	CrearCliente(id int, nombre string, correo string) (models.Cliente, error)
	RegistrarCliente(clientes []models.Cliente, cliente models.Cliente) []models.Cliente
	BuscarClientePorID(clientes []models.Cliente, id int) (models.Cliente, int, error)
}

// CalculadoraPedido agrupa los cálculos principales de una venta.
type CalculadoraPedido interface {
	CalcularSubtotal(items []models.ItemPedido) float64
	CalcularIVA(subtotal float64) float64
	CalcularTotal(subtotal float64, iva float64) float64
}

// PedidoManager define el contrato para la creación de pedidos.
type PedidoManager interface {
	CalculadoraPedido
	CrearItem(producto models.Producto, cantidad int) (models.ItemPedido, models.Producto, error)
	CrearPedido(id int, cliente models.Cliente, items []models.ItemPedido) (models.Pedido, error)
}
