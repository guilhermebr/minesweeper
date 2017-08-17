package types

type Game struct {
	Name   string
	Rows   int
	Cols   int
	Mines  int
	Status string
}

type GameService interface {
	Create(game Game) error
}

type GameStore interface {
	Insert(game Game) error
}
