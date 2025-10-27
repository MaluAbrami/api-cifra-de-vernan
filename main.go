package main

import (
	"github.com/MaluAbrami/api-cifra-de-vernan/controllers"
	_ "github.com/MaluAbrami/api-cifra-de-vernan/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Cifra de Vernam
// @version 1.0
// @description API que implementa a cifra de Vernam (criptografia e descriptografia de mensagens).
// @host localhost:8080
// @BasePath /
func main() {
	//cria instancia do servidor
	r := gin.Default()

	//rotas de criptografia e descriptografia
	r.POST("/criptografar", controllers.Criptografar)
	r.POST("/descriptografar", controllers.Descriptografar)

	//rota para a documentacao da API Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//inicia o servidor na porta 8080
	r.Run(":8080")
}
