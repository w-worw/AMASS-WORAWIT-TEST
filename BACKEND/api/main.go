package main

import (
	"amass-test/databases"
	"amass-test/middlewares"
	"amass-test/routers"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(
		middlewares.CORS(),
	)

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	db := databases.NewMysql(dbUser, dbPass, dbHost, dbPort, dbName)

	routers.NewRouter(r, db)

	if err := r.Run(":" + port); err != nil {
		fmt.Println("Failed to run server:", err)
	}
}
