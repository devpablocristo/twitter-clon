package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	cass "github.com/devpablocristo/monorepo/pkg/databases/nosql/cassandra/gocql"
	gorm "github.com/devpablocristo/monorepo/pkg/databases/sql/gorm"

	personmodels "github.com/devpablocristo/monorepo/projects/qh/internal/person/repository/models"
	usermodels "github.com/devpablocristo/monorepo/projects/qh/internal/user/repository/models"

	wire "github.com/devpablocristo/monorepo/projects/qh/wire"
)

func RunWebServer(ctx context.Context, deps *wire.Dependencies) error {
	if deps == nil {
		return errors.New("dependencies cannot be nil")
	}

	log.Println("Initializing routes...")

	// Manejo de errores en middlewares
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if ok {
				log.Printf("Error while initializing middlewares: %v\n", err)
			} else {
				log.Printf("Unknown panic occurred: %v\n", r)
			}
		}
	}()

	// Configurar middlewares globales primero
	if len(deps.Middlewares.Global) > 0 {
		deps.GinServer.GetRouter().Use(deps.Middlewares.Global...)
	}

	// Registrar rutas
	registerRoutes(deps)

	log.Println("Starting Gin server...")
	return deps.GinServer.RunServer(ctx)
}

func registerRoutes(deps *wire.Dependencies) {
	deps.PersonHandler.Routes()
	deps.UserHandler.Routes()
	deps.AutheHandler.Routes()
	deps.NotificationHandler.Routes()
	deps.TweetHandler.Routes()
}

func RunGormMigrations(ctx context.Context, repo gorm.Repository) error {
	log.Println("Starting database migrations...")

	// Verificar la conexión antes de proceder
	sqlDB, err := repo.Client().DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}
	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}

	// Lista de modelos a migrar
	modelsToMigrate := []any{
		&personmodels.Person{},
		&usermodels.User{},
		&usermodels.Follow{},
	}

	// Medir el tiempo de ejecución
	start := time.Now()
	if err := repo.AutoMigrate(modelsToMigrate...); err != nil {
		return fmt.Errorf("failed to migrate database models: %w", err)
	}
	duration := time.Since(start)

	log.Printf("Database migrations completed successfully in %s.", duration)
	return nil
}

// RunCassandraMigrations ejecuta las migraciones para Cassandra, creando el keyspace y las tablas necesarias.
func RunCassandraMigrations(ctx context.Context, repo cass.Repository) error {
	// Obtén el session de Cassandra.
	session := repo.GetSession()

	// 1. Crear el keyspace (si no existe).
	createKeyspaceCQL := `
		CREATE KEYSPACE IF NOT EXISTS mi_keyspace 
		WITH REPLICATION = { 'class': 'SimpleStrategy', 'replication_factor': 1 }`
	if err := session.Query(createKeyspaceCQL).WithContext(ctx).Exec(); err != nil {
		return fmt.Errorf("failed to create keyspace: %w", err)
	}
	log.Println("Keyspace 'mi_keyspace' created or already exists.")

	// 2. Crear la tabla "tweets".
	createTweetsTableCQL := `
		CREATE TABLE IF NOT EXISTS tweets (
			id uuid PRIMARY KEY,
			user_id text,
			content text,
			created_at timestamp
		)`
	if err := session.Query(createTweetsTableCQL).WithContext(ctx).Exec(); err != nil {
		return fmt.Errorf("failed to create table 'tweets': %w", err)
	}
	log.Println("Table 'tweets' created or already exists.")

	// 3. Crear la tabla desnormalizada "timeline_by_user".
	createTimelineTableCQL := `
		CREATE TABLE IF NOT EXISTS timeline_by_user (
			user_id text,
			created_at timestamp,
			tweet_id text,
			content text,
			PRIMARY KEY (user_id, created_at, tweet_id)
		) WITH CLUSTERING ORDER BY (created_at DESC)
	`
	if err := session.Query(createTimelineTableCQL).WithContext(ctx).Exec(); err != nil {
		return fmt.Errorf("failed to create table 'timeline_by_user': %w", err)
	}
	log.Println("Table 'timeline_by_user' created or already exists.")

	log.Println("Cassandra migrations completed successfully.")
	return nil
}
