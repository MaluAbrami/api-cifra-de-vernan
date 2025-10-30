package models

type ResponseDecifrarDTO struct {
	TextoClaro string `json:"textoClaro" binding:"required"`
}
