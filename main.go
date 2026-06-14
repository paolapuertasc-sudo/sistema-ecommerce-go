package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ecommerce-go/interfaces"
	"ecommerce-go/models"
	"ecommerce-go/services"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Se declaran los servicios usando interfaces.
	// Esto permite que el programa dependa de comportamientos y no de implementaciones concretas.
	var gestorProductos interfaces.ProductoManager = services.ProductoService{}
	var gestorClientes interfaces.ClienteManager = services.ClienteService{}
	var gestorPedidos interfaces.PedidoManager = services.PedidoService{}

	productos := cargarProductosIniciales(gestorProductos)
	clientes := cargarClientesIniciales(gestorClientes)
	pedidos := []models.Pedido{}

	siguienteIDProducto := len(productos) + 1
	siguienteIDCliente := len(clientes) + 1
	siguienteIDPedido := 1

	for {
		mostrarMenuPrincipal()
		opcion := leerEntero(reader, "Seleccione una opción: ")

		switch opcion {
		case 1:
			registrarProducto(reader, &productos, gestorProductos, &siguienteIDProducto)
		case 2:
			gestorProductos.ListarProductos(productos)
		case 3:
			registrarCliente(reader, &clientes, gestorClientes, &siguienteIDCliente)
		case 4:
			gestorClientes.ListarClientes(clientes)
		case 5:
			crearPedido(reader, &productos, clientes, &pedidos, gestorProductos, gestorClientes, gestorPedidos, &siguienteIDPedido)
		case 6:
			gestorPedidos.ListarPedidos(pedidos)
		case 7:
			fmt.Println("Gracias por usar el Sistema de Gestión de E-commerce.")
			return
		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}
}

func cargarProductosIniciales(gestor interfaces.ProductoManager) []models.Producto {
	productos := []models.Producto{}

	p1, _ := gestor.CrearProducto(1, "Laptop HP", 650.00, 5)
	p2, _ := gestor.CrearProducto(2, "Mouse inalámbrico", 15.00, 20)
	p3, _ := gestor.CrearProducto(3, "Teclado mecánico", 45.00, 10)
	p4, _ := gestor.CrearProducto(4, "Audífonos Bluetooth", 30.00, 8)

	productos = gestor.RegistrarProducto(productos, p1)
	productos = gestor.RegistrarProducto(productos, p2)
	productos = gestor.RegistrarProducto(productos, p3)
	productos = gestor.RegistrarProducto(productos, p4)

	return productos
}

func cargarClientesIniciales(gestor interfaces.ClienteManager) []models.Cliente {
	clientes := []models.Cliente{}

	cliente, _ := gestor.CrearCliente(1, "Paola Puertas", "paola@example.com")
	clientes = gestor.RegistrarCliente(clientes, cliente)

	return clientes
}

func mostrarMenuPrincipal() {
	fmt.Println("\n====================================")
	fmt.Println("    SISTEMA DE GESTIÓN E-COMMERCE")
	fmt.Println("====================================")
	fmt.Println("1. Registrar producto")
	fmt.Println("2. Listar productos")
	fmt.Println("3. Registrar cliente")
	fmt.Println("4. Listar clientes")
	fmt.Println("5. Crear pedido")
	fmt.Println("6. Ver historial de pedidos")
	fmt.Println("7. Salir")
	fmt.Println("====================================")
}

func registrarProducto(reader *bufio.Reader, productos *[]models.Producto, gestor interfaces.ProductoManager, siguienteID *int) {
	fmt.Println("\n===== REGISTRO DE PRODUCTO =====")

	nombre := leerTexto(reader, "Ingrese el nombre del producto: ")
	precio := leerDecimal(reader, "Ingrese el precio del producto: ")
	stock := leerEntero(reader, "Ingrese el stock del producto: ")

	producto, err := gestor.CrearProducto(*siguienteID, nombre, precio, stock)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	*productos = gestor.RegistrarProducto(*productos, producto)
	*siguienteID = *siguienteID + 1

	fmt.Println("Producto registrado correctamente.")
}

func registrarCliente(reader *bufio.Reader, clientes *[]models.Cliente, gestor interfaces.ClienteManager, siguienteID *int) {
	fmt.Println("\n===== REGISTRO DE CLIENTE =====")

	nombre := leerTexto(reader, "Ingrese el nombre del cliente: ")
	correo := leerTexto(reader, "Ingrese el correo del cliente: ")

	cliente, err := gestor.CrearCliente(*siguienteID, nombre, correo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	*clientes = gestor.RegistrarCliente(*clientes, cliente)
	*siguienteID = *siguienteID + 1

	fmt.Println("Cliente registrado correctamente.")
}

func crearPedido(
	reader *bufio.Reader,
	productos *[]models.Producto,
	clientes []models.Cliente,
	pedidos *[]models.Pedido,
	gestorProductos interfaces.ProductoManager,
	gestorClientes interfaces.ClienteManager,
	gestorPedidos interfaces.PedidoManager,
	siguienteID *int,
) {
	if len(clientes) == 0 {
		fmt.Println("No existen clientes registrados. Primero registre un cliente.")
		return
	}

	if len(*productos) == 0 {
		fmt.Println("No existen productos registrados. Primero registre un producto.")
		return
	}

	gestorClientes.ListarClientes(clientes)
	idCliente := leerEntero(reader, "Ingrese el ID del cliente que realizará el pedido: ")

	cliente, _, err := gestorClientes.BuscarClientePorID(clientes, idCliente)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	items := []models.ItemPedido{}

	for {
		gestorProductos.ListarProductos(*productos)
		idProducto := leerEntero(reader, "Ingrese el ID del producto a comprar, o 0 para finalizar: ")

		if idProducto == 0 {
			break
		}

		producto, index, err := gestorProductos.BuscarProductoPorID(*productos, idProducto)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		cantidad := leerEntero(reader, "Ingrese la cantidad a comprar: ")

		item, productoActualizado, err := gestorPedidos.CrearItem(producto, cantidad)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Se actualiza el stock del producto después de agregarlo al pedido.
		// Si el stock no es suficiente, el sistema no permite agregar el item.
		listaActualizada, err := gestorProductos.ActualizarProducto(*productos, index, productoActualizado)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		*productos = listaActualizada
		items = append(items, item)

		fmt.Println("Producto agregado al pedido correctamente.")
	}

	pedido, err := gestorPedidos.CrearPedido(*siguienteID, cliente, items)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	*pedidos = append(*pedidos, pedido)
	*siguienteID = *siguienteID + 1

	gestorPedidos.MostrarPedido(pedido)
}

func leerTexto(reader *bufio.Reader, mensaje string) string {
	fmt.Print(mensaje)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

func leerEntero(reader *bufio.Reader, mensaje string) int {
	for {
		entrada := leerTexto(reader, mensaje)
		valor, err := strconv.Atoi(entrada)
		if err != nil {
			fmt.Println("Debe ingresar un número entero válido.")
			continue
		}
		return valor
	}
}

func leerDecimal(reader *bufio.Reader, mensaje string) float64 {
	for {
		entrada := leerTexto(reader, mensaje)
		entrada = strings.ReplaceAll(entrada, ",", ".")

		valor, err := strconv.ParseFloat(entrada, 64)
		if err != nil {
			fmt.Println("Debe ingresar un número decimal válido.")
			continue
		}
		return valor
	}
}
