package mill

type Player struct {
	menLeftToPlace int
	menOnBoard     int
}

func NewPlayerNone() Player {
	return Player{
		menLeftToPlace: 0,
		menOnBoard:     0,
	}
}

func NewPlayer() Player {
	return Player{
		menLeftToPlace: 9,
		menOnBoard:     0,
	}
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

func (m *Mill) increaseMenOnBoard() {
	if m.currentPlayer == PlayerOne {
		m.menOnBoardPlayerOne++
	} else {
		m.menOnBoardPlayerTwo++
	}
}

func (m *Mill) decreaseOpponentsMenOnBoard() {
	if m.currentPlayer == PlayerOne {
		m.menOnBoardPlayerTwo--
	} else {
		m.menOnBoardPlayerOne--
	}
}
