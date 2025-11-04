package main

import (
	"log"
	"os"

	"github.com/franzego/stage03/internal"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// logic.TipsRandom()
	if err := godotenv.Load(); err != nil {
		log.Print(err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "GoFactsAgent is alive 🐹"})
	})
	r.GET("/fact", internal.GetRandomFacts)
	r.GET("/tips", internal.GetRandomTips)
	r.GET("/history", internal.GetRandomHistory)
	r.GET("/syntax", internal.GetRandomSyntax)

	r.POST("/detect", internal.DetectCategoryHandler)
	r.POST("telex/webhook", internal.PostTelexWebHook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	// fmt.Print(logic.FactRandom())
}
