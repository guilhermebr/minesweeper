package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guilhermebr/minesweeper/types"
	"github.com/sirupsen/logrus"
)

// title: create game
// path: /game
// method: POST
// responses:
//   201: Game created
//   400: Invalid json
//	 500: server error
func (s *Services) createGame(w http.ResponseWriter, r *http.Request) {
	var game types.Game

	log := s.logger.WithFields(logrus.Fields{
		"service": "game",
		"method":  "create",
	})

	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		log.Error(err)
		ErrInvalidJSON.Send(w)
		return
	}

	if err := s.GameService.Create(game); err != nil {
		log.WithField("err", err).Error("cannot create game")
		ErrInternalServer.Send(w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// title: start game
// path: /game/{name}/start
// method: POST
// responses:
//   200: OK
//   500: server error
func (s *Services) startGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	log := s.logger.WithFields(logrus.Fields{
		"service": "game",
		"method":  "start",
	})

	if err := s.GameService.Start(name); err != nil {
		log.WithField("err", err).Error("cannot start game")
		ErrInternalServer.Send(w)
		return
	}
	w.WriteHeader(http.StatusOK)
}
