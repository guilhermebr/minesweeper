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

	if err := s.GameService.Create(&game); err != nil {
		log.WithField("err", err).Error("cannot create game")
		ErrInternalServer.Send(w)
		return
	}
	Success(game, http.StatusCreated).Send(w)
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

	game, err := s.GameService.Start(name)
	if err != nil {
		log.WithField("err", err).Error("cannot start game")
		ErrInternalServer.Send(w)
		return
	}

	game2 := *game
	game2.Grid = nil

	Success(game2, http.StatusOK).Send(w)
}

// title: cell click
// path: /game/{name}/click
// method: POST
// responses:
//   200: OK
//   400: Invalid json
//   500: server error
func (s *Services) clickCell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	log := s.logger.WithFields(logrus.Fields{
		"service": "game",
		"method":  "click",
	})

	var cellPos struct {
		Row int `json:"row"`
		Col int `json:"col"`
	}

	if err := json.NewDecoder(r.Body).Decode(&cellPos); err != nil {
		log.Error(err)
		ErrInvalidJSON.Send(w)
		return
	}

	game, err := s.GameService.Click(name, cellPos.Row, cellPos.Col)
	if err != nil {
		log.WithField("err", err).Error("cannot click cell")
		ErrInternalServer.Send(w)
		return
	}
	cell := game.Grid[cellPos.Row][cellPos.Col]

	if game.Status != "over" && game.Status != "won" {
		game.Grid = nil
	}

	var result struct {
		Cell types.Cell
		Game types.Game
	}

	result.Cell = cell
	result.Game = *game

	Success(&result, http.StatusOK).Send(w)
}
