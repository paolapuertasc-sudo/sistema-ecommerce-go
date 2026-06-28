package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

type crearProductoRequest struct {
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock  int     `json:"stock"`
}

func (app *App) ProductosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.ListarProductos(w, r)
	case http.MethodPost:
		app.CrearProducto(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
	}
}

func (app *App) ListarProductos(w http.ResponseWriter, r *http.Request) {
	app.Store.Mutex.Lock()
	defer app.Store.Mutex.Unlock()
	productosDTO := []ProductoDTO{}
	for _, producto := range app.Store.Productos {
		productosDTO = append(productosDTO, productoToDTO(producto))
	}
	writeJSON(w, http.StatusOK, map[string]any{"productos": productosDTO})
}

func (app *App) CrearProducto(w http.ResponseWriter, r *http.Request) {
	var request crearProductoRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	app.Store.Mutex.Lock()
	defer app.Store.Mutex.Unlock()
	producto, err := app.ProductoService.CrearProducto(app.Store.SiguienteIDProducto, request.Nombre, request.Precio, request.Stock)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	app.Store.Productos = app.ProductoService.RegistrarProducto(app.Store.Productos, producto)
	app.Store.SiguienteIDProducto++
	writeJSON(w, http.StatusCreated, map[string]any{"mensaje": "producto registrado correctamente", "producto": productoToDTO(producto)})
}

func (app *App) ProductoPorIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}
	idTexto := strings.TrimPrefix(r.URL.Path, "/api/productos/")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		writeError(w, http.StatusBadRequest, "el ID del producto debe ser un número válido")
		return
	}
	app.Store.Mutex.Lock()
	defer app.Store.Mutex.Unlock()
	producto, _, err := app.ProductoService.BuscarProductoPorID(app.Store.Productos, id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"producto": productoToDTO(producto)})
}
