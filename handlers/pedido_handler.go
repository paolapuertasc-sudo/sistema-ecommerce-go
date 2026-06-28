package handlers

import (
	"ecommerce-go/models"
	"net/http"
)

type crearPedidoRequest struct {
	ClienteID int                `json:"cliente_id"`
	Items     []crearItemRequest `json:"items"`
}
type crearItemRequest struct {
	ProductoID int `json:"producto_id"`
	Cantidad   int `json:"cantidad"`
}

func (app *App) PedidosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.ListarPedidos(w, r)
	case http.MethodPost:
		app.CrearPedido(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
	}
}

func (app *App) ListarPedidos(w http.ResponseWriter, r *http.Request) {
	app.Store.Mutex.Lock()
	defer app.Store.Mutex.Unlock()
	pedidosDTO := []PedidoDTO{}
	for _, pedido := range app.Store.Pedidos {
		pedidosDTO = append(pedidosDTO, pedidoToDTO(pedido))
	}
	writeJSON(w, http.StatusOK, map[string]any{"pedidos": pedidosDTO})
}

func (app *App) CrearPedido(w http.ResponseWriter, r *http.Request) {
	var request crearPedidoRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	app.Store.Mutex.Lock()
	defer app.Store.Mutex.Unlock()
	cliente, _, err := app.ClienteService.BuscarClientePorID(app.Store.Clientes, request.ClienteID)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	if len(request.Items) == 0 {
		writeError(w, http.StatusBadRequest, "el pedido debe tener al menos un producto")
		return
	}
	items := []models.ItemPedido{}
	// Se usa una copia temporal para que el inventario real solo cambie si todo el pedido es válido.
	productosTemporales := append([]models.Producto{}, app.Store.Productos...)
	for _, itemRequest := range request.Items {
		producto, index, err := app.ProductoService.BuscarProductoPorID(productosTemporales, itemRequest.ProductoID)
		if err != nil {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		item, productoActualizado, err := app.PedidoService.CrearItem(producto, itemRequest.Cantidad)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		productosTemporales, err = app.ProductoService.ActualizarProducto(productosTemporales, index, productoActualizado)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		items = append(items, item)
	}
	pedido, err := app.PedidoService.CrearPedido(app.Store.SiguienteIDPedido, cliente, items)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	app.Store.Productos = productosTemporales
	app.Store.Pedidos = append(app.Store.Pedidos, pedido)
	app.Store.SiguienteIDPedido++
	writeJSON(w, http.StatusCreated, map[string]any{"mensaje": "pedido creado correctamente", "pedido": pedidoToDTO(pedido)})
}
