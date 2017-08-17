package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func Start(log *logrus.Logger) error {
	// API Routes
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", healthcheck).Methods("GET")

	// Middleware
	n := negroni.Classic()
	n.UseHandler(r)

	//Run Server
	log.Infoln("Server running on port :3000")
	http.ListenAndServe(":3000", n)
	return nil
}
