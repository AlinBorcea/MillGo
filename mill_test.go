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
