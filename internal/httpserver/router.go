package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/runAlgo/go-auth/internal/app"
	"github.com/runAlgo/go-auth/internal/user"
)

func NewRouter(a *app.App) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", health)

	userRepo := user.NewRepo(a.DB)

	userSvc := user.NewService(userRepo, a.Config.JWTSecret)
	userHandler := user.NewHandler(userSvc)

	r.POST("/register", userHandler.Register)
	return r
}
