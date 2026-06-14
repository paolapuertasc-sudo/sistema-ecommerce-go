package models

import "strings"

type Cliente struct {
	id     int
	nombre string
	correo string
}

func NuevoCliente(id int, nombre string, correo string) Cliente {
	return Cliente{
		id:     id,
		nombre: nombre,
		correo: correo,
	}
}

func (c Cliente) ID() int {
	return c.id
}

func (c Cliente) Nombre() string {
	return c.nombre
}

func (c Cliente) Correo() string {
	return c.correo
}

func (c Cliente) EsValido() bool {
	return strings.TrimSpace(c.nombre) != "" && strings.Contains(c.correo, "@")
}
