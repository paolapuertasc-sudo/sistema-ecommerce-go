# Sistema de Gestión de E-commerce en Go

## Descripción

Sistema de Gestión de e-commerce desarrollado en Go. La versión final incluye una página web visual y una API REST con serialización JSON.

## Objetivo

Administrar productos, clientes y pedidos de un e-commerce básico, integrando conocimientos de programación funcional, estructuras, paquetes, encapsulación, interfaces, manejo de errores, servicios web y JSON.

## Tecnologías

- Go / Golang
- HTML
- CSS
- JavaScript
- Servicios Web REST
- JSON
- Git y GitHub

## Ejecución

```bash
go run .
```

Abrir en el navegador:

```txt
http://localhost:8080
```

## API JSON

| Método | Endpoint | Funcionalidad |
|---|---|---|
| GET | `/api/productos` | Lista productos |
| POST | `/api/productos` | Registra producto |
| GET | `/api/productos/{id}` | Consulta producto por ID |
| GET | `/api/clientes` | Lista clientes |
| POST | `/api/clientes` | Registra cliente |
| GET | `/api/clientes/{id}` | Consulta cliente por ID |
| GET | `/api/pedidos` | Lista pedidos |
| POST | `/api/pedidos` | Crea pedido |

## Funcionalidades visuales

- Panel principal con conteos de productos, clientes y pedidos.
- Formulario para registrar productos.
- Tabla de productos.
- Formulario para registrar clientes.
- Tabla de clientes.
- Formulario para crear pedidos.
- Historial visual de pedidos.

## Integrantes

- Paola Puertas
- [Integrante 2]
- [Integrante 3]
- [Integrante 4]
