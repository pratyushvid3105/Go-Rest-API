package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pratyushvid3105/Go-Rest-API/models"
)

func signup(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user request data", "error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user. Try again later.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully!"})
}