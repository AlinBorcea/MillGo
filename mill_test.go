package mill

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestManually(t *testing.T) {
	mill := NewMill()
	err := &Success
	scanner := bufio.NewScanner(os.Stdin)
	placementStage := true
	moveStage := false

	var a, b, c, d int
	for err != nil {
		// input
		fmt.Print("Input: ")
		if placementStage && readInputAB(scanner, &a, &b) == nil {
			err = mill.PlaceMan(a, b)
		} else if moveStage && readInputABCD(scanner, &a, &b, &c, &d) == nil {
			err = mill.MoveMan(a, b, c, d)
		}

		fmt.Println(err)
	}
}

func readInputAB(scanner *bufio.Scanner, a, b *int) error {
	scanner.Scan()
	rawInput := scanner.Text()
	input := strings.Split(rawInput, " ")
	if len(input) != 2 {
		return fmt.Errorf("bad input")
	}
	*a, _ = strconv.Atoi(input[0])
	*b, _ = strconv.Atoi(input[1])
	return nil
}

func readInputABCD(scanner *bufio.Scanner, a, b, c, d *int) error {
	scanner.Scan()
	rawInput := scanner.Text()
	input := strings.Split(rawInput, " ")
	if len(input) != 4 {
		return fmt.Errorf("bad input")
	}
	*a, _ = strconv.Atoi(input[0])
	*b, _ = strconv.Atoi(input[1])
	*c, _ = strconv.Atoi(input[2])
	*d, _ = strconv.Atoi(input[3])
	return nil
}

func TestPlaceMan(t *testing.T) {
	mill := NewMill()
	err := mill.PlaceMan(0, 0)

	if err != &Success {
		t.Fatalf("the table is empty but cannot place man")
	}

	err = mill.PlaceMan(0, 0)
	if err != &ErrBadInput {
		t.Fatalf("placed man on another man")
	}

	err = mill.PlaceMan(0, 1)
	if err != &Success {
		t.Fatalf("failed to place man in empty cell")
	}

	fmt.Print(mill.board)
}

func TestMoveMan(t *testing.T) {
	mill := NewMill()
	mill.PlaceMan(0, 0)
	mill.PlaceMan(1, 0)

	err := mill.MoveMan(0, 0, 0, 0)
	if err == &Success {
		t.Fatalf("cannot move man to its current position")
	}

	err = mill.MoveMan(0, 0, 0, 1)
	if err != &Success {
		t.Fatalf("failed to move man")
	}

	err = mill.MoveMan(0, 0, 0, 7)
	if err == &Success {
		t.Fatalf("moved empty cell")
	}

	err = mill.MoveMan(1, 0, 0, 0)
	if err == &Success {
		t.Fatalf("illegal move")
	}
}

func TestTakeOpponentDown(t *testing.T) {
	mill := NewMill()
	var err *error = &Success

	mill.PlaceMan(1, 5)
	mill.PlaceMan(1, 6)

	err = mill.TakeManFromOpponent(1, 5)
	if err != &ErrBadInput {
		t.Fatalf("took down own man")
	}

	err = mill.TakeManFromOpponent(0, 3)
	if err != &ErrBadInput {
		t.Fatalf("there is no man to take")
	}

	err = mill.TakeManFromOpponent(1, 6)
	if err != &Success {
		t.Fatalf("could not take man")
	}
}
