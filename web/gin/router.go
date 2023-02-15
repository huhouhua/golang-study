package gin

import (
	"github.com/gin-gonic/gin"
	"golang-study/web/gin/authorization"
	"golang-study/web/gin/controller/authorize"
	"golang-study/web/gin/controller/user"
)

func register() {
	g := gin.Default()

	apiv1Group := g.Group("/api/v1")
	{
		authRouter := apiv1Group.Group("/authorize")
		{
			authCtrl := authorize.NewAuthzController()

			authRouter.POST("login", authCtrl.Login)
			authRouter.GET("logout", authCtrl.LoginOut)
		}

		userRouter := apiv1Group.Group("/user")
		{
			userRouter.Use(authorization.TokenAuth())
			userCtrl := user.NewUserController()

			userRouter.GET("/:id", userCtrl.UserInfo)
		}
	}

	g.Run(":8086")

}
