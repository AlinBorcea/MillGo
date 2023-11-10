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
	if m.currentPlayerId == PlayerOne && m.menLeftToPlacePlayerOne > 0 {
		return true
	}
	if m.currentPlayerId == PlayerTwo && m.menLeftToPlacePlayerTwo > 0 {
		return true
	}

	return false
}

func (m *Mill) decreaseMenLeft() {
	if m.currentPlayerId == PlayerOne {
		m.menLeftToPlacePlayerOne--
	} else {
		m.menLeftToPlacePlayerTwo--
	}
}

func (m *Mill) increaseMenOnBoard() {
	if m.currentPlayerId == PlayerOne {
		m.menOnBoardPlayerOne++
	} else {
		m.menOnBoardPlayerTwo++
	}
}

func (m *Mill) decreaseOpponentsMenOnBoard() {
	if m.currentPlayerId == PlayerOne {
		m.menOnBoardPlayerTwo--
	} else {
		m.menOnBoardPlayerOne--
	}
}
