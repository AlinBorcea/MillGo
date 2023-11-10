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
