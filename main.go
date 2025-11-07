package main

import (
	"fmt"
	"log"
	"os"
	"shorclick/handlers"
	"shorclick/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("環境変数読み込まれんらしいよ...")
	}
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST") // または環境変数から取得
	dbPort := "5432"      // または環境変数から取得
	log.Println("Connecting to database with user:", dbUser, "and db name:", dbName)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	log.Println("DSN:", dsn)

	// DB接続の初期化
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}
	log.Println("Database connection established")

	// マイグレーションの実行
	err = db.AutoMigrate(&models.ShortLink{})
	if err != nil {
		log.Fatalln("Failed to migrate database:", err)
	}
	log.Println("Database migrated successfully")

	// Ginのセットアップ
	r:= gin.Default()
	log.Println("Starting server on :8080")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to shorclick",
		})
	})
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "don't setup api token",
		})
	})
	r.GET("/:id", handlers.ShortLink(db))
	r.Run(":8080")
}