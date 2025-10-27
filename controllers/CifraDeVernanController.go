package controllers

import (
	"net/http"

	"github.com/MaluAbrami/api-cifra-de-vernan/models"
	"github.com/MaluAbrami/api-cifra-de-vernan/services"
	"github.com/gin-gonic/gin"
)

// @Summary Criptografar mensagem usando Cifra de Vernam
// @Description Aplica a cifra de Vernam ao texto fornecido com a chave informada.
// @Tags criptografia
// @Accept  json
// @Produce  json
// @Param request body models.MensagemRequest true "Texto e chave"
// @Success 200 {object} models.MensagemResponse
// @Router /criptografar [post]
func Criptografar(c *gin.Context) {
	var req models.MensagemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro:": err.Error()})
		return
	}

	result, err := services.AplicarCifraDeVernan(req.Texto, req.Chave)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.MensagemResponse{Resultado: result})
}

// @Summary Descriptografar mensagem com cifra de vernan
// @Description Aplica a cifra de Vernam ao texto fornecido com a chave informada.
// @Tags descriptografia
// @Accept  json
// @Produce  json
// @Param request body models.MensagemRequest true "Texto e chave"
// @Success 200 {object} models.MensagemResponse
// @Router /descriptografar [post]
func Descriptografar(c *gin.Context) {
	var req models.MensagemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro:": err.Error()})
		return
	}

	result, err := services.AplicarCifraDeVernan(req.Texto, req.Chave)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.MensagemResponse{Resultado: result})
}
