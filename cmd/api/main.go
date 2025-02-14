package main

import (
	"log"
	"net/http"

	"plc_project/internal/api"
)

func main() {
	// Utiliza o router configurado em api.SetupRoutes()
	router := api.SetupRoutes()

	log.Println("API Server iniciando na porta 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
