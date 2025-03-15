package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"lucares.github.com/minicloud/minicloud/adapters/repositories"
	"lucares.github.com/minicloud/minicloud/adapters/router"
	"lucares.github.com/minicloud/minicloud/domain/ports"
)

var ctx context.Context

func main() {
	ctx = context.Background()

	router := router.CreateRouter()

	router.Run()
}

func init() {
	ctx = context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DSN_MINICLOUD_DB"))
	if err != nil {
		log.Fatal("cant start db connection ")
	}

	dependencies := map[string]interface{}{
		ports.USER_REPOSITORY_KEY_CTX: repositories.NewRepository(db),
	}

	injectDepedencies(dependencies)
}

func injectDepedencies(dependencies map[string]interface{}) {
	for k, v := range dependencies {
		ctx = context.WithValue(ctx, k, v)
	}
}
