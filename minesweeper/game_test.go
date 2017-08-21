package minesweeper

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/guilhermebr/minesweeper/mocks"
	"github.com/guilhermebr/minesweeper/types"
)

func TestCreateGame(t *testing.T) {
	s := GameService{
		Store: &mocks.MockGameStore{
			OnInsert: func(game *types.Game) error {
				return nil
			},
		},
	}
	game := &types.Game{
		Name:  "mygame",
		Cols:  10,
		Rows:  11,
		Mines: 12,
	}

	if err := s.Create(game); err != nil {
		t.Fatal(err)
	}

	if game.Cols != 10 {
		t.Errorf("unexpected cols. want=10, got %d", game.Cols)
	}
	if game.Rows != 11 {
		t.Errorf("unexpected rows. want=11, got %d", game.Rows)
	}
	if game.Mines != 12 {
		t.Errorf("unexpected mines. want=12, got %d", game.Mines)
	}
	if game.Status != "new" {
		t.Errorf("unexpected status. want='new', got %d", game.Status)
	}
}
func TestCreateGame_Default(t *testing.T) {
	s := GameService{
		Store: &mocks.MockGameStore{
			OnInsert: func(game *types.Game) error {
				return nil
			},
		},
	}
	game := &types.Game{
		Name: "mygame",
	}

	if err := s.Create(game); err != nil {
		t.Fatal(err)
	}

	if game.Cols != defaultCols {
		t.Errorf("unexpected cols. want=%d, got %d", defaultCols, game.Cols)
	}
	if game.Rows != defaultRows {
		t.Errorf("unexpected rows. want=%d, got %d", defaultRows, game.Rows)
	}
	if game.Mines != defaultMines {
		t.Errorf("unexpected mines. want=%d, got %d", defaultMines, game.Mines)
	}
}

func TestCreateGame_MaxValues(t *testing.T) {
	s := GameService{
		Store: &mocks.MockGameStore{
			OnInsert: func(game *types.Game) error {
				return nil
			},
		},
	}
	game := &types.Game{
		Name:  "mygame",
		Cols:  9999,
		Rows:  9999,
		Mines: 9999,
	}

	rand.Seed(1)

	if err := s.Create(game); err != nil {
		t.Fatal(err)
	}

	if game.Cols != maxCols {
		t.Errorf("unexpected cols. want=%d, got %d", maxCols, game.Cols)
	}
	if game.Rows != maxRows {
		t.Errorf("unexpected rows. want=%d, got %d", maxRows, game.Rows)
	}
	if game.Mines != (maxCols * maxRows) {
		t.Errorf("unexpected mines. want=%d, got %d", (maxCols * maxRows), game.Mines)
	}
}

func TestStartGame(t *testing.T) {
	s := GameService{
		Store: &mocks.MockGameStore{
			OnGetByName: func(name string) (*types.Game, error) {
				return &types.Game{
					Name:  name,
					Cols:  2,
					Rows:  2,
					Mines: 1,
				}, nil
			},
			OnUpdate: func(game *types.Game) error {
				return nil
			},
		},
	}

	game, err := s.Start("mygame")
	if err != nil {
		t.Fatal(err)
	}

	if game.Status != "started" {
		t.Errorf("unexpected status. want='started', got %d", game.Status)
	}

	expected := []types.CellGrid{
		types.CellGrid{
			types.Cell{Mine: false, Clicked: false, Value: 1},
			types.Cell{Mine: true, Clicked: false, Value: 0},
		},
		types.CellGrid{
			types.Cell{Mine: false, Clicked: false, Value: 1},
			types.Cell{Mine: false, Clicked: false, Value: 1},
		},
	}
	if !reflect.DeepEqual(game.Grid, expected) {
		t.Errorf("unexpected grid. want=%v, got=%v", expected, game.Grid)
	}
}

