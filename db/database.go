package Database

import (
	"fmt"
	Model "kawan-usaha-api/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() *gorm.DB {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: fmt.Sprintf(
					"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
					os.Getenv("DB_HOST"),
					os.Getenv("DB_USER"),
					os.Getenv("DB_PASS"),
					os.Getenv("DB_NAME"),
					os.Getenv("DB_PORT"),
					os.Getenv("DB_SSLMODE"),
					os.Getenv("DB_TIMEZONE"),
				),
				PreferSimpleProtocol: true, // disables implicit prepared statement usage
			}),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Model
	if err = db.AutoMigrate(&Model.User{}); err != nil {
		log.Fatal(err.Error())
	}

	return db
}
