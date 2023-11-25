package db

import (
	"fmt"
	"log"
	"os"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {

	host := os.Getenv("DBHOST")
	dbUserName := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  sslmode=disable", host, dbUserName, pass, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	DB = db
	DB.AutoMigrate(&domain.User{})
	DB.AutoMigrate(&domain.Brand{})
	DB.AutoMigrate(&domain.Category{})
	DB.AutoMigrate(&domain.Product{})
	DB.AutoMigrate(&domain.Cart{})

}