func TestClickCell(t *testing.T) {
	s := GameService{
		Store: &mocks.MockGameStore{
			OnGetByName: func(name string) (*types.Game, error) {
				grid := []types.CellGrid{
					types.CellGrid{
						types.Cell{Mine: false, Clicked: false, Value: 1},
						types.Cell{Mine: true, Clicked: false, Value: 0},
					},
					types.CellGrid{
						types.Cell{Mine: false, Clicked: false, Value: 1},
						types.Cell{Mine: false, Clicked: false, Value: 1},
					},
				}
				return &types.Game{
					Name:   name,
					Cols:   2,
					Rows:   2,
					Mines:  1,
					Status: "started",
					Grid:   grid,
				}, nil
			},
			OnUpdate: func(game *types.Game) error {
				return nil
			},
		},
	}

	game, err := s.Click("test", 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	if game.Status != "started" {
		t.Errorf("unexpected status. want='started', got %s", game.Status)
	}

	expected := []types.CellGrid{
		types.CellGrid{
			types.Cell{Mine: false, Clicked: true, Value: 1},
			types.Cell{Mine: true, Clicked: false, Value: 0},
		},
		types.CellGrid{
			types.Cell{Mine: false, Clicked: false, Value: 1},
			types.Cell{Mine: false, Clicked: false, Value: 1},
		},
	}
	if !reflect.DeepEqual(game.Grid, expected) {
		t.Errorf("unexpected grid. want=%v, got=%v", expected, game.Grid)
	}
}

func TestClickCell_MineCell(t *testing.T) {
	s := GameService{
		Store: &mocks.MockGameStore{
			OnGetByName: func(name string) (*types.Game, error) {
				grid := []types.CellGrid{
					types.CellGrid{
						types.Cell{Mine: false, Clicked: false, Value: 1},
						types.Cell{Mine: true, Clicked: false, Value: 0},
					},
					types.CellGrid{
						types.Cell{Mine: false, Clicked: false, Value: 1},
						types.Cell{Mine: false, Clicked: false, Value: 1},
					},
				}
				return &types.Game{
					Name:   name,
					Cols:   2,
					Rows:   2,
					Mines:  1,
					Status: "started",
					Grid:   grid,
				}, nil
			},
			OnUpdate: func(game *types.Game) error {
				return nil
			},
		},
	}

	game, err := s.Click("test", 0, 1)
	if err != nil {
		t.Fatal(err)
	}

	if game.Status != "over" {
		t.Errorf("unexpected status. want='started', got %s", game.Status)
	}

	expected := []types.CellGrid{
		types.CellGrid{
			types.Cell{Mine: false, Clicked: false, Value: 1},
			types.Cell{Mine: true, Clicked: true, Value: 0},
		},
		types.CellGrid{
			types.Cell{Mine: false, Clicked: false, Value: 1},
			types.Cell{Mine: false, Clicked: false, Value: 1},
		},
	}
	if !reflect.DeepEqual(game.Grid, expected) {
		t.Errorf("unexpected grid. want=%v, got=%v", expected, game.Grid)
	}
}

func TestClickCell_Won(t *testing.T) {
	s := GameService{
		Store: &mocks.MockGameStore{
			OnGetByName: func(name string) (*types.Game, error) {
				grid := []types.CellGrid{
					types.CellGrid{
						types.Cell{Mine: false, Clicked: true, Value: 1},
						types.Cell{Mine: true, Clicked: false, Value: 0},
					},
					types.CellGrid{
						types.Cell{Mine: false, Clicked: true, Value: 1},
						types.Cell{Mine: false, Clicked: false, Value: 1},
					},
				}
				return &types.Game{
					Name:   name,
					Cols:   2,
					Rows:   2,
					Mines:  1,
					Clicks: 2,
					Status: "started",
					Grid:   grid,
				}, nil
			},
			OnUpdate: func(game *types.Game) error {
				return nil
			},
		},
	}

	game, err := s.Click("test", 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	if game.Status != "won" {
		t.Errorf("unexpected status. want='won', got %s", game.Status)
	}

	expected := []types.CellGrid{
		types.CellGrid{
			types.Cell{Mine: false, Clicked: true, Value: 1},
			types.Cell{Mine: true, Clicked: false, Value: 0},
		},
		types.CellGrid{
			types.Cell{Mine: false, Clicked: true, Value: 1},
			types.Cell{Mine: false, Clicked: true, Value: 1},
		},
	}
	if !reflect.DeepEqual(game.Grid, expected) {
		t.Errorf("unexpected grid. want=%v, got=%v", expected, game.Grid)
	}
}
