package database

import (
	"log"
	"os"

	"github.com/pllehrman/trivia-example/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := "host=db user=username password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected to db!")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations!")
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance {
		Db:db,
	}

}
