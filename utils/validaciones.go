package utils

import (
	"errors"
	"strings"
)

func ValidarID(id int) error {
	if id <= 0 {
		return errors.New("el ID debe ser mayor a cero")
	}
	return nil
}

func ValidarTextoNoVacio(campo string, valor string) error {
	if strings.TrimSpace(valor) == "" {
		return errors.New("el campo " + campo + " no puede estar vacío")
	}
	return nil
}

func ValidarPrecio(precio float64) error {
	if precio <= 0 {
		return errors.New("el precio debe ser mayor a cero")
	}
	return nil
}

func ValidarStock(stock int) error {
	if stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	return nil
}

func ValidarCantidad(cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}
	return nil
}

func ValidarCorreo(correo string) error {
	if err := ValidarTextoNoVacio("correo", correo); err != nil {
		return err
	}

	if !strings.Contains(correo, "@") || !strings.Contains(correo, ".") {
		return errors.New("el correo ingresado no tiene un formato válido")
	}

	return nil
}
