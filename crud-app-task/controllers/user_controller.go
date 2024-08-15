package controllers

import (
	"crud-app-task/models"
	"crud-app-task/services"
	"crud-app-task/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Please provide valid format of user")
		return
	}
	if user.Username == "" || user.Password == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "user name password should not be empty")
		return
	}

	if err := ctrl.userService.CreateUser(&user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, "User created successfully", user)
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ctrl.userService.GetUserByID(uint(id))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "User retrieved successfully", user)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "User input is not correct")
		return
	}
	user.ID = uint(id)
	if err := ctrl.userService.UpdateUser(&user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "User updated successfully", user)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.userService.DeleteUser(uint(id)); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "User deleted successfully", nil)
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.userService.GetAllUsers()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Users retrieved successfully", users)
}
