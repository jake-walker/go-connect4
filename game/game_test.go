package game

import (
	"testing"
)

const p_ = PieceEmpty
const pa = PieceA
const pb = PieceB

func TestIsWinnerEmpty(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
		},
	}

	actualWinnerA := myGame.IsWinner(pa)
	expectWinnerA := false
	actualWinnerB := myGame.IsWinner(pb)
	expectWinnerB := false

	if actualWinnerA != expectWinnerA || actualWinnerB != expectWinnerB {
		t.Fatalf("Winner A is %v, want %v. Winner B is %v, want %v.", actualWinnerA, expectWinnerA, actualWinnerB, expectWinnerB)
	}
}

func TestIsWinnerHorizontal(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, pa},
			{p_, p_, p_, p_, p_, pa},
			{p_, p_, p_, p_, p_, pa},
			{p_, p_, p_, p_, p_, pa},
		},
	}

	actualWinnerA := myGame.IsWinner(pa)
	expectWinnerA := true
	actualWinnerB := myGame.IsWinner(pb)
	expectWinnerB := false

	if actualWinnerA != expectWinnerA || actualWinnerB != expectWinnerB {
		t.Fatalf("Winner A is %v, want %v. Winner B is %v, want %v.", actualWinnerA, expectWinnerA, actualWinnerB, expectWinnerB)
	}
}

func TestIsWinnerVertical(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, pa, pa, pa, pa},
		},
	}

	actualWinnerA := myGame.IsWinner(pa)
	expectWinnerA := true
	actualWinnerB := myGame.IsWinner(pb)
	expectWinnerB := false

	if actualWinnerA != expectWinnerA || actualWinnerB != expectWinnerB {
		t.Fatalf("Winner A is %v, want %v. Winner B is %v, want %v.", actualWinnerA, expectWinnerA, actualWinnerB, expectWinnerB)
	}
}

func TestIsWinnerDiagonalNeSw(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, pa, p_, p_, p_},
			{p_, p_, p_, pa, p_, p_},
			{p_, p_, p_, p_, pa, p_},
			{p_, p_, p_, p_, p_, pa},
		},
	}

	actualWinnerA := myGame.IsWinner(pa)
	expectWinnerA := true
	actualWinnerB := myGame.IsWinner(pb)
	expectWinnerB := false

	if actualWinnerA != expectWinnerA || actualWinnerB != expectWinnerB {
		t.Fatalf("Winner A is %v, want %v. Winner B is %v, want %v.", actualWinnerA, expectWinnerA, actualWinnerB, expectWinnerB)
	}
}

func TestIsWinnerDiagonalNwSe(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, pa},
			{p_, p_, p_, p_, pa, p_},
			{p_, p_, p_, pa, p_, p_},
			{p_, p_, pa, p_, p_, p_},
		},
	}

	actualWinnerA := myGame.IsWinner(pa)
	expectWinnerA := true
	actualWinnerB := myGame.IsWinner(pb)
	expectWinnerB := false

	if actualWinnerA != expectWinnerA || actualWinnerB != expectWinnerB {
		t.Fatalf("Winner A is %v, want %v. Winner B is %v, want %v.", actualWinnerA, expectWinnerA, actualWinnerB, expectWinnerB)
	}
}

func TestIsWinnerPartial(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{p_, p_, p_, p_, pb, pa},
			{p_, p_, p_, pb, pa, pa},
			{p_, p_, p_, pa, pa, pa},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, pa, pa, pa, pb},
			{p_, p_, p_, pb, pb, pb},
			{p_, p_, p_, pa, pa, pb},
		},
	}

	actualWinnerA := myGame.IsWinner(pa)
	expectWinnerA := false
	actualWinnerB := myGame.IsWinner(pb)
	expectWinnerB := false

	if actualWinnerA != expectWinnerA || actualWinnerB != expectWinnerB {
		t.Fatalf("Winner A is %v, want %v. Winner B is %v, want %v.", actualWinnerA, expectWinnerA, actualWinnerB, expectWinnerB)
	}
}

