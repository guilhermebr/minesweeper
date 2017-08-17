package minesweeper

import (
	"errors"

	"github.com/guilhermebr/minesweeper/types"
)

type GameService struct {
	Store types.GameStore
}

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
	game.Status = "new"

	err := s.Store.Insert(game)
	return err
}
