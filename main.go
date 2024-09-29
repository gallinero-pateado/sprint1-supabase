package main

import (
	"log"
	"net/http"
	"time"

	"backend/api"
	utils "backend/api/utils"
	_ "backend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Backend API
// @version 1.0
// @description API para el backend de una aplicación web con autenticación de Firebase.
// @contact.name moisesnks
// @contact.url https://github.com/moisesnks
// @contact.email moisesnks@utem.cl
// @BasePath /
// @host localhost:8081
func main() {
	// Configurar y obtener los clientes de Firebase
	firestoreClient, authClient, storageClient, err := configSetup()
	if err != nil {
		log.Fatalf("Error configurando Firebase y clientes: %v", err)
	}

	// Cerrar cliente de Firestore al finalizar la aplicación
	defer firestoreClient.Close()

	// Conectar a la base de datos
	db, err := utils.OpenGormDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Configurar router Gin
	r := gin.New()

	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Middleware JSON Logger
	r.Use(gin.Logger())

	// Middleware Recovery
	r.Use(gin.Recovery())

	// Configurar Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Redireccionar la raíz a la documentación de Swagger
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Configurar rutas desde el paquete de autenticación (api)
	api.SetupRouter(r, firestoreClient, authClient, storageClient, db)

	// Iniciar el servidor solo después de la inicialización completa
	port := "8081"
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
