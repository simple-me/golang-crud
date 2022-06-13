package conn

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgres() *gorm.DB {
	dbURL := os.Getenv("PG_CONNSTRING")
	if os.Getenv("PG_CONNSTRING") == "" {
		// Establish a default conn string for testing
		dbURL = "postgres://root:secret@127.0.0.1:5432/products"
	}
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
