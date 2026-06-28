package data

import (
	"sync"

	"ecommerce-go/interfaces"
	"ecommerce-go/models"
)

// Store mantiene los datos temporalmente en memoria.
// Se usa Mutex para proteger los datos cuando llegan solicitudes web simultáneas.
type Store struct {
	Mutex               sync.Mutex
	Productos           []models.Producto
	Clientes            []models.Cliente
	Pedidos             []models.Pedido
	SiguienteIDProducto int
	SiguienteIDCliente  int
	SiguienteIDPedido   int
}

func NewStore(productoService interfaces.ProductoManager, clienteService interfaces.ClienteManager) *Store {
	productos := []models.Producto{}
	clientes := []models.Cliente{}
	p1, _ := productoService.CrearProducto(1, "Laptop HP", 650.00, 5)
	p2, _ := productoService.CrearProducto(2, "Mouse inalámbrico", 15.00, 20)
	p3, _ := productoService.CrearProducto(3, "Teclado mecánico", 45.00, 10)
	p4, _ := productoService.CrearProducto(4, "Audífonos Bluetooth", 30.00, 8)
	productos = productoService.RegistrarProducto(productos, p1)
	productos = productoService.RegistrarProducto(productos, p2)
	productos = productoService.RegistrarProducto(productos, p3)
	productos = productoService.RegistrarProducto(productos, p4)
	c1, _ := clienteService.CrearCliente(1, "Paola Puertas", "paola@example.com")
	clientes = clienteService.RegistrarCliente(clientes, c1)
	return &Store{Productos: productos, Clientes: clientes, Pedidos: []models.Pedido{}, SiguienteIDProducto: 5, SiguienteIDCliente: 2, SiguienteIDPedido: 1}
}
