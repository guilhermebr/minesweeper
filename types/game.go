package types

type Cell struct {
	Mine    bool
	Clicked bool
	Value   int
}

type CellGrid []Cell

type Game struct {
	Name   string
	Rows   int
	Cols   int
	Mines  int
	Status string
	Grid   []CellGrid
}

type GameService interface {
	Create(game Game) error
	Start(name string) error
}

type GameStore interface {
	Insert(game Game) error
	Update(game Game) error
	GetByName(name string) (Game, error)
}
