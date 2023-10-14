package mill

import "errors"

type Player int

const (
	PlayerNone Player = iota
	PlayerOne
	PlayerTwo
)

var (
	Success     error = nil
	ErrBadInput error = errors.New("input is invalid")
)

type Mill struct {
	board         [3][8]Player
	currentPlayer Player
}

func NewMill() *Mill {
	return &Mill{currentPlayer: PlayerOne}
}

func (m *Mill) PlaceMan(a, b int) *error {
	err := m.placeCellUnrestricted(m.currentPlayer, a, b)
	if err == &Success {
		m.nextPlayer()
	}

	return err
}

func (m *Mill) MoveMan(a, b, c, d int) *error {
	err := m.moveCellToNeighbor(a, b, c, d)
	if err == &Success {
		m.nextPlayer()
	}

	return err
}

func (m *Mill) placeCellUnrestricted(p Player, a, b int) *error {
	if m.board[a][b] == PlayerNone {
		m.board[a][b] = p
		return m.isMill(a, b)
	}
	return &ErrBadInput
}

func (m *Mill) moveCellToNeighbor(a, b, c, d int) *error {
	if m.board[a][b] != m.currentPlayer {
		return nil
	}

	if m.board[c][d] != PlayerNone {
		return nil
	}

	m.board[c][d] = m.board[a][b]
	m.board[a][b] = PlayerNone

	return &Success
}

func (m *Mill) isMill(a, b int) *error {
	return &Success
}

func (m *Mill) nextPlayer() {
	if m.currentPlayer == PlayerOne {
		m.currentPlayer = PlayerTwo
	} else {
		m.currentPlayer = PlayerOne
	}
}
