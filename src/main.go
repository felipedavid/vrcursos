package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/felipedavid/vrcursos/src/application/controllers"
	"github.com/felipedavid/vrcursos/src/application/routes"
	"github.com/felipedavid/vrcursos/src/infrastructure/config"
	"github.com/felipedavid/vrcursos/src/infrastructure/repository/postgres"
	"github.com/joho/godotenv"
)

const (
	migrationsPath = "file://migrations"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "err", err)
		os.Exit(-1)
	}

	addr := os.Getenv("ADDR")
	databaseUrl := os.Getenv("DATABASE_URL")

	db := setupDatabase(databaseUrl)

	userRepo := postgres.NewPostgresStudentRepository(db)
	courseRepo := postgres.NewPostgresCourseRepository(db)

	userControllers := controllers.NewStudentController(userRepo)
	courseControllers := controllers.NewCourseController(courseRepo)

	routes := routes.DefineRoutes(userControllers, courseControllers)

	slog.Info("Starting web server", "addr", addr)
	err = http.ListenAndServe(addr, routes)
	slog.Error("Unable to start web server", "err", err)
	os.Exit(-1)
}

// setupDatabase stablishes a connection to the database and runs the migrations
func setupDatabase(databaseUrl string) *sql.DB {
	db, err := config.ConnectToDatabase(databaseUrl)
	if err != nil {
		slog.Error("Unable to stablish database connection", "err", err)
		os.Exit(-1)
	}

	err = config.RunUpMigrations(db, migrationsPath)
	if err != nil && err.Error() != "no change" {
		slog.Error("Error while running migrations", "err", err)
		os.Exit(-1)
	}

	if err != nil && err.Error() != "no change" {
		slog.Info("No migrations to apply, all migrations are up to date")
	} else {
		slog.Info("Migrations applied successfully!")
	}

	return db
}
