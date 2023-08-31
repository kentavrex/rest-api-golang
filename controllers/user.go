package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-rest-api-template/models"
	"net/http"
	"strconv"
)

// @BasePath /api/v1

// FindUsers godoc
// @Summary find users
// @Schemes
// @Description fetch all users data
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} user data
// @Router /users [get]
func FindUsers(c *gin.Context) {
	var user models.User

	var users []models.User
	models.DB.Model(&user).Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// CreateUser godoc
// @Summary create user
// @Schemes
// @Description create user entry
// @Tags users
// @Accept json
// @Produce json
// @Success 201 {string} user data
// @Router /users [post]
func CreateUser(c *gin.Context) {
	user := models.User{}

	models.DB.Model(&user).Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// FindUser godoc
// @Summary find user
// @Schemes
// @Description find user entry by id
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {string} user data
// @Router /users/{id} [get]
func FindUser(c *gin.Context) {
	var user_id uint

	if err := models.DB.Model(models.User{ID: user_id}).First(&user_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user_id})
}

// DeleteUser godoc
// @Summary delete user
// @Schemes
// @Description delete user entry by id
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 204 {string} empty content
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	user_id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	fmt.Println(user_id)
	if err := models.DB.Model(models.User{ID: uint(user_id)}).First(&user_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	models.DB.Model(models.User{ID: uint(user_id)}).Delete(&user_id)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
