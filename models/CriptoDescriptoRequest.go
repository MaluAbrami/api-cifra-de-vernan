package models

type CriptoDescriptoRequest struct {
	Texto string `json:"texto" binding:"required"`
	Chave string `json:"chave" binding:"required"`
}
