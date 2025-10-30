package controllers

import (
	"net/http"

	"github.com/MaluAbrami/api-cifra-de-vernan/models"
	"github.com/MaluAbrami/api-cifra-de-vernan/services"
	"github.com/gin-gonic/gin"
)

// @Summary Cifrar uma mensagem com uma chave
// @Description Aplica a cifra de Vernam ao texto fornecido com a chave informada.
// @Tags criptografia
// @Accept  json
// @Produce  json
// @Param request body models.RequestCifrarDTO true "Texto e chave"
// @Success 200 {object} models.ResponseCifrarDTO
// @Router /cifrar [post]
func Cifrar(c *gin.Context) {
	var req models.RequestCifrarDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro:": err.Error()})
		return
	}

	result, err := services.AplicarCifraDeVernan(req.TextoClaro, req.Chave)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.ResponseCifrarDTO{TextoCifrado: result})
}

// @Summary Decifrar mensagem com cifra de vernan
// @Description Aplica a cifra de Vernam ao texto fornecido com a chave informada.
// @Tags descriptografia
// @Accept  json
// @Produce  json
// @Param request body models.RequestDecifrarDTO true "Texto e chave"
// @Success 200 {object} models.ResponseDecifrarDTO
// @Router /decifrar [post]
func Decifrar(c *gin.Context) {
	var req models.RequestDecifrarDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro:": err.Error()})
		return
	}

	result, err := services.AplicarCifraDeVernan(req.TextoCifrado, req.Chave)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.ResponseDecifrarDTO{TextoClaro: result})
}
