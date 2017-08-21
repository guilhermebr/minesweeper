package mocks

import "github.com/guilhermebr/minesweeper/types"

type MockGameService struct {
	OnCreate func(game *types.Game) error
	OnStart  func(name string) (*types.Game, error)
	OnClick  func(name string, i, j int) (*types.Game, error)
}

func (m *MockGameService) Create(game *types.Game) error {
	return m.OnCreate(game)
}

func (m *MockGameService) Start(name string) (*types.Game, error) {
	return m.OnStart(name)
}

func (m *MockGameService) Click(name string, i, j int) (*types.Game, error) {
	return m.OnClick(name, i, j)
}

type MockGameStore struct {
	OnInsert    func(game *types.Game) error
	OnUpdate    func(game *types.Game) error
	OnGetByName func(name string) (*types.Game, error)
}

func (m *MockGameStore) Insert(game *types.Game) error {
	return m.OnInsert(game)
}

func (m *MockGameStore) Update(game *types.Game) error {
	return m.OnUpdate(game)
}

func (m *MockGameStore) GetByName(name string) (*types.Game, error) {
	return m.OnGetByName(name)
}
