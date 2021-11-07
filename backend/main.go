package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/terrytay/twitter/backend/infrastructures/db"
)

func main() {

	// Setup logger
	l := log.New(os.Stdout, "twitter", log.LstdFlags)

	// Load enivronment variables into app
	loadEnv(l)

	// Initialize database
	firebase, err := db.NewFirebase(l)
	if err != nil {
		l.Fatal(err)
	}
	db := db.NewDatabase(firebase)
	l.Println("[MAIN]", "Database connection established")
	defer db.Client.Close()

}

func loadEnv(l *log.Logger) {
	if err := godotenv.Load(".env"); err != nil {
		l.Fatal(err)
	}
}
