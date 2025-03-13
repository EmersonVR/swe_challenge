package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"swe_challenge/internal/repository"
)

// NewRouter configura y retorna un router con los endpoints REST definidos.
func NewRouter(accionRepo *repository.AccionRepository) http.Handler {
	r := chi.NewRouter()

	// Middleware de CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Permitir peticiones desde el frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Cache de pre-flight
	}))

	// Endpoint para obtener todas las acciones
	r.Get("/acciones", func(w http.ResponseWriter, r *http.Request) {
		limit := 1000  // Número de registros por página
		offset := 0  // Default (página 1)
	
		queryLimit := r.URL.Query().Get("limit")
		queryOffset := r.URL.Query().Get("offset")
	
		if queryLimit != "" {
			l, err := strconv.Atoi(queryLimit)
			if err == nil {
				limit = l
			}
		}
		if queryOffset != "" {
			o, err := strconv.Atoi(queryOffset)
			if err == nil {
				offset = o
			}
		}
	
		acciones, err := accionRepo.GetAll(limit, offset)
		if err != nil {
			http.Error(w, "Error al obtener acciones", http.StatusInternalServerError)
			return
		}
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(acciones)
	})
	

	// Endpoint para obtener las mejores recomendaciones de acciones
	r.Get("/recomendaciones", func(w http.ResponseWriter, r *http.Request) {
		recomendaciones, err := accionRepo.GetRecommendedActions()
		if err != nil {
			http.Error(w, "Error al obtener recomendaciones", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(recomendaciones)
	})

	// Endpoint base opcional
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("¡Bienvenido a la API de Acciones!"))
	})

	return r
}
