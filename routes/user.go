package routes

import (
	"example/rest_api/models"
	"example/rest_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user) //Convert JSON to object

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Fail to convert JSON to object"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't signup user data"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"success": globalResponseSuccess, "user": user})
}

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user) //Convert JSON to object

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Fail to convert JSON to object"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't authenticate user1"})
		return		
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't authenticate user2"})
		return		
	}

	context.JSON(http.StatusBadRequest, gin.H{"message": "Login successful", "token" : token})
}