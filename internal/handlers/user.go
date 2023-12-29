package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"users/internal/data"
	"users/internal/data/memoryStore"
	"users/internal/dto"
	"users/internal/utils"
)

func AddUser(c *gin.Context) {
	var request dto.AddUserResponseBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	user, err := UsersQ(c).Add(data.User{
		Name:  request.Name,
		Email: request.Email,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	userResponse := dto.ToUserResponseBody(user)

	c.JSON(http.StatusOK, utils.ApiResponse("success", userResponse))
}

func GetUser(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	user, err := UsersQ(c).FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	userResponse := dto.ToUserResponseBody(user)

	c.JSON(http.StatusOK, utils.ApiResponse("success", userResponse))
}

func GetUsers(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	usersPerPage, err := strconv.Atoi(c.Query("usersPerPage"))
	if err != nil || usersPerPage < 1 {
		usersPerPage = 10
	}

	users, err := UsersQ(c).GetAll(page, usersPerPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	usersResponse := make([]dto.UserResponseBody, len(users))
	for i, user := range users {
		usersResponse[i] = dto.ToUserResponseBody(user)
	}

	c.JSON(http.StatusOK, utils.ApiResponse("success", usersResponse))
}

func UpdateUser(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	var request dto.UpdateUserResponseBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	user, err := UsersQ(c).Update(id, data.User{
		Name: request.Name,
	})
	if err != nil {
		if errors.Is(err, memoryStore.NotFoundError) {
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	userResponse := dto.ToUserResponseBody(user)

	c.JSON(http.StatusOK, utils.ApiResponse("success", userResponse))
}

func DeleteUser(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	err = UsersQ(c).Delete(id)
	if err != nil {
		if errors.Is(err, memoryStore.NotFoundError) {
			c.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse("success", utils.Null()))
}
