package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/leogregianin/brcep/api"
)

// CepHandler ..
type CepHandler struct {
	PreferredAPI string
	CepApis      map[string]api.API
}

type responseError struct {
	Error string `json:"error"`
}

var prefferedApiError = &responseError{
	Error: "preferred api not available",
}

// Handle handles the request ..
func (h *CepHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cep := strings.Split(r.URL.Path, "/")[1]

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	preferredAPI, ok := h.CepApis[h.PreferredAPI]
	if !ok {
		j, _ := json.Marshal(prefferedApiError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(j)
		return
	}

	result, err := preferredAPI.Fetch(cep)
	if err != nil {
		j, _ := json.Marshal(&responseError{Error: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(j)
		return
	}

	result.Sanitize()

	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
