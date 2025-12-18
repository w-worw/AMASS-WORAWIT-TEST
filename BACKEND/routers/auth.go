package routers

import (
	"amass-test/handlers"
	"amass-test/repositories"
	"amass-test/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAuthRouter(
	r *gin.Engine,
	db *gorm.DB,
) {
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	auth := handlers.NewAuthHandler(authService)
	authGroup := r.Group("/auth")
	authGroup.POST("/login", auth.Login)
	authGroup.POST("/register", auth.Register)
}
