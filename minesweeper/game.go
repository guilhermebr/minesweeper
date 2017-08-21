package minesweeper

import (
	"errors"
	"fmt"

	"github.com/guilhermebr/minesweeper/types"
)

type GameService struct {
	Store types.GameStore
}

const (
	defaultRows  = 6
	defaultCols  = 6
	defaultMines = 12
	maxRows      = 30
	maxCols      = 30
)

func (s *GameService) Create(game *types.Game) error {
	if game.Name == "" {
		return errors.New("no Game name")
	}

	if game.Rows == 0 {
		game.Rows = defaultRows
	}
	if game.Cols == 0 {
		game.Cols = defaultCols
	}
	if game.Mines == 0 {
		game.Mines = defaultMines
	}

	if game.Rows > maxRows {
		game.Rows = maxRows
	}
	if game.Cols > maxCols {
		game.Cols = maxCols
	}
	if game.Mines > (game.Cols * game.Rows) {
		game.Mines = (game.Cols * game.Rows)
	}
	game.Status = "new"

	err := s.Store.Insert(game)
	return err
}

func (s *GameService) Start(name string) (*types.Game, error) {
	game, err := s.Store.GetByName(name)
	if err != nil {
		return nil, err
	}

	buildBoard(game)

	game.Status = "started"
	err = s.Store.Update(game)
	fmt.Printf("%#v\n", game.Grid)
	return game, err
}

func (s *GameService) Click(name string, i, j int) (*types.Game, error) {
	game, err := s.Store.GetByName(name)
	if err != nil {
		return nil, err
	}

	if err := clickCell(game, i, j); err != nil {
		return nil, err
	}

	if err := s.Store.Update(game); err != nil {
		return nil, err
	}

	return game, nil
}
