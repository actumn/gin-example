package app

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func GinEngine(handler *Handler) *gin.Engine {
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "test Zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: handler.authenticator,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*UserClaims); ok {
				return jwt.MapClaims{
					"user_id": user.UserId,
				}
			}

			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			return jwt.ExtractClaims(c)
		},
		Authorizator: func(user interface{}, c *gin.Context) bool {
			if _, ok := user.(*jwt.MapClaims); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("jwt.New Error")
	}

	// sign up
	ginEngine.POST("/signup", handler.signUp)

	// login
	ginEngine.POST("/login", authMiddleware.LoginHandler)

	// auth
	authGroup := ginEngine.Group("/auth")
	authGroup.Use(authMiddleware.MiddlewareFunc())
	{
		authGroup.GET("", handler.authProfile)
		authGroup.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	// api
	apiGroup := ginEngine.Group("/api")
	apiGroup.Use(authMiddleware.MiddlewareFunc())
	{
		apiGroup.GET("/books/:id", handler.book)
		apiGroup.POST("/books", handler.postBook)
		apiGroup.PUT("/books/:id", handler.putBook)
		apiGroup.DELETE("/books/:id", handler.deleteBook)
	}

	return ginEngine
}

type UserClaims struct {
	UserId uint
}

func ConvertClaims(claims jwt.MapClaims) UserClaims {
	return UserClaims{
		UserId: uint(claims["user_id"].(float64)),
	}
}
