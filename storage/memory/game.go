package memory

import (
	"errors"

	"github.com/guilhermebr/minesweeper/types"
)

type GameStore struct {
	db *DB
}

func NewGameStore(db *DB) *GameStore {
	return &GameStore{db: db}
}

func (s *GameStore) Insert(game types.Game) error {
	if _, ok := s.db.games[game.Name]; ok {
		return errors.New("game already exist")
	}
	s.db.games[game.Name] = game
	return nil
}

func (s *GameStore) Update(game types.Game) error {
	if _, ok := s.db.games[game.Name]; !ok {
		return errors.New("game do not exist")
	}
	s.db.games[game.Name] = game
	return nil
}

func (s *GameStore) GetByName(name string) (types.Game, error) {
	if game, ok := s.db.games[name]; ok {
		return game, nil
	}
	return types.Game{}, errors.New("game not found")
}
