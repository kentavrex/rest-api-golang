package controllers

import (
	"golang-rest-api-template/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// FindUserSegments godoc
// @Summary find segments
// @Schemes
// @Description fetch all user_segment data
// @Tags user_segments
// @Accept json
// @Produce json
// @Success 200 {string} segment data
// @Router /user_segments [get]
func FindUserSegments(c *gin.Context) {
	var user_segment models.UserSegment

	var user_segments []models.UserSegment
	models.DB.Model(&user_segment).Find(&user_segments)

	c.JSON(http.StatusOK, gin.H{"data": user_segments})
}

// CreateUserSegment godoc
// @Summary create segment
// @Schemes
// @Description create segment entry
// @Tags user_segments
// @Accept json
// @Produce json
// @Param user_id body int true "User id"
// @Param segment_id body int true "Segment id"
// @Success 201 {string} segment data
// @Router /user_segments [post]
func CreateUserSegment(c *gin.Context) {
	var user_id int
	var segment_id int

	user_segment := models.UserSegment{UserId: user_id, SegmentId: segment_id}

	models.DB.Model(&user_segment).Create(&user_segment)

	c.JSON(http.StatusCreated, gin.H{"data": user_segment})
}

// FindUserSegment godoc
// @Summary find user_segment
// @Schemes
// @Description find user_segment entry by id
// @Tags user_segments
// @Accept json
// @Produce json
// @Param user_segment_id path int true "UserSegment ID"
// @Success 200 {string} user_segment data
// @Router /user_segments/{id} [get]
func FindUserSegment(c *gin.Context) {

	var user_segment_id uint

	if err := models.DB.Model(models.UserSegment{ID: user_segment_id}).First(&user_segment_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_segment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user_segment_id})
}

// UpdateUserSegment godoc
// @Summary update user user_segment
// @Schemes
// @Description update user_segment entry by id
// @Tags user_segments
// @Accept json
// @Produce json
// @Param user_segment_id body int true "User id"
// @Param user_id body int true "User id"
// @Param segment_id body int true "Segment id"
// @Success 200 {string} user_segment data
// @Router /user_segments/{id} [put]
func UpdateUserSegment(c *gin.Context) {
	var user_segment_id uint
	var user_id int
	var segment_id int

	if err := models.DB.Model(models.UserSegment{ID: user_segment_id}).First(&segment_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_segment not found"})
		return
	}

	if err := models.DB.Model(models.User{ID: uint(user_id)}).First(&segment_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := models.DB.Model(models.Segment{ID: uint(segment_id)}).First(&segment_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "segment not found"})
		return
	}

	models.DB.Model(models.UserSegment{ID: user_segment_id, UserId: user_id, SegmentId: segment_id}).Updates(models.UserSegment{UserId: user_id, SegmentId: segment_id})

	c.JSON(http.StatusOK, gin.H{"data": segment_id})
}

// DeleteUserSegment godoc
// @Summary delete segment
// @Schemes
// @Description delete segment entry by id
// @Tags user_segments
// @Accept json
// @Produce json
// @Param user_segment_id path uint true "UserSegment ID"
// @Success 204 {string} empty content
// @Router /user_segments/{id} [delete]
func DeleteUserSegment(c *gin.Context) {
	var user_segment_id uint

	if err := models.DB.Model(models.UserSegment{ID: user_segment_id}).First(&user_segment_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "segment not found"})
		return
	}

	models.DB.Model(models.UserSegment{ID: user_segment_id}).Delete(&user_segment_id)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
