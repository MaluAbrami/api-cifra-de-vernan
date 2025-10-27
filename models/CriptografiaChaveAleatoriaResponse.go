package models

type CriptografiaChaveAleatoriaResponse struct {
	TextoCriptografado string `json:"texto_criptografado"`
	ChaveUsada         string `json:"chave_usada"`
}
