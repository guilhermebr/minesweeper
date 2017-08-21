package types

type Cell struct {
	Mine    bool `json:"mine"`
	Clicked bool `json:"clicked"`
	Value   int  `json:"value"`
}

type CellGrid []Cell

type Game struct {
	Name   string     `json:"name"`
	Rows   int        `json:"rows"`
	Cols   int        `json:"cols"`
	Mines  int        `json:"mines"`
	Status string     `json:"status"`
	Grid   []CellGrid `json:"grid,omitempty"`
	Clicks int        `json:"-"`
}

type GameService interface {
	Create(game *Game) error
	Start(name string) (*Game, error)
	Click(name string, i, j int) (*Game, error)
}

type GameStore interface {
	Insert(game *Game) error
	Update(game *Game) error
	GetByName(name string) (*Game, error)
}
