package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/guilhermebr/minesweeper/mocks"
	"github.com/guilhermebr/minesweeper/types"
	"github.com/sirupsen/logrus"
)

func TestCreateGame_Success(t *testing.T) {
	log := logrus.StandardLogger()
	services := &Services{
		logger: log,
		GameService: &mocks.MockGameService{
			OnCreate: func(game *types.Game) error {
				if game.Name != "teste" {
					t.Fatalf("unexpected name. want=teste, got=%s", game.Name)
				}
				if game.Rows != 10 {
					t.Fatalf("unexpected rows. want=10, got=%d", game.Rows)
				}
				if game.Cols != 12 {
					t.Fatalf("unexpected cols. want=12, got=%d", game.Cols)
				}
				if game.Mines != 30 {
					t.Fatalf("unexpected mines. want=30, got=%d", game.Mines)
				}
				game.Status = "new"
				return nil
			},
		},
	}

	data := `{"name":"teste","rows": 10, "cols":12, "mines": 30}`
	b := strings.NewReader(data)
	req, err := http.NewRequest("POST", "/game", b)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router(services).ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: want %v, got %v",
			http.StatusCreated, status)
	}

	// Check the response body.
	expected := `{"success":true,"status":201,"result":{"name":"teste","rows":10,"cols":12,"mines":30,"status":"new"}}`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: want %v, got %v",
			expected, rr.Body.String())
	}
}

func TestCreateGame_InvalidJson(t *testing.T) {
	log := logrus.StandardLogger()
	services := &Services{
		logger: log,
		GameService: &mocks.MockGameService{
			OnCreate: func(game *types.Game) error {
				t.Fatal("game create should not be called")
				return nil
			},
		},
	}

	data := `{name:teste}`
	b := strings.NewReader(data)
	req, err := http.NewRequest("POST", "/game", b)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router(services).ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: want %v, got %v",
			http.StatusBadRequest, status)
	}

	// Check the response body.
	expected := `{"type":"invalid_json","message":"Invalid or malformed JSON"}`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: want %v, got %v",
			expected, rr.Body.String())
	}
}

func TestCreateGame_ServerError(t *testing.T) {
	log := logrus.StandardLogger()
	services := &Services{
		logger: log,
		GameService: &mocks.MockGameService{
			OnCreate: func(game *types.Game) error {
				return errors.New("some error")
			},
		},
	}

	data := `{"name":"teste"}`
	b := strings.NewReader(data)
	req, err := http.NewRequest("POST", "/game", b)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router(services).ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: want %v, got %v",
			http.StatusInternalServerError, status)
	}

	// Check the response body.
	expected := `{"type":"server_error","message":"Internal server error. The error has been logged and we are working on it"}`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: want %v, got %v",
			expected, rr.Body.String())
	}
}

func TestStartGame_Success(t *testing.T) {
	log := logrus.StandardLogger()
	services := &Services{
		logger: log,
		GameService: &mocks.MockGameService{
			OnStart: func(name string) (*types.Game, error) {
				if name != "teste" {
					t.Fatalf("unexpected name. want=teste, got=%s", name)
				}

				grid := []types.CellGrid{
					types.CellGrid{
						types.Cell{Mine: false, Clicked: false, Value: 1},
						types.Cell{Mine: false, Clicked: false, Value: 1},
					},
					types.CellGrid{
						types.Cell{Mine: false, Clicked: false, Value: 1},
						types.Cell{Mine: true, Clicked: false, Value: 0},
					},
				}

				return &types.Game{
					Name:   "teste",
					Status: "started",
					Rows:   2,
					Cols:   2,
					Mines:  4,
					Grid:   grid,
				}, nil
			},
		},
	}

	req, err := http.NewRequest("POST", "/game/teste/start", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router(services).ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: want %v, got %v",
			http.StatusCreated, status)
	}

	// Check the response body.
	expected := `{"success":true,"status":200,"result":{"name":"teste","rows":2,"cols":2,"mines":4,"status":"started"}}`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: want %v, got %v",
			expected, rr.Body.String())
	}
}

func TestStartGame_ServerError(t *testing.T) {
	log := logrus.StandardLogger()
	services := &Services{
		logger: log,
		GameService: &mocks.MockGameService{
			OnStart: func(name string) (*types.Game, error) {
				return nil, errors.New("error")
			},
		},
	}

	req, err := http.NewRequest("POST", "/game/teste/start", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router(services).ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: want %v, got %v",
			http.StatusInternalServerError, status)
	}

	// Check the response body.
	expected := `{"type":"server_error","message":"Internal server error. The error has been logged and we are working on it"}`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: want %v, got %v",
			expected, rr.Body.String())
	}
}

