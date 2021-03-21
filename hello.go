package main

import (
	"fmt"

	ct "github.com/daviddengcn/go-colortext"
	"jakew.me/connect4/game"
)

func main() {
	myGame := game.Game{
		CurrentTurn: game.PieceA,
	}

	for !myGame.IsFinished() {
		printBoard(myGame)
		askPlace(&myGame)
	}
}

func askPlace(g *game.Game) {
	switch g.CurrentTurn {
	case game.PieceA:
		fmt.Println("It is player A's turn!")
	case game.PieceB:
		fmt.Println("It is player B's turn!")
	}

	var i = 0

	for i < 1 || i > game.BoardX {
		fmt.Print("Where would you like to go? ")
		fmt.Scan(&i)
	}

	err := g.PlacePiece(i-1, g.CurrentTurn)
	if err != nil {
		fmt.Println("Whoops! That was an invalid move")
		askPlace(g)
		return
	}

	if g.CurrentTurn == game.PieceA {
		g.CurrentTurn = game.PieceB
	} else {
		g.CurrentTurn = game.PieceA
	}
}

func printBoard(g game.Game) {
	for col := 1; col < game.BoardX+1; col++ {
		fmt.Printf(" %v", col)
	}
	fmt.Print("\n")

	for y := 0; y < game.BoardY; y++ {
		for x := 0; x < game.BoardX; x++ {
			switch g.Board[x][y] {
			case game.PieceEmpty:
				fmt.Print(" _")
			case game.PieceA:
				ct.Foreground(ct.Red, true)
				fmt.Print(" X")
				ct.ResetColor()
			case game.PieceB:
				ct.Foreground(ct.Green, true)
				fmt.Print(" O")
				ct.ResetColor()
			}
		}
		fmt.Println("")
	}
}
