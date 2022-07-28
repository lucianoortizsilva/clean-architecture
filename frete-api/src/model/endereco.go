package model

type Endereco struct {
	Cep        string `json:"cep,omitempty"`
	Logradouro string `json:"logradouro,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Cidade     string `json:"cidade,omitempty"`
	Estado     string `json:"estado,omitempty"`
}
