package game

import (
	"errors"
)

const BoardX = 7
const BoardY = 6

const (
	PieceEmpty = iota
	PieceA     = iota
	PieceB     = iota
)

type Game struct {
	Board       [BoardX][BoardY]int
	CurrentTurn int
}

// Get the coordinates of the given piece's winning move
func (g Game) WinningMove(piece int) (x1, y1, x2, y2, x3, y3, x4, y4 int) {
	// Get the coordinates for a move that wins horizontally
	for x := 0; x < BoardX-3; x++ {
		for y := 0; y < BoardY; y++ {
			if g.Board[x][y] == piece && g.Board[x+1][y] == piece && g.Board[x+2][y] == piece && g.Board[x+3][y] == piece {
				// log.Printf("Found horizontal win at (%v, %v)", x, y)
				return x, y, x + 1, y, x + 2, y, x + 3, y
			}
		}
	}

	// Get the coordinates for a move that wins vertically
	for x := 0; x < BoardX; x++ {
		for y := 0; y < BoardY-3; y++ {
			if g.Board[x][y] == piece && g.Board[x][y+1] == piece && g.Board[x][y+2] == piece && g.Board[x][y+3] == piece {
				// log.Printf("Found vertical win at (%v, %v)", x, y)
				return x, y, x, y + 1, x, y + 2, x, y + 3
			}
		}
	}

	// Get the coordinates for a move that wins diagonally (NE/SW)
	for x := 0; x < BoardX-3; x++ {
		for y := 0; y < BoardY-3; y++ {
			if g.Board[x][y] == piece && g.Board[x+1][y+1] == piece && g.Board[x+2][y+2] == piece && g.Board[x+3][y+3] == piece {
				// log.Printf("Found NE/SW diagonal win at (%v, %v)", x, y)
				return x, y, x + 1, y + 1, x + 2, y + 2, x + 3, y + 3
			}
		}
	}

	// Get the coordinates for a move that wins diagonally (NW/SE)
	for x := 3; x < BoardX; x++ {
		for y := 0; y < BoardY-3; y++ {
			if g.Board[x][y] == piece && g.Board[x-1][y+1] == piece && g.Board[x-2][y+2] == piece && g.Board[x-3][y+3] == piece {
				// log.Printf("Found NW/SE diagonal win at (%v, %v)", x, y)
				return x, y, x - 1, y + 1, x - 2, y + 2, x - 3, y + 3
			}
		}
	}

	// If no other moves were found
	return -1, -1, -1, -1, -1, -1, -1, -1
}

// Check if the given piece has won
func (g Game) IsWinner(piece int) bool {
	x1, y1, x2, y2, x3, y3, x4, y4 := g.WinningMove(piece)
	// The given piece has won if all the coordinates are -1
	return x1 != -1 && y1 != -1 && x2 != -1 && y2 != -1 && x3 != -1 && y3 != -1 && x4 != -1 && y4 != -1
}

// Check if the board is full
func (g Game) IsFull() bool {
	for x := 0; x < BoardX; x++ {
		for y := 0; y < BoardY; y++ {
			if g.Board[x][y] == PieceEmpty {
				return false
			}
		}
	}

	return true
}

// Check if the game is finished
func (g Game) IsFinished() bool {
	// If somebody has won or the board is full, the game is finished
	return g.IsWinner(PieceA) || g.IsWinner(PieceB) || g.IsFull()
}

// Check if a given move is valid
func (g Game) IsValidMove(col int) bool {
	// Ensure the column is within range and the column is not full
	return !(col < 0 || col > BoardX) && g.Board[col][0] == PieceEmpty
}

// Place a piece in a given column
func (g *Game) PlacePiece(col int, piece int) error {
	// Ensure the move is valid
	if !g.IsValidMove(col) {
		return errors.New("invalid move")
	}

	// Start from the bottom of the column, and fill the first available empty square
	for i := BoardY - 1; i >= 0; i-- {
		if g.Board[col][i] == PieceEmpty {
			g.Board[col][i] = piece
			return nil
		}
	}

	return nil
}

// End the current turn
func (g *Game) SwapTurn() {
	// Swap the turn, so if it is player A's turn, switch to player B and vice-versa
	if g.CurrentTurn == PieceA {
		g.CurrentTurn = PieceB
	} else {
		g.CurrentTurn = PieceA
	}
}
