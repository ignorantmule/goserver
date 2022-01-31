package goserver

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginswagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/marmotedu/goserver/internal/goserver/constant"
	"github.com/marmotedu/goserver/internal/goserver/controller/v1/post"
	"github.com/marmotedu/goserver/internal/goserver/controller/v1/user"
	"github.com/marmotedu/goserver/internal/goserver/store/mysql"
	"github.com/marmotedu/goserver/internal/pkg/middleware"
	"github.com/marmotedu/goserver/pkg/core"
	"github.com/marmotedu/goserver/pkg/errno"
	"github.com/marmotedu/goserver/pkg/token"
)

// loadRouter loads the middlewares, routes, handlers.
func loadRouter(g *gin.Engine, mw ...gin.HandlerFunc) {
	installMiddleware(g, mw...)
	installController(g)
}

// installMiddleware install Middlewares.
func installMiddleware(g *gin.Engine, mw ...gin.HandlerFunc) {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
}

// installController install controllers.
func installController(g *gin.Engine) {
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	// swagger api docs
	g.GET("/swagger/*any", ginswagger.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	storeIns, _ := mysql.GetMySQLFactoryOr()
	userController := user.NewUserController(storeIns)

	// api for authentication functionalities
	g.POST("/login", userController.Login)

	// The user handlers, requiring authentication
	v1 := g.Group("/v1")
	{
		// user RESTful resource
		userv1 := v1.Group("/users")
		{
			userv1.POST("", userController.Create)
			userv1.Use(authMiddleware())
			userv1.DELETE(":name", userController.Delete)
			userv1.PUT(":name", userController.Update)
			userv1.GET("", userController.List)
			userv1.GET(":name", userController.Get)
		}

		// secret RESTful resource
		postv1 := v1.Group("/posts", authMiddleware())
		{
			postController := post.NewPostController(storeIns)

			postv1.POST("", postController.Create)
			postv1.DELETE("", postController.DeleteCollection)
			postv1.DELETE(":postID", postController.Delete)
			postv1.PUT(":postID", postController.Update)
			postv1.GET("", postController.List)
			postv1.GET(":postID", postController.Get)
		}
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		username, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()

			return
		}

		c.Set(constant.XUsernameKey, username)
		c.Next()
	}
}
