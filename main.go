package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"lucares.github.com/minicloud/minicloud/adapters/repositories"
	"lucares.github.com/minicloud/minicloud/adapters/router"
	"lucares.github.com/minicloud/minicloud/domain/ports"
)

var ctx context.Context

func main() {
	router := router.CreateRouter()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	log.Info().Msg("starting server at :8000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Msgf("erro when run app: %e", err)
	}
}

func init() {
	// Create ctx
	ctx = context.Background()

	// Configure logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("error loading .env file")
	}

	// Create pool db
	db, err := sql.Open("postgres", os.Getenv("DSN_MINICLOUD_DB"))
	if err != nil {
		log.Fatal().Msg("cant start db connection ")
	}
	seedDb(ctx, db)

	// Inject dependencies into ctx
	dependencies := map[string]interface{}{
		ports.USER_REPOSITORY_KEY_CTX: repositories.NewRepository(db),
	}

	ctx = injectDependencies(ctx, dependencies)
	log.Info().Msg("Dependencies injected")

}

func seedDb(ctx context.Context, db *sql.DB) {
	_, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS users(
		id VARCHAR(255) NOT NULL,
		name VARCHAR(255) UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		PRIMARY KEY(id)
	)`)
	if err != nil {
		log.Fatal().Msgf("erro when seed db: %e", err)
	}
	log.Info().Msg("Seed users success")

	log.Info().Msg("Database seed success")
}

func injectDependencies(ctx context.Context, dependencies map[string]interface{}) context.Context {
	for k, v := range dependencies {
		ctx = context.WithValue(ctx, k, v)
	}

	return ctx
}
