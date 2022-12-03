package initializers

import (
	"github.com/lilendian0x00/go-mvc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDatabase(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error when connecting to DB: %s\n", err.Error())
	}
	log.Println("Connected to DB Successfully!")
}

func SyncDB()  {
	err := DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalln(err)
	}
}