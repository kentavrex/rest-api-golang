package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	app "rest-api-golang"
	"strconv"
)

// @Summary Create segment
// @Tags segments
// @Description create segments
// @ID create-segment
// @Accept  json
// @Produce  json
// @Param input body app.Segment true "segment info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segments [post]
func (h *Handler) createSegment(c *gin.Context) {
	var input app.Segment

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Segment.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllSegmentsResponse struct {
	Data []app.Segment `json:"data"`
}

// @Summary Get all segments
// @Tags segments
// @Description get all segments
// @ID get-all-segments
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segments [get]
func (h *Handler) getAllSegments(c *gin.Context) {
	segments, err := h.services.Segment.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllSegmentsResponse{
		Data: segments,
	})
}

// @Summary Get segment by id
// @Tags segments
// @Description get segment by id
// @ID get-segment-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "segment id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segments/{id} [get]
func (h *Handler) getSegmentById(c *gin.Context) {
	segmentId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	segment, err := h.services.Segment.GetById(segmentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, segment)
}

// @Summary Update segment
// @Tags segments
// @Description update segment
// @ID update-segment
// @Accept  json
// @Produce  json
// @Param input body app.Segment true "segment info"
// @Param id path string true "segment id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segments/{id} [put]
func (h *Handler) updateSegment(c *gin.Context) {
	segmentId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input app.UpdateSegmentInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Segment.Update(segmentId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete segment
// @Tags segments
// @Description delete id
// @ID delete-segment
// @Accept  json
// @Produce  json
// @Param id path string true "segment id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segments/{id} [delete]
func (h *Handler) deleteSegment(c *gin.Context) {
	segmentId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Segment.Delete(segmentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
