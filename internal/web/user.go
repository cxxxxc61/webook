package web

import (
	"github.com/gin-gonic/gin"
)

// 业务处理
func Registerusersroutes(sever *gin.Engine, u *UserHandler) {
	s := sever.Group("/users")
	s.POST("/signup", u.Signup)
	s.POST("/post", u.Post)
	//s.POST("/login", u.Login)
	s.POST("/login", u.LoginJWT)
	//s.POST("/edit", u.Profile)
	s.GET("/profile", u.Profile)
}
