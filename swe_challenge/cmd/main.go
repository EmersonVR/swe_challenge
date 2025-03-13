package main

import (
	"log"
	"net/http"

	"swe_challenge/internal/api"
	"swe_challenge/internal/repository"
	"swe_challenge/internal/services"
)

func main() {
	// Conectarse a la base de datos CockroachDB
	db, err := repository.NewDBConnection(
		"dusk-sage-5083.jxf.gcp-us-east1.cockroachlabs.cloud", // Host del cluster
		26257,                                                 // Puerto
		"pruebatec",                                           // Usuario
		"XF-yM8ZYvdWF-7GuzVRf9Q",                               // Password
		"stockdb",                                             // Base de datos
	)
	if err != nil {
		log.Fatalf("Error conectando a DB: %v\n", err)
	}
	defer db.Close()

	log.Println("¡Conexión a CockroachDB exitosa!")

	// Definir número de páginas a traer
	maxPages := 20 // Puedes modificar este valor según necesites

	// Llamar a la API externa e insertar datos con paginación
	items, err := services.FetchDataWithPagination(
		"https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6MSwiZW1haWwiOiJlbWVyc29udmFsbGVjaWxsYXJ1aXpAZ21haWwuY29tIiwiZXhwIjoxNzQxNjQ2MjAzLCJpZCI6IjAiLCJwYXNzd29yZCI6IicgT1IgJzEnPScxIn0.YHn1rVAPwaakOcr8XX0mvczCrKfP0r5M576HJls_MK8",
		maxPages,
	)
	if err != nil {
		log.Printf("Error al llamar la API externa: %v\n", err)
	} else {
		log.Printf("Se recibieron %d items en total de la API\n", len(items))

		accionRepo := repository.NewAccionRepository(db)
		newRecords := 0
		for _, item := range items {
			inserted, err := accionRepo.InsertAccion(item)
			if err != nil {
				log.Println("Error insertando item:", err)
			}
			if inserted {
				newRecords++
			}
		}

		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM acciones").Scan(&count)
		if err != nil {
			log.Printf("Error en SELECT COUNT(*): %v\n", err)
		} else {
			log.Printf("Total de registros en 'acciones': %d\n", count)
		}

		if newRecords == 0 {
			log.Println("Resumen: Datos ya existen.")
		} else {
			log.Println("Resumen: Datos insertados de manera exitosa.")
		}
	}

	// Usar Chi para definir el router y exponer los endpoints REST
	accionRepo := repository.NewAccionRepository(db)
	r := api.NewRouter(accionRepo)

	log.Println("Servidor corriendo en http://localhost:9090")
	http.ListenAndServe(":9090", r)
}