func TestIsFullEmpty(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
		},
	}

	actual := myGame.IsFull()
	expect := false

	if actual != expect {
		t.Fatalf("IsFull is %v, want %v.", actual, expect)
	}
}

func TestIsFullFull(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{pb, pa, pb, pa, pb, pa},
			{pa, pb, pa, pb, pa, pb},
			{pa, pb, pa, pb, pa, pb},
			{pb, pa, pb, pa, pb, pa},
			{pa, pb, pa, pb, pa, pb},
			{pa, pb, pa, pb, pa, pb},
			{pb, pa, pb, pa, pb, pa},
		},
	}

	actual := myGame.IsFull()
	expect := true

	if actual != expect {
		t.Fatalf("IsFull is %v, want %v.", actual, expect)
	}
}

func TestPlacePieceEmpty(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
		},
	}

	if !myGame.IsValidMove(0) {
		t.Fatalf("Expected a valid move in position 0")
	}

	err := myGame.PlacePiece(0, pa)

	if err != nil {
		t.Fatalf("Expected no error")
	}

	if myGame.Board[0][0] != p_ {
		t.Fatalf("Expected 0,0 to be empty")
	}
	if myGame.Board[0][1] != p_ {
		t.Fatalf("Expected 0,1 to be empty")
	}
	if myGame.Board[0][2] != p_ {
		t.Fatalf("Expected 0,2 to be empty")
	}
	if myGame.Board[0][3] != p_ {
		t.Fatalf("Expected 0,3 to be empty")
	}
	if myGame.Board[0][4] != p_ {
		t.Fatalf("Expected 0,4 to be empty")
	}
	if myGame.Board[0][5] != pa {
		t.Fatalf("Expected 0,5 to be A")
	}
}

func TestPlacePiecePartial(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{p_, p_, p_, pb, pb, pb},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
		},
	}

	if !myGame.IsValidMove(0) {
		t.Fatalf("Expected a valid move in position 0")
	}

	err := myGame.PlacePiece(0, pa)

	if err != nil {
		t.Fatalf("Expected no error")
	}

	if myGame.Board[0][0] != p_ {
		t.Fatalf("Expected 0,0 to be empty")
	}
	if myGame.Board[0][1] != p_ {
		t.Fatalf("Expected 0,1 to be empty")
	}
	if myGame.Board[0][2] != pa {
		t.Fatalf("Expected 0,2 to be A")
	}
	if myGame.Board[0][3] != pb {
		t.Fatalf("Expected 0,3 to be B")
	}
	if myGame.Board[0][4] != pb {
		t.Fatalf("Expected 0,4 to be B")
	}
	if myGame.Board[0][5] != pb {
		t.Fatalf("Expected 0,5 to be B")
	}
}

func TestPlacePieceFull(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
		Board: [7][6]int{
			{pb, pb, pb, pb, pb, pb},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
			{p_, p_, p_, p_, p_, p_},
		},
	}

	if myGame.IsValidMove(0) {
		t.Fatalf("Expected an invalid move in position 0")
	}

	err := myGame.PlacePiece(0, pa)

	if err == nil {
		t.Fatalf("Expected an error")
	}

	if myGame.Board[0][0] != pb {
		t.Fatalf("Expected 0,0 to be B")
	}
	if myGame.Board[0][1] != pb {
		t.Fatalf("Expected 0,1 to be B")
	}
	if myGame.Board[0][2] != pb {
		t.Fatalf("Expected 0,2 to be B")
	}
	if myGame.Board[0][3] != pb {
		t.Fatalf("Expected 0,3 to be B")
	}
	if myGame.Board[0][4] != pb {
		t.Fatalf("Expected 0,4 to be B")
	}
	if myGame.Board[0][5] != pb {
		t.Fatalf("Expected 0,5 to be B")
	}
}

func TestSwapTurn(t *testing.T) {
	myGame := Game{
		CurrentTurn: PieceA,
	}

	myGame.SwapTurn()

	if myGame.CurrentTurn != PieceB {
		t.Fatalf("Expected first swap to be B")
	}

	myGame.SwapTurn()

	if myGame.CurrentTurn != PieceA {
		t.Fatalf("Expected second swap to be A")
	}
}
