package main

import (
	"net/http"

	"github.com/MaluAbrami/api-cifra-de-vernan/controllers"
	_ "github.com/MaluAbrami/api-cifra-de-vernan/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Middleware simples para liberar CORS no Gin
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// @title API Cifra de Vernam
// @version 1.0
// @description API que implementa a cifra de Vernam (criptografia e descriptografia de mensagens).
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.POST("/cifrar", controllers.Cifrar)
	r.POST("/decifrar", controllers.Decifrar)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
