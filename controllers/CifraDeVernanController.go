package controllers

import (
	"net/http"

	"github.com/MaluAbrami/api-cifra-de-vernan/models"
	"github.com/MaluAbrami/api-cifra-de-vernan/services"
	"github.com/gin-gonic/gin"
)

// @Summary Criptografar com chave aleatória
// @Description Aplica a cifra de Vernam ao texto fornecido com chave aleatória.
// @Tags criptografia
// @Accept  json
// @Produce  json
// @Param request body models.CriptografiaChaveAleatoriaRequest true "Texto e chave"
// @Success 200 {object} models.CriptografiaChaveAleatoriaResponse
// @Router /criptografar-com-chave-aleatoria [post]
func CriptografarComChaveAleatoria(c *gin.Context) {
	var req models.CriptografiaChaveAleatoriaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro:": err.Error()})
		return
	}

	result, chave, err := services.Criptografar(req.Texto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.CriptografiaChaveAleatoriaResponse{TextoCriptografado: result, ChaveUsada: chave})
}

// @Summary Criptografar com sua chave
// @Description Aplica a cifra de Vernam ao texto fornecido com a chave informada.
// @Tags criptografia
// @Accept  json
// @Produce  json
// @Param request body models.CriptoDescriptoRequest true "Texto e chave"
// @Success 200 {object} models.CriptoDescriptoResponse
// @Router /criptografar-com-chave-propria [post]
func CriptografarComChavePropria(c *gin.Context) {
	var req models.CriptoDescriptoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro:": err.Error()})
		return
	}

	result, err := services.AplicarCifraDeVernan(req.Texto, req.Chave)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.CriptoDescriptoResponse{Resultado: result})
}

// @Summary Descriptografar mensagem com cifra de vernan
// @Description Aplica a cifra de Vernam ao texto fornecido com a chave informada.
// @Tags descriptografia
// @Accept  json
// @Produce  json
// @Param request body models.CriptoDescriptoRequest true "Texto e chave"
// @Success 200 {object} models.CriptoDescriptoResponse
// @Router /descriptografar [post]
func Descriptografar(c *gin.Context) {
	var req models.CriptoDescriptoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro:": err.Error()})
		return
	}

	result, err := services.AplicarCifraDeVernan(req.Texto, req.Chave)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.CriptoDescriptoResponse{Resultado: result})
}
