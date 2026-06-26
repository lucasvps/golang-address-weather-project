package clients

import "example.com/address-weather-project/internal/domain"

type ViaCepResponse struct {
	Cep         string
	Logradouro  string
	Complemento string
	Unidade     string
	Bairro      string
	Localidade  string
	Uf          string
	Estado      string
	Regiao      string
	Ibge        string
	Gia         string
	Ddd         string
	Siafi       string
}

func (v ViaCepResponse) ToAddress() domain.Address {
	return domain.Address{
		PostalCode: v.Cep,
		Street:     v.Logradouro,
		City:       v.Localidade,
		State:      v.Uf,
	}
}
