package routes

// import (
// 	"jwt-golang/controllers"
// 	"jwt-golang/middlewares"

// 	"github.com/gin-gonic/gin"
// )

// type ExtendedUserController struct {
// 	*controllers.UserController
// }

// func (uc *ExtendedUserController) RegisterLoginRoutes(rg *gin.RouterGroup) {
// 	userroute := rg.Group("/user")
// 	userroute.POST("/login", uc.Login)
// 	userroute.POST("/register", uc.Register)
// }

// func (uc *ExtendedUserController) RegisterJwtCheckRoutes(rg *gin.RouterGroup) {
// 	userroute := rg.Group("/check")
// 	userroute.Use(middlewares.JwtAuthMiddleware())
// 	userroute.GET("/get/:id", uc.GetUser)
// }
