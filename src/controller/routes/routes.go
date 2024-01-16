package routes

import (
	"github.com/WesleySCorrea/crud-completo-golang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userID", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userID", userController.UpdateUser)
	r.DELETE("/deleteUser/:userID", userController.DeleteUser)

}
