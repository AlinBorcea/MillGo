package main

import (
	"fmt"

	"github.com/AlinBorcea/mill"
)

func main() {
	game := mill.NewMill()
	var err *error

	for game.Status() != mill.StatusGameDone {
		switch game.Status() {
		case mill.StatusTurnDone:
			game.NextPlayer()
		case mill.StatusAwaitPlaceMan:
			err = placeMan(game)
			handlePlaceManError(err)
		case mill.StatusAwaitMoveMan:
			err = moveMan(game)
			handleMoveManError(err)
		case mill.StatusAwaitTargetMan:
			err = takeMan(game)
			handleTakeManError(err)
		}
	}
}

func placeMan(game *mill.Mill) *error {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)
	return game.PlaceMan(a, b)
}

func handlePlaceManError(err *error) {
	fmt.Println(err)
}

func moveMan(game *mill.Mill) *error {
	var a, b, c, d int

	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	return game.MoveMan(a, b, c, d)
}

func handleMoveManError(err *error) {
	fmt.Println(err)
}

func takeMan(game *mill.Mill) *error {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)
	return game.TakeManFromOpponent(a, b)
}

func handleTakeManError(err *error) {
	fmt.Println(err)
}
