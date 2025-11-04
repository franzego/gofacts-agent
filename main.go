package main

import (
	"fmt"
	"log"

	"github.com/franzego/stage03/internal"
	"github.com/franzego/stage03/logic"
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

	r.Run(":8080")
	fmt.Print(logic.FactRandom())
}
