package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api-template/models"
	"net/http"
)

// @BasePath /api/v1

// FindSegments godoc
// @Summary find segments
// @Schemes
// @Description fetch all segments data
// @Tags segments
// @Accept json
// @Produce json
// @Success 200 {string} segment data
// @Router /segments [get]
func FindSegments(c *gin.Context) {
	var segment models.Segment

	var segments []models.Segment
	models.DB.Model(&segment).Find(&segments)

	c.JSON(http.StatusOK, gin.H{"data": segments})
}

// CreateSegment godoc
// @Summary create segment
// @Schemes
// @Description create segment entry
// @Tags segments
// @Accept json
// @Produce json
// @Param slug body string true "Segment slug"
// @Success 201 {string} segment data
// @Router /segments [post]
func CreateSegment(c *gin.Context) {
	var slug string

	//id, _ := strconv.Atoi(r.Body.Get("id"))

	segment := models.Segment{Slug: slug}

	models.DB.Create(&segment)

	c.JSON(http.StatusCreated, gin.H{"data": segment})
}

// FindSegment godoc
// @Summary find segment
// @Schemes
// @Description find segment entry by id
// @Tags segments
// @Accept json
// @Produce json
// @Param segment_id path int true "Segment ID"
// @Success 200 {string} segment data
// @Router /segments/{id} [get]
func FindSegment(c *gin.Context) {
	var segment_id uint

	if err := models.DB.Model(models.Segment{ID: segment_id}).First(&segment_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "segment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": segment_id})
}

// UpdateSegment godoc
// @Summary update segment
// @Schemes
// @Description update segment entry by id
// @Tags segments
// @Accept json
// @Produce json
// @Param segment_id path int true "Segment ID"
// @Param slug body string true "Segment slug"
// @Success 200 {string} segment data
// @Router /segments/{id} [put]
func UpdateSegment(c *gin.Context) {
	var segment_id uint

	var slug string

	if err := models.DB.Model(models.Segment{ID: segment_id}).First(&segment_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "segment not found"})
		return
	}

	if err := c.ShouldBindJSON(&slug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(models.Segment{ID: segment_id}).Updates(models.Segment{ID: segment_id, Slug: slug})

	c.JSON(http.StatusOK, gin.H{"data": segment_id})
}

// DeleteSegment godoc
// @Summary delete segment
// @Schemes
// @Description delete segment entry by id
// @Tags segments
// @Accept json
// @Produce json
// @Param segment_id path uint true "Segment ID"
// @Success 204 {string} empty content
// @Router /segments/{id} [delete]
func DeleteSegment(c *gin.Context) {
	var segment_id uint

	if err := models.DB.Model(models.Segment{ID: segment_id}).First(&segment_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "segment not found"})
		return
	}

	models.DB.Model(models.Segment{ID: segment_id}).Delete(&segment_id)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
