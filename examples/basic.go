package main

import (
	"fmt"
	"os"

	"github.com/AlinBorcea/mill"
)

func main() {
	game := mill.NewMill()
	var err *error

	fmt.Println("Game started")
	for game.Status() != mill.StatusGameDone {
		switch game.Status() {
		case mill.StatusTurnDone:
			fmt.Println("Going to next player")
			game.NextPlayer()
		case mill.StatusAwaitPlaceMan:
			game.PrintTable(os.Stdout)
			fmt.Println("Place man")
			err = placeMan(game)
			handlePlaceManError(err)
		case mill.StatusAwaitMoveMan:
			game.PrintTable(os.Stdout)
			fmt.Println("Move man")
			err = moveMan(game)
			handleMoveManError(err)
		case mill.StatusAwaitTargetMan:
			fmt.Println("Pick enemy")
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
