// Command migrate applies goose migrations embedded in the binary.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/LeGiaBaoDev/Go_Wander/backend/migrations"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	command := "up"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	pool, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("open database: %v", err)
	}
	defer func() {
		if closeErr := pool.Close(); closeErr != nil {
			log.Printf("close database: %v", closeErr)
		}
	}()

	goose.SetBaseFS(migrations.FS)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("set dialect: %v", err)
	}

	switch command {
	case "up":
		err = goose.Up(pool, ".")
	case "down":
		err = goose.Down(pool, ".")
	case "status":
		err = goose.Status(pool, ".")
	default:
		err = fmt.Errorf("unknown command %q (use up|down|status)", command)
	}
	if err != nil {
		log.Fatalf("migrate %s: %v", command, err)
	}
}
