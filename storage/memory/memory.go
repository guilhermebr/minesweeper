package memory

import "github.com/guilhermebr/minesweeper/types"

type DB struct {
	games map[string]*types.Game
}

func New() *DB {
	return &DB{
		games: make(map[string]*types.Game),
	}
}
