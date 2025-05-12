package main

import (
	"example/gorm/database"
	"example/gorm/database/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Init()
	router := gin.Default()

	router.POST("/add-country",
		func(ctx *gin.Context) {
			var newCountry models.Countries

			if err := ctx.BindJSON(&newCountry); err != nil {
				ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := db.Create(&newCountry).Error; err != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			}

			ctx.IndentedJSON(http.StatusOK, gin.H{"success": true, "message": "new country added successfully", "newCountry": newCountry})
		})

	router.GET("/get-all-countries",
		func(ctx *gin.Context) {
			allCountries := []models.Countries{}

			if err := db.Find(&allCountries).Error; err != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
				return
			}

			ctx.IndentedJSON(http.StatusOK, gin.H{"success": true, "allCountries": allCountries})
		})

	router.Run(":8080")
}
