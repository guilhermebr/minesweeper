package minesweeper

import (
	"errors"
	"math/rand"
	"time"

	"github.com/guilhermebr/minesweeper/types"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func buildBoard(game *types.Game) {
	numCells := game.Cols * game.Rows
	cells := make(types.CellGrid, numCells)

	// Randomly set mines
	i := 0
	for i < game.Mines {
		idx := rand.Intn(numCells)
		if !cells[idx].Mine {
			cells[idx].Mine = true
			i++
		}
	}

	game.Grid = make([]types.CellGrid, game.Rows)
	for row := range game.Grid {
		game.Grid[row] = cells[(game.Cols * row):(game.Cols * (row + 1))]
	}

	// Set cell values
	for i, row := range game.Grid {
		for j, cell := range row {
			if cell.Mine {
				setAdjacentValues(game, i, j)
			}
		}
	}
}

func setAdjacentValues(game *types.Game, i, j int) {
	for z := i - 1; z < i+2; z++ {
		if z < 0 || z > game.Rows-1 {
			continue
		}
		for w := j - 1; w < j+2; w++ {
			if w < 0 || w > game.Cols-1 {
				continue
			}
			if z == i && w == j {
				continue
			}
			game.Grid[z][w].Value++
		}
	}
}

func clickCell(game *types.Game, i, j int) error {
	if game.Grid[i][j].Clicked {
		return errors.New("cell already clicked")
	}
	game.Grid[i][j].Clicked = true
	if game.Grid[i][j].Mine {
		game.Status = "over"
		return nil
	}
	game.Clicks += 1
	if checkWon(game) {
		game.Status = "won"
	}

	return nil
}

func checkWon(game *types.Game) bool {
	return game.Clicks == ((game.Rows * game.Cols) - game.Mines)
}
