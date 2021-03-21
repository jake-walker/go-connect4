package game

import "testing"

func TestImmediateWins2(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{p_, p_, p_, p_, pa, pa},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
		},
	}

	actual := immediateWins(myGame, pa)
	expect := -1

	if actual != expect {
		t.Fatalf("immediateWins is %v, want %v.", actual, expect)
	}
}

func TestImmediateWins3(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{p_, p_, p_, pa, pa, pa},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
		},
	}

	actual := immediateWins(myGame, pa)
	expect := 0

	if actual != expect {
		t.Fatalf("immediateWins is %v, want %v.", actual, expect)
	}
}

func TestImmediateWins4(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{p_, p_, pa, pa, pa, pa},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
		},
	}

	actual := immediateWins(myGame, pa)
	expect := 0

	if actual != expect {
		t.Fatalf("immediateWins is %v, want %v.", actual, expect)
	}
}
