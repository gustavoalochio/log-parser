package engine

import (
	"log-parser/entity"
	"log-parser/parser/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestEngine_Initgame(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	parser := mocks.NewMockParser(ctrl)

	t.Run("init", func(t *testing.T) {
		//// arrange
		engine := NewEngine(parser)
		//// action
		engineResult := engine.InitGame()
		//// assert
		assert.NotNil(t, engineResult.Game)
		assert.Equal(t, engineResult.Game.Running, true)
		assert.Equal(t, engineResult.Game.Running, true)
		assert.NotEmpty(t, engineResult.Game.Players)
		assert.NotEmpty(t, engineResult.Game.PlayersById)
		assert.NotEmpty(t, engineResult.Game.PlayersByName)
	})
}

func TestEngine_PrintResultGame(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	parser := mocks.NewMockParser(ctrl)

	t.Run("equal", func(t *testing.T) {
		//// arrange
		engine := NewEngine(parser)
		gameInfo := entity.ResultGameInformation{
			TotalKills: 0,
			Players:    []string{"Isgalamido"},
			Kills:      map[string]int{"Isgalamido": 0},
		}
		gameCounter := 1

		successResult := `{"game_1":{"total_kills":0,"players":["Isgalamido"],"kills":{"Isgalamido":0}}}`
		//// action
		result := engine.BuildResultGame(gameInfo, gameCounter)
		//// assert
		assert.Equal(t, result, successResult)
	})

	t.Run("not-equal", func(t *testing.T) {
		//// arrange
		engine := NewEngine(parser)
		gameInfo := entity.ResultGameInformation{
			TotalKills: 1,
			Players:    []string{"Isgalamido"},
			Kills:      map[string]int{"Isgalamido": 1},
		}
		gameCounter := 1

		successResult := `{"game_1":{"total_kills":0,"players":["Isgalamido"],"kills":{"Isgalamido":0}}}`
		//// action
		result := engine.BuildResultGame(gameInfo, gameCounter)
		//// assert
		assert.NotEqual(t, result, successResult)
	})
}

func TestEngine_PrintDeadReasons(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	parser := mocks.NewMockParser(ctrl)

	t.Run("equal", func(t *testing.T) {
		//// arrange
		engine := NewEngine(parser)
		game := entity.Game{
			Running:           true,
			Players:           []*entity.Player{},
			PlayersByName:     make(map[string]*entity.Player),
			PlayersById:       make(map[string]*entity.Player),
			DeadReasons:       map[string]int{"MOD_ROCKET": 1, "MOD_ROCKET_SPLASH": 2},
			DefaultPlayerName: "<world>",
		}

		gameCounter := 20
		successResult := `{"game-20":{"MOD_ROCKET":1,"MOD_ROCKET_SPLASH":2}}`
		//// action
		result := engine.BuildDeadReasons(&game, gameCounter)
		//// assert
		assert.Equal(t, result, successResult)
	})

	t.Run("not-equal", func(t *testing.T) {
		//// arrange
		engine := NewEngine(parser)
		game := entity.Game{
			Running:           true,
			Players:           []*entity.Player{},
			PlayersByName:     make(map[string]*entity.Player),
			PlayersById:       make(map[string]*entity.Player),
			DeadReasons:       map[string]int{"MOD_ROCKET": 1, "MOD_ROCKET_SPLASH": 2},
			DefaultPlayerName: "<world>",
		}

		gameCounter := 20
		successResult := `{"game-20":{"MOD_ROCKET":0,"MOD_ROCKET_SPLASH":2}}`
		//// action
		result := engine.BuildDeadReasons(&game, gameCounter)
		//// assert
		assert.NotEqual(t, result, successResult)
	})

}

func TestEngine_Kill(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	parser := mocks.NewMockParser(ctrl)

	t.Run("KILL +1", func(t *testing.T) {
		//// arrange
		engine := NewEngine(parser)

		game := &entity.Game{
			Running:           true,
			Players:           []*entity.Player{},
			PlayersByName:     make(map[string]*entity.Player),
			PlayersById:       make(map[string]*entity.Player),
			DeadReasons:       map[string]int{},
			DefaultPlayerName: "<world>",
		}

		game.AddPlayer(&entity.Player{ID: "0", Name: game.DefaultPlayerName})
		game.AddPlayer(&entity.Player{ID: "2", Name: "Isgalamido"})
		game.AddPlayer(&entity.Player{ID: "3", Name: "Mocinha"})

		killedInfo := entity.PlayerKilledInformation{
			KillerName: "Isgalamido",
			DeadName:   "Mocinha",
			Reason:     "MOD_ROCKET_SPLASH",
		}

		//// action
		engine.Kill(game, killedInfo)

		//// assert
		assert.Equal(t, game.PlayersByName["Isgalamido"].Kills, 1)
	})

	t.Run("KILL -1", func(t *testing.T) {
		//// arrange
		engine := NewEngine(parser)

		game := &entity.Game{
			Running:           true,
			Players:           []*entity.Player{},
			PlayersByName:     make(map[string]*entity.Player),
			PlayersById:       make(map[string]*entity.Player),
			DeadReasons:       map[string]int{},
			DefaultPlayerName: "<world>",
		}

		game.AddPlayer(&entity.Player{ID: "0", Name: game.DefaultPlayerName})
		game.AddPlayer(&entity.Player{ID: "2", Name: "Isgalamido"})
		game.AddPlayer(&entity.Player{ID: "3", Name: "Mocinha"})

		killedInfo := entity.PlayerKilledInformation{
			KillerName: "<world>",
			DeadName:   "Isgalamido",
			Reason:     "MOD_TRIGGER_HURT",
		}

		//// action
		engine.Kill(game, killedInfo)

		//// assert
		assert.Equal(t, game.PlayersByName["Isgalamido"].Kills, -1)
	})
}
