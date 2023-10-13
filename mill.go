package mill

type Player int

const (
	PlayerNone Player = iota
	PlayerOne
	PlayerTwo
)

type Mill struct {
	board [3][8]int
}

func NewMill() *Mill {
	return &Mill{}
}

func (m *Mill) PlaceMan(p Player, a, b int) error {
	return nil
}
