package http

import (
	"github.com/gorilla/websocket"
	"snake/dao"
	"snake/game"
	"snake/generators"
)

type GameConn struct {
	conn    *websocket.Conn
	level   *game.Level
	gameId  string
	perm    []int
	counter int
}

type MoveMsg struct {
	Direction string `json:"direction"`
	CloseGame bool   `json:"closeGame"`
	NewGame   bool   `json:"newGame"`
}

type EndGameAnsw struct {
	Permutation []int  `json:"permutation"`
	Counter     int    `json:"counter"`
	GameResult  string `json:"gameResult"`
	Prize       string `json:"prize"`
}

type MoveAnsw struct {
	GameMap [][]string `json:"gameMap"`
	Steps   int64      `json:"step"`
	Counter int        `json:"counter"`
}

type ErrAnsw struct {
	msg string
}

func NewGameConn(conn *websocket.Conn, gameId string) GameConn {
	gameConn := GameConn{conn: conn, gameId: gameId}
	return gameConn
}

func (gameConn *GameConn) Play() {
	defer gameConn.conn.Close()
	for {
		err := gameConn.setupGame()
		if err != nil {
			_ = gameConn.conn.WriteJSON(ErrAnsw{msg: err.Error()})
			return
		}
		_ = gameConn.conn.WriteJSON(
			MoveAnsw{
				GameMap: gameConn.level.Map(),
				Steps:   gameConn.level.Steps(),
				Counter: gameConn.counter})

		var moveMsg MoveMsg
		for {
			err = gameConn.conn.ReadJSON(&moveMsg)
			if err != nil {
				_ = gameConn.conn.WriteJSON(ErrAnsw{msg: err.Error()})
				return
			}
			moveAnsw := gameConn.handleGame(moveMsg)
			if gameConn.level.Status() != game.STATUS_UNFINISHED {
				break
			}
			_ = gameConn.conn.WriteJSON(moveAnsw)
			if moveMsg.CloseGame {
				return
			}
		}
		gameConn.handleEndGame()
		err = gameConn.conn.ReadJSON(&moveMsg)
		if err != nil {
			_ = gameConn.conn.WriteJSON(ErrAnsw{msg: err.Error()})
			return
		}
		if !moveMsg.NewGame {
			return
		}
	}
}

func (gameConn *GameConn) setupGame() error {
	dao.IncCounter(gameConn.gameId)
	_map := dao.GetMap(gameConn.gameId)
	gameConn.counter = _map.Counter
	seed, err := generators.GenerateSeed(_map.Init[:], []byte(_map.Secret), uint64(_map.Counter))
	perm := make([]int, 256)
	for i, el := range seed {
		perm[i] = int(el)
	}
	if err != nil {
		return err
	}
	level, err := generators.GenerateLevel(seed)
	if err != nil {
		return err
	}
	gameConn.level = level
	gameConn.perm = perm
	_ = gameConn.level.Step(game.DIRECTION_RIGHT)
	return nil
}

func (gameConn GameConn) handleEndGame() {
	strStatus := "win"
	if gameConn.level.Status() == game.STATUS_LOSE {
		strStatus = "lose"
	}
	_ = gameConn.conn.WriteJSON(EndGameAnsw{Permutation: gameConn.perm, Counter: gameConn.counter, GameResult: strStatus})
}

func (gameConn *GameConn) handleGame(msg MoveMsg) MoveAnsw {
	direction := gameConn.level.Snake.Direction()
	switch msg.Direction {
	case UP:
		direction = game.DIRECTION_UP
	case DOWN:
		direction = game.DIRECTION_DOWN
	case LEFT:
		direction = game.DIRECTION_LEFT
	case RIGHT:
		direction = game.DIRECTION_RIGHT
	}
	_ = gameConn.level.Step(direction)
	return MoveAnsw{GameMap: gameConn.level.Map(), Steps: gameConn.level.Steps(), Counter: gameConn.counter}
}
