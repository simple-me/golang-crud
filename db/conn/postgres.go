package conn
import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "log"
  "os"
)

func GetPostgres() *gorm.DB {
  dbURL := os.Getenv("PG_CONNSTRING")
  db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    return db
}
