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
	firebase := db.NewFirebase()
	db := db.NewDatabase(l, firebase)
	defer db.Client.Close()

}

func loadEnv(l *log.Logger) {
	if err := godotenv.Load(".env"); err != nil {
		l.Fatal(err)
	}
}
