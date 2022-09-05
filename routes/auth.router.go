package routes

import (
	"github.com/LakeCombs/gethired/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())

}
