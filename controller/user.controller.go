package controller

import (
	models "gethired/model"
	"gethired/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}

}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	userID := ctx.Param(":id")
	user, err := uc.UserService.GetUser(&userID)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func (uc *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")

}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")

}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroutes := rg.Group("/user")
	userroutes.POST("/create", uc.CreateUser)
	userroutes.GET("/:id", uc.GetUser)
	userroutes.GET("/", uc.GetAll)
	userroutes.PUT("/update", uc.UpdateUser)
	userroutes.DELETE("/delete/:id", uc.DeleteUser)

}
