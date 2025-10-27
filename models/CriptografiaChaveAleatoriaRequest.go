package models

type CriptografiaChaveAleatoriaRequest struct {
	Texto string `json:"texto" binding:"required"`
}
