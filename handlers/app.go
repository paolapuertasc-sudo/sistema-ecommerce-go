package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"ecommerce-go/data"
	"ecommerce-go/interfaces"
)

type App struct {
	Store           *data.Store
	ProductoService interfaces.ProductoManager
	ClienteService  interfaces.ClienteManager
	PedidoService   interfaces.PedidoManager
}

func NewApp(store *data.Store, productoService interfaces.ProductoManager, clienteService interfaces.ClienteManager, pedidoService interfaces.PedidoManager) *App {
	return &App{Store: store, ProductoService: productoService, ClienteService: clienteService, PedidoService: pedidoService}
}

func (app *App) RegistrarRutas(mux *http.ServeMux) {
	mux.HandleFunc("/", app.IndexHandler)
	mux.HandleFunc("/api", app.HomeAPIHandler)
	mux.HandleFunc("/api/productos", app.ProductosHandler)
	mux.HandleFunc("/api/productos/", app.ProductoPorIDHandler)
	mux.HandleFunc("/api/clientes", app.ClientesHandler)
	mux.HandleFunc("/api/clientes/", app.ClientePorIDHandler)
	mux.HandleFunc("/api/pedidos", app.PedidosHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

func (app *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		writeError(w, http.StatusNotFound, "ruta no encontrada")
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "no se pudo cargar la página", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (app *App) HomeAPIHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"mensaje": "API del Sistema de Gestión de E-commerce", "servicios": []string{"GET /api/productos", "POST /api/productos", "GET /api/productos/{id}", "GET /api/clientes", "POST /api/clientes", "GET /api/clientes/{id}", "GET /api/pedidos", "POST /api/pedidos"}})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
func decodeJSON(r *http.Request, destino any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(destino)
}
