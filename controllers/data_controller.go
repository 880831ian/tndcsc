package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tndcsc/services"
)

func GetData(c *gin.Context) {
	data, err := services.FetchData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}

	dataType := c.Query("type")
	switch dataType {
	case "swim":
		c.JSON(http.StatusOK, gin.H{
			"swim_pool_count":      data.SwimCount,
			"swim_pool_percentage": data.SwimPercentage,
		})
	case "fitness":
		c.JSON(http.StatusOK, gin.H{
			"fitness_center_count":      data.FitnessCount,
			"fitness_center_percentage": data.FitnessPercentage,
		})
	default:
		c.JSON(http.StatusOK, data)
	}
}
