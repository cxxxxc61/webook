package main

import (
	"github.com/cxxxxc61/webook/internal/web"
	"github.com/cxxxxc61/webook/internal/web/middleware"
	"github.com/cxxxxc61/webook/repository"
	"github.com/cxxxxc61/webook/repository/dao"
	"github.com/cxxxxc61/webook/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db := initdb()
	u := initUser(db)
	sever := initwebsever()
	web.Registerusersroutes(sever, u)
	sever.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务

	//text()

	//sever := gin.Default()
	//u := &web.UserHandler{}
	//sever.POST("/users/login", func(c *gin.Context) {
	//	type Signupreq struct {
	//		Email           string `json:"email"`
	//		ConfirmPassword string `json:"confirm_password"`
	//		Password        string `json:"password"`
	//	}
	//
	//	var req Signupreq
	//	if err := c.Bind(&req); err != nil {
	//		return
	//	}
	//	c.String(http.StatusOK, "注册成功")
	//})
	//sever.POST("/post", u.Post)
	//sever.POST("/users/signup", u.Signup)
	//sever.POST("/login", u.Login)
	//sever.POST("/edit", u.Profile)
	//sever.GET("/profile", u.Profile)
	//sever.Run(":8080")
}
func initdb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:BazKT%HlbsP3@tcp(117.50.198.118:30336)/cxc_webook"))
	if err != nil {
		panic(err)
	}
	err = dao.Inittable(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initUser(db *gorm.DB) *web.UserHandler {
	d := dao.NewUserDao(db)
	repo := repository.NewUserRepository(d)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initwebsever() *gin.Engine {
	sever := gin.Default()
	sever.Use(cors.New(cors.Config{
		//详细地址
		//AllowOrigins:     []string{"https://foo.com"},
		AllowMethods: []string{"POST", "GET"},
		//业务请求中可以带上的头
		AllowHeaders: []string{"Origin", "Content-Type"},
		//允许传入的头
		ExposeHeaders: []string{"x-jwt-token"},
		//用户认证信息
		AllowCredentials: true,
		//类型地址
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "your.com")
		},
		MaxAge: 12 * time.Minute,
	}))

	store, err := redis.NewStore(16,
		"tcp", "127.0.0.1:6379",
		"", "cxc20060601",
		[]byte("bHO2mkqCDKSB2GsqikJGlQURD0KtwiuZI4zpWZYolG7QCE64hTM0r6O5VhrdjFHt"))
	if err != nil {
		panic(err)
	}
	sever.Use(sessions.Sessions("mysession", store))
	sever.Use(middleware.NewLoginjwtMiddlewareBuild().
		Ignorepath("/users/login").
		Ignorepath("/users/signup").Build())

	return sever
}
