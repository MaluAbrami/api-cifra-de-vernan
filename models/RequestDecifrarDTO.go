package models

type RequestDecifrarDTO struct {
	TextoCifrado string `json:"textoCifrado"`
	Chave        string `json:"chave"`
}
