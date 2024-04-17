package handler

import (
	"go-web-api/config"
	"go-web-api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler adalah fungsi yang diekspor yang kompatibel dengan Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Koneksi ke database
	configs.ConnectDB()

	// Inisialisasi Gorilla Mux Router
	router := mux.NewRouter()

	// Setup subrouter untuk /api
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Definisikan rute
	routes.AuthRoutes(apiRouter)
	routes.UserRoutes(apiRouter)

	// Gunakan router sebagai handler
	router.ServeHTTP(w, r)
}