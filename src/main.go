package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/felipedavid/vrcursos/src/application/routes"
	"github.com/felipedavid/vrcursos/src/core/model"
	"github.com/felipedavid/vrcursos/src/infrastructure/config"
	"github.com/felipedavid/vrcursos/src/infrastructure/repository/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "err", err)
		return
	}

	addr := os.Getenv("ADDR")
	databaseUrl := os.Getenv("DATABASE_URL")

	db, err := config.ConnectToDatabase(databaseUrl)
	if err != nil {
		slog.Error("Unable to sablish database connection", "err", err)
		return
	}

	userRepo := postgres.NewPostgresStudentRepository(db)

	err = userRepo.Save(context.Background(), &model.Student{
		Name: "Felipe David",
	})
	if err != nil {
		slog.Error("Unable to save user", "err", err)
		return
	}

	mux := routes.DefineRoutes()

	slog.Info("Starting web server", "addr", addr)
	err = http.ListenAndServe(addr, mux)
	slog.Error("Unable to start server", "err", err)
	os.Exit(-1)
}
