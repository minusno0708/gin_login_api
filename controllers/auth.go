package controllers

import (
    "net/http"

    "jwt-gin/models"
    "jwt-gin/utils/token"

    "github.com/gin-gonic/gin"
)

type RegisterInput struct {
    UserID string `json:"user_id" binding:"required"`
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
    var input RegisterInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := models.User{UserID: input.UserID, Username: input.Username, Password: input.Password}

    user, err := user.Save()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": user,
    })
}

type LoginInput struct {
    UserID string `json:"user_id" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
    var input LoginInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := models.GenerateToken(input.UserID, input.Password)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": token,
    })
}

func CurrentUser(c *gin.Context) {
    userId, err := token.ExtractTokenId(c)

    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    var user models.User

    err = models.DB.First(&user, userId).Error

    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": user.PrepareOutput(),
    })
}