package minesweeper

import (
	"errors"

	"github.com/guilhermebr/minesweeper/types"
)

type GameService struct {
	Store types.GameStore
}

const (
	max_rows = 30
	max_cols = 30
)

func (s *GameService) Create(game types.Game) error {
	if game.Name == "" {
		return errors.New("no Game name")
	}

	if game.Rows == 0 {
		game.Rows = 6
	}
	if game.Cols == 0 {
		game.Cols = 6
	}
	if game.Mines == 0 {
		game.Mines = 12
	}

	if game.Rows > max_rows {
		game.Rows = max_rows
	}
	if game.Cols > max_cols {
		game.Cols = max_cols
	}
	if game.Mines > (game.Cols * game.Rows) {
		game.Mines = (game.Cols * game.Rows)
	}
	game.Status = "new"

	err := s.Store.Insert(game)
	return err
}

func (s *GameService) Start(name string) error {
	game, err := s.Store.GetByName(name)
	if err != nil {
		return err
	}

	buildBoard(&game)

	game.Status = "started"
	err = s.Store.Update(game)
	return err
}
