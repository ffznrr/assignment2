package main

import (
	"assignment2/handler"
	"assignment2/structur"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = 5432
	dbName   = "postgres"
)

func main() {
	database := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&structur.Items{}, &structur.Orders{})

	router := gin.Default()

	router.POST("/orders", handler.CreateOrders)
	router.GET("/orders/:id", handler.ReadOrders)
	router.PUT("/orders/:id", handler.UpdateOrder)
	router.DELETE("/orders/:id", handler.DeleteOrder)

	router.Run(":8180")
}
