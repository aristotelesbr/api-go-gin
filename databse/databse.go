package databse

import (
	"log"

	"github.com/aristotelesbr/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabse() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("Failed to open database: ", err)
	}

	DB.AutoMigrate(&models.Student{})
}
