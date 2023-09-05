package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	app "rest-api-golang"
	"strconv"
)

// @Summary Create user
// @Tags users
// @Description create user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body app.CreateUserInput true "user info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users [post]
func (h *Handler) createUser(c *gin.Context) {
	var input app.CreateUserInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.User.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllUsersResponse struct {
	Data []app.UserGet `json:"data"`
}

// @Summary Get all users
// @Tags users
// @Description get all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.User.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: users,
	})
}

// @Summary Get user by id
// @Tags users
// @Description get user by id
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "user id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id} [get]
func (h *Handler) getUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := h.services.User.GetById(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Update user
// @Tags users
// @Description update user
// @ID update-user
// @Accept  json
// @Produce  json
// @Param input body app.UpdateUserInput true "user info"
// @Param id path string true "user id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id} [put]
func (h *Handler) updateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input app.UpdateUserInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.User.Update(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete user
// @Tags users
// @Description delete id
// @ID delete-user
// @Accept  json
// @Produce  json
// @Param id path string true "user id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.User.Delete(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Add segments to user
// @Tags users
// @Description add segments to user
// @ID add-segments-to-user
// @Accept  json
// @Produce  json
// @Param input body app.AddUserSegmentInput true "segments ids"
// @Param id path string true "user id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id}/segments [post]
func (h *Handler) addUserSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input app.AddUserSegmentInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.User.AddSegments(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Get user segments
// @Tags users
// @Description get user segments
// @ID get-user-segments
// @Accept  json
// @Produce  json
// @Param id path string true "user id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id}/segments [get]
func (h *Handler) getUserSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	userSegments, err := h.services.User.GetSegments(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userSegments)
}

// @Summary Delete all user segments
// @Tags users
// @Description delete all user segments
// @ID delete-user-segments
// @Accept  json
// @Produce  json
// @Param id path string true "user id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id}/segments [delete]
func (h *Handler) deleteUserSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.User.DeleteSegments(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
