package controllers

import (
	"os"
	"testing"

	"github.com/felipedavid/vrcursos/src/infrastructure/database"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbUrl := os.Getenv("TEST_DATABASE_URL")

	db, err := database.ConnectToDatabase(dbUrl)
	if err != nil {
		panic(err)
	}

	err = database.RunUpMigrations(db, "file://migrations")
	if err != nil && err.Error() != "no change" {
		panic(err)
	}

	m.Run()

	err = database.RunDownMigrations(db, "file://migrations")
	if err != nil && err.Error() != "no change" {
		panic(err)
	}
}

func TestStudentsList(t *testing.T) {

}
