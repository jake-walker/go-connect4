package game

import (
	"errors"
	"log"
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

func (g Game) IsWinner(piece int) bool {
	for x := 0; x < BoardX-3; x++ {
		for y := 0; y < BoardY; y++ {
			if g.Board[x][y] == piece && g.Board[x+1][y] == piece && g.Board[x+2][y] == piece && g.Board[x+3][y] == piece {
				log.Printf("Found horizontal win at (%v, %v)", x, y)
				return true
			}
		}
	}

	for x := 0; x < BoardX; x++ {
		for y := 0; y < BoardY-3; y++ {
			if g.Board[x][y] == piece && g.Board[x][y+1] == piece && g.Board[x][y+2] == piece && g.Board[x][y+3] == piece {
				log.Printf("Found vertical win at (%v, %v)", x, y)
				return true
			}
		}
	}

	for x := 0; x < BoardX-3; x++ {
		for y := 0; y < BoardY-3; y++ {
			if g.Board[x][y] == piece && g.Board[x+1][y+1] == piece && g.Board[x+2][y+2] == piece && g.Board[x+3][y+3] == piece {
				log.Printf("Found NE/SW diagonal win at (%v, %v)", x, y)
				return true
			}
		}
	}

	for x := 3; x < BoardX; x++ {
		for y := 0; y < BoardY-3; y++ {
			if g.Board[x][y] == piece && g.Board[x-1][y+1] == piece && g.Board[x-2][y+2] == piece && g.Board[x-3][y+3] == piece {
				log.Printf("Found NW/SE diagonal win at (%v, %v)", x, y)
				return true
			}
		}
	}

	return false
}

func (g Game) IsFinished() bool {
	if g.IsWinner(PieceA) || g.IsWinner(PieceB) {
		return true
	}

	for x := 0; x < BoardX; x++ {
		for y := 0; y < BoardY; y++ {
			if g.Board[x][y] == PieceEmpty {
				return false
			}
		}
	}

	return true
}

func (g Game) IsValidMove(col int) bool {
	return g.Board[col][0] == PieceEmpty && !(col < 0 || col > BoardX)
}

func (g *Game) PlacePiece(col int, piece int) error {
	if !g.IsValidMove(col) {
		return errors.New("invalid move")
	}

	for i := BoardY - 1; i >= 0; i-- {
		if g.Board[col][i] == PieceEmpty {
			g.Board[col][i] = piece
			return nil
		}
	}

	return nil
}
