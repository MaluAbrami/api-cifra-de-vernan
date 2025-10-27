package models

type MensagemRequest struct {
	Texto string `json:"texto" binding:"required"`
	Chave string `json:"chave" binding:"required"`
}
