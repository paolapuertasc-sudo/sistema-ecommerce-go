package services

import "sistema-ecommerce-go/models"

func CrearCliente(id int, nombre string, correo string) models.Cliente {
	return models.Cliente{ID: id, Nombre: nombre, Correo: correo}
}
