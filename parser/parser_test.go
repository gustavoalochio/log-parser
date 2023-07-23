package parser

import (
	"log-parser/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_ParserInitGame(t *testing.T) {
	parser := NewParser()

	t.Run("success", func(t *testing.T) {
		//// arrange
		str := ` 20:37 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\bot_minplayers\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`
		//// action
		result := parser.ParserInitGame(str)
		//// assert
		assert.Equal(t, result, true)
	})

	t.Run("error", func(t *testing.T) {
		//// arrange
		str := ` 20:37 Init1Game: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\bot_minplayers\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`
		//// action
		result := parser.ParserInitGame(str)
		//// assert
		assert.Equal(t, result, false)
	})
}

func TestParser_ParserInterval(t *testing.T) {
	parser := NewParser()

	t.Run("success", func(t *testing.T) {
		//// arrange
		str := ` 20:37 ------------------------------------------------------------`
		//// action
		result := parser.ParserInterval(str)
		//// assert
		assert.Equal(t, result, true)
	})

	t.Run("error", func(t *testing.T) {
		//// arrange
		str := ` 20:37 -1-----------------------------------------------------------`
		//// action
		result := parser.ParserInterval(str)
		//// assert
		assert.Equal(t, result, false)
	})
}

func TestParser_ParserClientConnect(t *testing.T) {
	parser := NewParser()

	t.Run("success", func(t *testing.T) {
		//// arrange
		str := ` 21:51 ClientConnect: 3`
		//// action
		result := parser.ParserClientConnect(str)
		//// assert
		assert.Equal(t, result, true)
	})

	t.Run("error", func(t *testing.T) {
		//// arrange
		str := ` 21:51 ClientConnect:`
		//// action
		result := parser.ParserClientConnect(str)
		//// assert
		assert.Equal(t, result, false)
	})
}

func TestParser_ParserClientUserInfoChanged(t *testing.T) {
	parser := NewParser()

	t.Run("success", func(t *testing.T) {
		//// arrange
		str := ` 21:51 ClientUserinfoChanged: 3 n\Dono da Bola\t\0\model\sarge/krusade\hmodel\sarge/krusade\g_redteam\\g_blueteam\\c1\5\c2\5\hc\95\w\0\l\0\tt\0\tl\0`
		//// action
		result := parser.ParserClientUserInfoChanged(str)
		//// assert
		assert.Equal(t, result, true)
	})

	t.Run("error", func(t *testing.T) {
		//// arrange
		str := ` 21:51 ClientUserinfoChanged:  n\Dono da Bola\t\0\model\sarge/krusade\hmodel\sarge/krusade\g_redteam\\g_blueteam\\c1\5\c2\5\hc\95\w\0\l\0\tt\0\tl\0`
		//// action
		result := parser.ParserClientUserInfoChanged(str)
		//// assert
		assert.Equal(t, result, false)
	})
}

func TestParser_ParserKilled(t *testing.T) {
	parser := NewParser()

	t.Run("success", func(t *testing.T) {
		//// arrange
		str := ` 22:06 Kill: 2 3 7: Isgalamido killed Mocinha by MOD_ROCKET_SPLASH`
		//// action
		result := parser.ParserKilled(str)
		//// assert
		assert.Equal(t, result, true)
	})

	t.Run("error", func(t *testing.T) {
		//// arrange
		str := ` 22:06 Kill: 2 3 7: Isgalamidokilled Mocinha by MOD_ROCKET_SPLASH`
		//// action
		result := parser.ParserKilled(str)
		//// assert
		assert.Equal(t, result, false)
	})
}

func TestParser_GetInformationPlayerUpdate(t *testing.T) {
	parser := NewParser()
	playerInformation := entity.UpdatePlayerInformation{
		PlayerId: "2",
		Name:     "Isgalamido",
	}

	t.Run("success", func(t *testing.T) {
		//// arrange
		str := ` 21:15 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\uriel/zael\hmodel\uriel/zael\g_redteam\\g_blueteam\\c1\5\c2\5\hc\100\w\0\l\0\tt\0\tl\0`
		//// action
		result, _ := parser.GetInformationPlayerUpdate(str)
		//// assert
		assert.Equal(t, result, playerInformation)
	})

	t.Run("error", func(t *testing.T) {
		//// arrange
		str := ` 21:15 ClientUserinfoChanged:2 n\Isgalamido\t\0\model\uriel/zael\hmodel\uriel/zael\g_redteam\\g_blueteam\\c1\5\c2\5\hc\100\w\0\l\0\tt\0\tl\0`
		//// action
		_, err := parser.GetInformationPlayerUpdate(str)
		//// assert
		assert.NotNil(t, err)
	})
}

func TestParser_GetInformationPlayerConnected(t *testing.T) {
	parser := NewParser()
	playerInformation := entity.PlayerConnectedInformation{
		PlayerId: "2",
	}
	t.Run("success", func(t *testing.T) {
		//// arrange
		str := ` 21:15 ClientConnect: 2`
		//// action
		result, _ := parser.GetInformationPlayerConnected(str)
		//// assert
		assert.Equal(t, result, playerInformation)
	})

	t.Run("error", func(t *testing.T) {
		//// arrange
		str := ` 21:15 ClientConnect:`
		//// action
		_, err := parser.GetInformationPlayerConnected(str)
		//// assert
		assert.NotNil(t, err)
	})
}

func TestParser_GetInformationKilledPlayer(t *testing.T) {
	parser := NewParser()
	playerInformation := entity.PlayerKilledInformation{
		KillerName: "Isgalamido",
		DeadName:   "Mocinha",
		Reason:     "MOD_ROCKET_SPLASH",
	}
	t.Run("success", func(t *testing.T) {
		//// arrange
		str := ` 22:40 Kill: 2 2 7: Isgalamido killed Mocinha by MOD_ROCKET_SPLASH`
		//// action
		result, _ := parser.GetInformationKilledPlayer(str)
		//// assert
		assert.Equal(t, result, playerInformation)
	})

	t.Run("error", func(t *testing.T) {
		//// arrange
		str := ` 22:40 Kill: 2 2 7: Isgalamido killedMocinha by MOD_ROCKET_SPLASH`
		//// action
		_, err := parser.GetInformationKilledPlayer(str)
		//// assert
		assert.NotNil(t, err)
	})
}
