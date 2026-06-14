package services

import (
	"errors"
	"fmt"
	"strings"

	"ecommerce-go/models"
	"ecommerce-go/utils"
)

type ClienteService struct{}

func (s ClienteService) CrearCliente(id int, nombre string, correo string) (models.Cliente, error) {
	if err := utils.ValidarID(id); err != nil {
		return models.Cliente{}, err
	}
	if err := utils.ValidarTextoNoVacio("nombre del cliente", nombre); err != nil {
		return models.Cliente{}, err
	}
	if err := utils.ValidarCorreo(correo); err != nil {
		return models.Cliente{}, err
	}

	return models.NuevoCliente(id, strings.TrimSpace(nombre), strings.TrimSpace(correo)), nil
}

func (s ClienteService) RegistrarCliente(clientes []models.Cliente, cliente models.Cliente) []models.Cliente {
	// Se devuelve una nueva lista con el cliente agregado.
	// Esto evita modificar directamente la lista recibida como parámetro.
	nuevaLista := append([]models.Cliente{}, clientes...)
	return append(nuevaLista, cliente)
}

func (s ClienteService) ListarClientes(clientes []models.Cliente) {
	fmt.Println("\n===== CLIENTES REGISTRADOS =====")

	if len(clientes) == 0 {
		fmt.Println("No existen clientes registrados.")
		return
	}

	for _, cliente := range clientes {
		fmt.Printf("ID: %d | Cliente: %s | Correo: %s\n",
			cliente.ID(),
			cliente.Nombre(),
			cliente.Correo(),
		)
	}
}

func (s ClienteService) BuscarClientePorID(clientes []models.Cliente, id int) (models.Cliente, int, error) {
	for index, cliente := range clientes {
		if cliente.ID() == id {
			return cliente, index, nil
		}
	}

	return models.Cliente{}, -1, errors.New("no se encontró un cliente con el ID ingresado")
}
