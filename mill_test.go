package mill

import (
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

func TestEnemyHasVulnerableMan(t *testing.T) {
	mill := NewMill()

	mill.PlaceMan(0, 0)
	mill.PlaceMan(1, 0)
	mill.PlaceMan(2, 3)
	mill.PlaceMan(2, 5)

	if !mill.EnemyHasVulnerableMan() {
		t.Fatalf("the enemy does have a vulnerable man")
	}

}