func TestClickCell_Success(t *testing.T) {
	log := logrus.StandardLogger()
	services := &Services{
		logger: log,
		GameService: &mocks.MockGameService{
			OnClick: func(name string, i, j int) (*types.Game, error) {
				grid := []types.CellGrid{
					types.CellGrid{
						types.Cell{Mine: false, Clicked: false, Value: 1},
						types.Cell{Mine: false, Clicked: false, Value: 1},
					},
					types.CellGrid{
						types.Cell{Mine: false, Clicked: false, Value: 1},
						types.Cell{Mine: true, Clicked: false, Value: 0},
					},
				}

				return &types.Game{
					Name:   name,
					Status: "started",
					Rows:   2,
					Cols:   2,
					Mines:  1,
					Grid:   grid,
				}, nil
			},
		},
	}

	data := `{"name":"teste","row": 0, "col":1}`
	b := strings.NewReader(data)
	req, err := http.NewRequest("POST", "/game/teste/click", b)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router(services).ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: want %v, got %v",
			http.StatusOK, status)
	}

	// Check the response body.
	expected := `{"success":true,"status":200,"result":{"Cell":{"mine":false,"clicked":false,"value":1},"Game":{"name":"teste","rows":2,"cols":2,"mines":1,"status":"started"}}}`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: want %v, got %v",
			expected, rr.Body.String())
	}
}

func TestClickCell_GameOver(t *testing.T) {
	log := logrus.StandardLogger()
	services := &Services{
		logger: log,
		GameService: &mocks.MockGameService{
			OnClick: func(name string, i, j int) (*types.Game, error) {
				grid := []types.CellGrid{
					types.CellGrid{
						types.Cell{Mine: false, Clicked: false, Value: 1},
						types.Cell{Mine: false, Clicked: false, Value: 1},
					},
					types.CellGrid{
						types.Cell{Mine: false, Clicked: false, Value: 1},
						types.Cell{Mine: true, Clicked: false, Value: 0},
					},
				}

				return &types.Game{
					Name:   name,
					Status: "over",
					Rows:   2,
					Cols:   2,
					Mines:  1,
					Grid:   grid,
				}, nil
			},
		},
	}

	data := `{"name":"teste","row": 1, "col":1}`
	b := strings.NewReader(data)
	req, err := http.NewRequest("POST", "/game/teste/click", b)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router(services).ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: want %v, got %v",
			http.StatusOK, status)
	}

	// Check the response body.
	expected := `{"success":true,"status":200,"result":{"Cell":{"mine":true,"clicked":false,"value":0},"Game":{"name":"teste","rows":2,"cols":2,"mines":1,"status":"over","grid":[[{"mine":false,"clicked":false,"value":1},{"mine":false,"clicked":false,"value":1}],[{"mine":false,"clicked":false,"value":1},{"mine":true,"clicked":false,"value":0}]]}}}`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: want %v, got %v",
			expected, rr.Body.String())
	}
}

func TestClickCell_InvalidJson(t *testing.T) {
	log := logrus.StandardLogger()
	services := &Services{
		logger: log,
		GameService: &mocks.MockGameService{
			OnClick: func(name string, i, j int) (*types.Game, error) {
				t.Fatal("game click should not be called")
				return nil, nil
			},
		},
	}

	data := `"name":"teste"`
	b := strings.NewReader(data)
	req, err := http.NewRequest("POST", "/game/teste/click", b)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router(services).ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: want %v, got %v",
			http.StatusBadRequest, status)
	}

	// Check the response body.
	expected := `{"type":"invalid_json","message":"Invalid or malformed JSON"}`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: want %v, got %v",
			expected, rr.Body.String())
	}
}

func TestClickCell_ServerError(t *testing.T) {
	log := logrus.StandardLogger()
	services := &Services{
		logger: log,
		GameService: &mocks.MockGameService{
			OnClick: func(name string, i, j int) (*types.Game, error) {
				return nil, errors.New("error")
			},
		},
	}

	data := `{"name":"teste","row": 2, "col":2}`
	b := strings.NewReader(data)
	req, err := http.NewRequest("POST", "/game/teste/click", b)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router(services).ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: want %v, got %v",
			http.StatusInternalServerError, status)
	}

	// Check the response body.
	expected := `{"type":"server_error","message":"Internal server error. The error has been logged and we are working on it"}`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: want %v, got %v",
			expected, rr.Body.String())
	}
}
