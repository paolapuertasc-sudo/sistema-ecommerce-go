# Sistema de Gestión de E-commerce en Go

Proyecto académico desarrollado en Golang para gestionar productos, clientes y pedidos de una tienda virtual básica.

## Funcionalidades principales

- Menú interactivo por consola.
- Registro de productos.
- Listado de productos.
- Registro de clientes.
- Listado de clientes.
- Creación de pedidos.
- Validación de stock.
- Cálculo de subtotal, IVA y total.
- Historial de pedidos.

## Conceptos aplicados

- Encapsulación mediante atributos privados y métodos públicos.
- Manejo de errores con el tipo `error`.
- Interfaces para definir comportamientos de productos, clientes y pedidos.
- Comentarios en funcionalidades relevantes.
- Organización por paquetes.
- Enfoque funcional en funciones que reciben datos y devuelven resultados.

## Estructura del proyecto

```txt
sistema-ecommerce-go-v2/
├── go.mod
├── main.go
├── models/
│   ├── producto.go
│   ├── cliente.go
│   └── pedido.go
├── services/
│   ├── producto_service.go
│   ├── cliente_service.go
│   └── pedido_service.go
├── interfaces/
│   └── interfaces.go
└── utils/
    └── validaciones.go
```

## Ejecución

Desde la carpeta principal del proyecto ejecutar:

```bash
go run .
```

Opcionalmente se puede formatear el código con:

```bash
gofmt -w .
```

## Menú del sistema

```txt
====================================
    SISTEMA DE GESTIÓN E-COMMERCE
====================================
1. Registrar producto
2. Listar productos
3. Registrar cliente
4. Listar clientes
5. Crear pedido
6. Ver historial de pedidos
7. Salir
====================================
```
