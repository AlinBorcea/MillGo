package mill

import (
	"fmt"
	"testing"
)

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
