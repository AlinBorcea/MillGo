package mill

type Player struct {
	menLeftToPlace int
	menOnBoard     int
}

func NewPlayerNone() *Player {
	return &Player{
		menLeftToPlace: 0,
		menOnBoard:     0,
	}
}

func NewPlayer() *Player {
	return &Player{
		menLeftToPlace: 9,
		menOnBoard:     0,
	}
}

func (p *Player) hasMenLeft() bool {
	return p.menLeftToPlace > 0
}

func (p *Player) decreaseMenLeft() {
	p.menLeftToPlace--
}

func (p *Player) decreaseMenOnBoard() {
	p.menOnBoard--
}

func (p *Player) increaseMenOnBoard() {
	p.menOnBoard++
}
