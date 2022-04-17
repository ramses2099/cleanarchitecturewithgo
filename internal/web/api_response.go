package web

import (
	"encoding/json"
	"net/http"
)

type ReponseAPI struct {
	Success bool        `json:"success"`
	Status  int         `json:"status,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func Success(result interface{}, status int) *ReponseAPI {
	return &ReponseAPI{Success: true, Status: status, Result: result}
}

func (r *ReponseAPI) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r)
}