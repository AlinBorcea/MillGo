package mill

import (
	"errors"
)

type PlayerId int
type GameStatus int

const (
	PlayerNone PlayerId = iota
	PlayerOne
	PlayerTwo
)

const (
	StatusDefault GameStatus = iota
	StatusTurnDone
	StatusAwaitTargetMan
)

var (
	Success         error = nil
	ErrBadInput           = errors.New("input is invalid")
	ErrNoMenLeft          = errors.New("no men left")
	ErrItIsAMill          = errors.New("it is a mill")
	ErrItIsNotAMill       = errors.New("it is not a mill")
	ErrFail               = errors.New("fail")
)

type Mill struct {
	board           [3][8]PlayerId
	players         [3]*Player
	currentPlayerId PlayerId
	status          GameStatus

	menLeftToPlacePlayerOne int
	menLeftToPlacePlayerTwo int
	menOnBoardPlayerOne     int
	menOnBoardPlayerTwo     int
}

func NewMill() (m *Mill) {
	return &Mill{
		players:         [3]*Player{NewPlayerNone(), NewPlayer(), NewPlayer()},
		currentPlayerId: PlayerOne,
		status:          StatusDefault,

		menLeftToPlacePlayerOne: 9,
		menLeftToPlacePlayerTwo: 9,
		menOnBoardPlayerOne:     0,
		menOnBoardPlayerTwo:     0,
	}
}

func (m *Mill) Status() GameStatus {
	return m.status
}

func (m *Mill) PlaceMan(a, b int) *error {
	player := m.currentPlayer()

	if !player.hasMenLeft() {
		return &ErrNoMenLeft
	}

	if !m.placeCellUnrestricted(a, b) {
		return &ErrBadInput
	}

	if m.isMill(a, b) {
		m.status = StatusAwaitTargetMan
	} else {
		m.status = StatusTurnDone
	}

	player.decreaseMenLeft()
	player.increaseMenOnBoard()
	return &Success
}

func (m *Mill) MoveMan(a, b, c, d int) *error {
	if !m.moveCellToNeighbor(a, b, c, d) {
		return &ErrBadInput
	}

	if m.isMill(c, d) {
		m.status = StatusAwaitTargetMan
	} else {
		m.status = StatusTurnDone
	}

	return &Success
}

func (m *Mill) NextPlayer() {
	if m.currentPlayerId == PlayerOne {
		m.currentPlayerId = PlayerTwo
	} else {
		m.currentPlayerId = PlayerOne
	}
	m.status = StatusDefault
}

func (m *Mill) TakeManFromOpponent(a, b int) *error {
	if m.isMill(a, b) {
		return &ErrItIsAMill
	}

	if m.board[a][b] == m.currentPlayerId || m.board[a][b] == PlayerNone {
		return &ErrBadInput
	}

	m.board[a][b] = PlayerNone
	m.opponent().decreaseMenOnBoard()
	return &Success
}

func (m *Mill) EnemyHasVulnerableMan() bool {
	for i := range m.board {
		for j := range m.board[i] {
			if m.board[i][j] != m.currentPlayerId && m.isMill(i, j) {
				return true
			}
		}
	}
	return false
}

func (m *Mill) currentPlayer() *Player {
	return m.players[m.currentPlayerId]
}

func (m *Mill) opponent() *Player {
	if m.currentPlayerId == PlayerOne {
		return m.players[PlayerTwo]
	} else {
		return m.players[PlayerOne]
	}
}

func (m *Mill) placeCellUnrestricted(a, b int) bool {
	if m.board[a][b] == PlayerNone {
		m.board[a][b] = m.currentPlayerId
		return true
	}
	return false
}

func (m *Mill) moveCellToNeighbor(a, b, c, d int) bool {
	if m.board[a][b] != m.currentPlayerId {
		return false
	}

	if m.board[c][d] != PlayerNone {
		return false
	}

	if a == c && (b-d == 1 || d-b == 1) {
		goto ok
	} else if (b == 1 || b == 3 || b == 5 || b == 7) && b == d && (a-c == 1 || c-a == 1) {
		goto ok
	} else {
		goto fail
	}

fail:
	return false

ok:
	m.board[c][d] = m.board[a][b]
	m.board[a][b] = PlayerNone
	return true
}

func (m *Mill) isMill(a, b int) bool {
	if m.board[a][b] == PlayerNone {
		return false
	}

	if b >= 0 && b <= 2 {
		if m.board[a][0] == m.board[a][1] && m.board[a][2] == m.board[a][b] {
			return true
		}
	}
	if b >= 2 && b <= 4 {
		if m.board[a][2] == m.board[a][3] && m.board[a][4] == m.board[a][b] {
			return true
		}
	}
	if b >= 4 && b <= 6 {
		if m.board[a][4] == m.board[a][5] && m.board[a][6] == m.board[a][b] {
			return true
		}
	}
	if b >= 6 && b <= 7 {
		if m.board[a][6] == m.board[a][7] && m.board[a][0] == m.board[a][b] {
			return true
		}
	}

	if b == 1 || b == 3 || b == 5 || b == 7 {
		if m.board[0][b] == m.board[1][b] && m.board[2][b] == m.board[a][b] {
			return true
		}
	}

	return false
}
