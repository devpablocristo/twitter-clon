package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/devpablocristo/monorepo/projects/qh/wire"
)

func main() {
	// Create a context with cancellation to handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Capture system signals for clean termination
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		<-sigChan // Wait for a signal
		log.Println("Received termination signal. Shutting down the application...")
		cancel()
	}()

	// Initialize dependencies using Wire
	deps, err := wire.Initialize()
	if err != nil {
		log.Fatalf("Error initializing dependencies: %s", err)
	}

	// Cargar datos de prueba en el repositorio de personas.
	if err := seedTestData(ctx, deps.PersonUseCases, deps.UserUseCases, deps.TweetUseCases); err != nil {
		log.Printf("Error seeding test data: %v", err)
	}

	if err := RunGormMigrations(ctx, deps.GormRepository); err != nil {
		log.Fatalf("Failed to run Gorm's migrations: %v", err)
	}

	if err := RunCassandraMigrations(ctx, deps.CassandraRepository); err != nil {
		log.Fatalf("Failed to run Cassandra's migrations: %v", err)
	}

	if err := RunWebServer(ctx, deps); err != nil {
		log.Fatalf("Error starting the application: %s", err)
	}

	log.Println("Application terminated successfully.")
}
