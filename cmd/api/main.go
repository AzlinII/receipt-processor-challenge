package main

import (
	"log"
	"net/http"

	"github.com/AzlinII/receipt-processor-challenge/internal/handlers"
	"github.com/AzlinII/receipt-processor-challenge/internal/repo"
	"github.com/AzlinII/receipt-processor-challenge/internal/services"
)

func main() {
	// Create a new router
	router := http.NewServeMux()

	// init services
	pointsDB := repo.NewPointsDB()
	pointsService := services.NewPointsService(pointsDB)

	// init handler
	handler := handlers.NewHandler(pointsService)
	handler.Init(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("running server...")
	srv.ListenAndServe()
}
