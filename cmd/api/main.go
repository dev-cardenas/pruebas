package main

import (
	"log"
	"pruebas/configs"
	"pruebas/internal/shared/infrastructure/db/postgres"
	userHttp "pruebas/internal/user/infrastructure/http"
)

func main() {
	// Cargar configuración
	cfg := configs.LoadConfig()

	// Conectar a la base de datos usando solo la configuración de la DB
	db, err := postgres.NewPostgresDB(cfg.Database) // Aquí se pasa cfg.Database
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Crear el router de usuario
	router := userHttp.NewRouter(db)

	// Iniciar servidor
	router.Run(":8080")
}
