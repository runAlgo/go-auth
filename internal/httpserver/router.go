package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/runAlgo/go-auth/internal/app"
	"github.com/runAlgo/go-auth/internal/middleware"
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
	r.POST("/login", userHandler.Login)

	// list all data/files (protected)
	api := r.Group("/api")
	api.Use(middleware.AuthRequired(a.Config.JWTSecret))

	api.GET("/files", func(c *gin.Context) {

		userID, _ := middleware.GetUserID(c)
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"userId": userID,
			"files":  []any{},
		})
	})

	api.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":    true,
			"files": []any{},
		})
	})

	admin := api.Group("/admin")
	admin.Use(middleware.RequireAdmin())
	
	admin.GET("/restricted", func(c *gin.Context) {
		role, _ := middleware.GetRole(c)
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"role": role,
		})
	})
	return r
}
