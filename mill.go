package mill

import (
	"errors"
)

type Player int

const (
	PlayerNone Player = iota
	PlayerOne
	PlayerTwo
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
	board         [3][8]Player
	currentPlayer Player

	menLeftToPlacePlayerOne int
	menLeftToPlacePlayerTwo int
}

func NewMill() *Mill {
	return &Mill{
		currentPlayer:           PlayerOne,
		menLeftToPlacePlayerOne: 9,
		menLeftToPlacePlayerTwo: 9,
	}
}

func (m *Mill) PlaceMan(a, b int) *error {
	if !m.hasMenLeft() {
		return &ErrNoMenLeft
	}

	if m.placeCellUnrestricted(a, b) {
		m.decreaseMenLeft()
		m.nextPlayer()
		return &Success
	}

	return &ErrBadInput
}

func (m *Mill) MoveMan(a, b, c, d int) *error {
	if m.moveCellToNeighbor(a, b, c, d) {
		m.nextPlayer()
		return &Success
	} else {
		return &ErrBadInput
	}
}

func (m *Mill) TakeManFromOpponent(a, b int) *error {
	if m.isMill(a, b) {
		return &ErrBadInput
	}

	if m.board[a][b] == m.currentPlayer {
		return &ErrBadInput
	}

	return &Success
}

func (m *Mill) placeCellUnrestricted(a, b int) bool {
	if m.board[a][b] == PlayerNone {
		m.board[a][b] = m.currentPlayer
		return true
	}
	return false
}

func (m *Mill) moveCellToNeighbor(a, b, c, d int) bool {
	if m.board[a][b] != m.currentPlayer {
		return false
	}

	if m.board[c][d] != PlayerNone {
		return false
	}

	if a == c && (b-d == 1 || d-b == 1) {
		goto ok
	} else if (b == 1 || b == 3 || b == 5 || b == 7) && b == d && (a-b == 1 || b-a == 1) {
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

func (m *Mill) hasMenLeft() bool {
	if m.currentPlayer == PlayerOne && m.menLeftToPlacePlayerOne > 0 {
		return true
	}
	if m.currentPlayer == PlayerTwo && m.menLeftToPlacePlayerTwo > 0 {
		return true
	}

	return false
}

func (m *Mill) decreaseMenLeft() {
	if m.currentPlayer == PlayerOne {
		m.menLeftToPlacePlayerOne--
	} else {
		m.menLeftToPlacePlayerTwo--
	}
}

func (m *Mill) isMill(a, b int) bool {
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

func (m *Mill) nextPlayer() {
	if m.currentPlayer == PlayerOne {
		m.currentPlayer = PlayerTwo
	} else {
		m.currentPlayer = PlayerOne
	}
}
