package internal

import (
	"net/http"

	"github.com/franzego/stage03/logic"
	"github.com/franzego/stage03/models"
	"github.com/gin-gonic/gin"
)

func GetRandomTips(c *gin.Context) {
	f := logic.TipsRandom()
	c.JSON(http.StatusOK, models.TipsResponse{
		Tips: f,
	})
}
func GetRandomFacts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": logic.FactRandom()})
}
func GetRandomSyntax(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": logic.SyntaxRandom()})
}
func GetRandomHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": logic.HistoryRandom()})
}

func DetectCategoryHandler(c *gin.Context) {
	var req struct {
		Message string `json:"message"`
	}
	// if req.Message == "" {
	// 	c.JSON(http.StatusBadRequest, models.ErrorResponse{
	// 		Message: "You did not add any message",
	// 	})
	// 	return
	// }
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Invalid request type",
			Error:   err.Error(),
		})
		return
	}
	response := logic.DetectCategory(req.Message)
	c.JSON(http.StatusOK, gin.H{"response": response})
}

func PostTelexWebHook(c *gin.Context) {
	var req models.TelexRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Invalid Request",
			Error:   err.Error(),
		})
		return
	}
	response := logic.DetectCategory(req.Data.Content)
	c.JSON(http.StatusOK, gin.H{"response": response})
}
