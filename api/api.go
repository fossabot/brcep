package api

import (
	"regexp"
)

// BrCepResult holds the standardized JSON result from the API
type BrCepResult struct {
	Cep         string `json:"cep"`
	Endereco    string `json:"endereco"`
	Bairro      string `json:"bairro"`
	Complemento string `json:"complemento"`
	Cidade      string `json:"cidade"`
	Uf          string `json:"uf"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	DDD         string `json:"ddd"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
}

// API holds the interface that represents an API capable of fetching
// a CEP an return a standardized struct
type API interface {
	// Fetch should fetch the result from the
	// API and return as BrCepResult
	Fetch(cep string) (*BrCepResult, error)
}

var cepSanitizer = regexp.MustCompile("[^0-9]+")

// Sanitize replaces string with "" ..
func (r *BrCepResult) Sanitize() {
	r.Cep = cepSanitizer.ReplaceAllString(r.Cep, "")
}
