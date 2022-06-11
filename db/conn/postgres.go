package conn

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgres() *gorm.DB {
	dbURL := os.Getenv("PG_CONNSTRING")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func GetPostgresTest() *gorm.DB {
	dbURL := "postgres://root:secret@127.0.0.1:5432/products"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
