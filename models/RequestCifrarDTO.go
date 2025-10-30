package models

type RequestCifrarDTO struct {
	TextoClaro string `json:"textoClaro" binding:"required"`
	Chave      string `json:"chave" binding:"required"`
}
