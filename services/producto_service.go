package services

import (
	"errors"
	"fmt"
	"strings"

	"ecommerce-go/models"
	"ecommerce-go/utils"
)

type ProductoService struct{}

func (s ProductoService) CrearProducto(id int, nombre string, precio float64, stock int) (models.Producto, error) {
	if err := utils.ValidarID(id); err != nil {
		return models.Producto{}, err
	}
	if err := utils.ValidarTextoNoVacio("nombre del producto", nombre); err != nil {
		return models.Producto{}, err
	}
	if err := utils.ValidarPrecio(precio); err != nil {
		return models.Producto{}, err
	}
	if err := utils.ValidarStock(stock); err != nil {
		return models.Producto{}, err
	}

	return models.NuevoProducto(id, strings.TrimSpace(nombre), precio, stock), nil
}

func (s ProductoService) RegistrarProducto(productos []models.Producto, producto models.Producto) []models.Producto {
	// Se crea una copia del slice antes de agregar el nuevo producto.
	// Esto mantiene un estilo más funcional, evitando modificar directamente la lista original.
	nuevaLista := append([]models.Producto{}, productos...)
	return append(nuevaLista, producto)
}

func (s ProductoService) ListarProductos(productos []models.Producto) {
	fmt.Println("\n===== CATÁLOGO DE PRODUCTOS =====")

	if len(productos) == 0 {
		fmt.Println("No existen productos registrados.")
		return
	}

	for _, producto := range productos {
		fmt.Printf("ID: %d | Producto: %s | Precio: $%.2f | Stock: %d\n",
			producto.ID(),
			producto.Nombre(),
			producto.Precio(),
			producto.Stock(),
		)
	}
}

func (s ProductoService) BuscarProductoPorID(productos []models.Producto, id int) (models.Producto, int, error) {
	for index, producto := range productos {
		if producto.ID() == id {
			return producto, index, nil
		}
	}

	return models.Producto{}, -1, errors.New("no se encontró un producto con el ID ingresado")
}

func (s ProductoService) ActualizarProducto(productos []models.Producto, index int, productoActualizado models.Producto) ([]models.Producto, error) {
	if index < 0 || index >= len(productos) {
		return productos, errors.New("no se pudo actualizar el producto porque el índice no es válido")
	}

	// Se actualiza sobre una copia del slice.
	// Así se protege la lista original y se devuelve una nueva versión actualizada.
	nuevaLista := append([]models.Producto{}, productos...)
	nuevaLista[index] = productoActualizado

	return nuevaLista, nil
}
