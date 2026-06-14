package interfaces

import "ecommerce-go/models"

type ProductoManager interface {
	CrearProducto(id int, nombre string, precio float64, stock int) (models.Producto, error)
	RegistrarProducto(productos []models.Producto, producto models.Producto) []models.Producto
	ListarProductos(productos []models.Producto)
	BuscarProductoPorID(productos []models.Producto, id int) (models.Producto, int, error)
	ActualizarProducto(productos []models.Producto, index int, productoActualizado models.Producto) ([]models.Producto, error)
}

type ClienteManager interface {
	CrearCliente(id int, nombre string, correo string) (models.Cliente, error)
	RegistrarCliente(clientes []models.Cliente, cliente models.Cliente) []models.Cliente
	ListarClientes(clientes []models.Cliente)
	BuscarClientePorID(clientes []models.Cliente, id int) (models.Cliente, int, error)
}

type CalculadoraPedido interface {
	CalcularSubtotal(items []models.ItemPedido) float64
	CalcularIVA(subtotal float64) float64
	CalcularTotal(subtotal float64, iva float64) float64
}

type ReportePedido interface {
	MostrarPedido(pedido models.Pedido)
	ListarPedidos(pedidos []models.Pedido)
}

type PedidoManager interface {
	CalculadoraPedido
	ReportePedido
	CrearItem(producto models.Producto, cantidad int) (models.ItemPedido, models.Producto, error)
	CrearPedido(id int, cliente models.Cliente, items []models.ItemPedido) (models.Pedido, error)
}
