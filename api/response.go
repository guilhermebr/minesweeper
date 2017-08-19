package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Status  int         `json:"status,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func Success(result interface{}, status int) *Response {
	return &Response{
		Success: true,
		Status:  status,
		Result:  result,
	}
}
func (r *Response) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r)
}
