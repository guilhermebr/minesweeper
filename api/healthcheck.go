package api

import "net/http"

// title: healthcheck
// path: /healthcheck
// method: GET
// responses:
//   200: OK
func (s *Services) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
