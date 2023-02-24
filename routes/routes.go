package routes

import (
	"jwt-golang/controllers"
	"jwt-golang/middlewares"

	"github.com/gin-gonic/gin"
)

type RouterService struct {
	Usercontroller controllers.UserController
}

func NewRouterService(usercontroller controllers.UserController) RouterService {
	return RouterService{
		Usercontroller: usercontroller,
	}
}

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

func (rs *RouterService) RegisterLoginRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/login", rs.Usercontroller.Login)
	userroute.POST("/register", rs.Usercontroller.Register)
}

func (rs *RouterService) RegisterJwtCheckRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/check")
	userroute.Use(middlewares.JwtAuthMiddleware())
	userroute.GET("/get/:id", rs.Usercontroller.GetUser)
}
