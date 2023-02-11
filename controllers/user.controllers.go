package controllers

import (
	"jwt-golang/middlewares"
	"jwt-golang/models"
	"jwt-golang/services"
	"jwt-golang/token"
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

func (uc *UserController) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := uc.UserService.LoginUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Login verified", "token": token})
}

func (uc *UserController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := uc.UserService.RegisterUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Registration success", "token": token})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user_id, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id == "" {
		id = user_id
	}

	u, err := uc.UserService.GetUserByID(&user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func (uc *UserController) RegisterLoginRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/login", uc.Login)
	userroute.POST("/register", uc.Register)
}

func (uc *UserController) RegisterJwtCheckRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/check")
	userroute.Use(middlewares.JwtAuthMiddleware())
	userroute.GET("/get/:id", uc.GetUser)
}