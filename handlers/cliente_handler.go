package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

type crearClienteRequest struct {
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
}

func (app *App) ClientesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.ListarClientes(w, r)
	case http.MethodPost:
		app.CrearCliente(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
	}
}

func (app *App) ListarClientes(w http.ResponseWriter, r *http.Request) {
	app.Store.Mutex.Lock()
	defer app.Store.Mutex.Unlock()
	clientesDTO := []ClienteDTO{}
	for _, cliente := range app.Store.Clientes {
		clientesDTO = append(clientesDTO, clienteToDTO(cliente))
	}
	writeJSON(w, http.StatusOK, map[string]any{"clientes": clientesDTO})
}

func (app *App) CrearCliente(w http.ResponseWriter, r *http.Request) {
	var request crearClienteRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	app.Store.Mutex.Lock()
	defer app.Store.Mutex.Unlock()
	cliente, err := app.ClienteService.CrearCliente(app.Store.SiguienteIDCliente, request.Nombre, request.Correo)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	app.Store.Clientes = app.ClienteService.RegistrarCliente(app.Store.Clientes, cliente)
	app.Store.SiguienteIDCliente++
	writeJSON(w, http.StatusCreated, map[string]any{"mensaje": "cliente registrado correctamente", "cliente": clienteToDTO(cliente)})
}

func (app *App) ClientePorIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}
	idTexto := strings.TrimPrefix(r.URL.Path, "/api/clientes/")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		writeError(w, http.StatusBadRequest, "el ID del cliente debe ser un número válido")
		return
	}
	app.Store.Mutex.Lock()
	defer app.Store.Mutex.Unlock()
	cliente, _, err := app.ClienteService.BuscarClientePorID(app.Store.Clientes, id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"cliente": clienteToDTO(cliente)})
}
