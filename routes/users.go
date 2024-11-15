package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pratyushvid3105/Go-Rest-API/models"
	"github.com/pratyushvid3105/Go-Rest-API/utils"
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

func login(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user request data", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not authenticate user", "error": err.Error()})
		return
	}
	
	var token string
	token, err = utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch user token", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successfull!", "token": token})
}

func getUsers(context *gin.Context){
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch users. Try again later.", "error": err})
		return
	}
	context.JSON(http.StatusOK, users)
}