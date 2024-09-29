// auth/routes.go

package api

import (
	"github.com/gin-gonic/gin"

	"backend/api/controllers"
	"backend/api/middleware"

	"firebase.google.com/go/auth"
	"gorm.io/gorm"
)

// SetupRouter configura las rutas para el módulo de autenticación
func SetupRouter(r *gin.Engine, db *gorm.DB, authClient *auth.Client) {
	authRoutes := r.Group("/")
	{
		authRoutes.POST("/login", controllers.LoginUser)
		authRoutes.POST("/register", func(c *gin.Context) {
			controllers.RegisterUser(c, db)
		})
		authRoutes.POST("/verify-code", func(c *gin.Context) {
			controllers.VerifyCode(c, db, authClient)
		})
		authRoutes.POST("/resend-code", middleware.AuthMiddleware(authClient), func(c *gin.Context) {
			controllers.ResendCode(c, db)
		})
		authRoutes.PATCH("/update", middleware.AuthMiddleware(authClient), func(c *gin.Context) {
			controllers.UpdateProfile(c, db, authClient)
		})
		authRoutes.POST("/upload-photo", middleware.AuthMiddleware(authClient), func(c *gin.Context) {
			controllers.UploadPhoto(c, db, authClient)
		})
		authRoutes.POST("/forgot-password", func(c *gin.Context) {
			controllers.ForgotPassword(c, authClient, db)
		})
		authRoutes.POST("/change-password", middleware.JWTMiddleware(), func(c *gin.Context) {
			controllers.ChangePassword(c, authClient)
		})
		authRoutes.GET("/validate-token", middleware.AuthMiddleware(authClient), func(c *gin.Context) {
			controllers.ValidateToken(c, authClient)
		})
	}
}
