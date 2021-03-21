package main

import (
	"fmt"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/jake-walker/go-connect4/game"
)

func main() {
	myGame := game.Game{
		CurrentTurn: game.PieceA,
	}

	printBoard(myGame)

	for !myGame.IsFinished() {
		if myGame.CurrentTurn == game.HumanPiece {
			fmt.Println("It is your turn!")
			askPlace(&myGame)
		} else {
			fmt.Println("It is computer's turn!")
			move := game.DoComputerMove(myGame)
			myGame.PlacePiece(move, game.ComputerPiece)
			myGame.CurrentTurn = game.HumanPiece
		}

		printBoard(myGame)
	}

	if myGame.IsWinner(game.HumanPiece) {
		fmt.Println("Congrats! You beat the computer")
	} else {
		fmt.Println("Unlucky, the computer beat you this time")
	}
}

func askPlace(g *game.Game) {
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

	g.SwapTurn()
}

func printBoard(g game.Game) {
	var winX1, winY1, winX2, winY2, winX3, winY3, winX4, winY4 int = -1, -1, -1, -1, -1, -1, -1, -1
	if g.IsWinner(game.PieceA) {
		winX1, winY1, winX2, winY2, winX3, winY3, winX4, winY4 = g.WinningMove(game.PieceA)
	} else if g.IsWinner(game.PieceB) {
		winX1, winY1, winX2, winY2, winX3, winY3, winX4, winY4 = g.WinningMove(game.PieceB)
	}

	fmt.Print("\n")
	for col := 1; col < game.BoardX+1; col++ {
		fmt.Printf(" %v", col)
	}
	fmt.Print("\n")

	for y := 0; y < game.BoardY; y++ {
		for x := 0; x < game.BoardX; x++ {
			fmt.Print(" ")

			if (x == winX1 && y == winY1) || (x == winX2 && y == winY2) || (x == winX3 && y == winY3) || (x == winX4 && y == winY4) {
				ct.Background(ct.Yellow, false)
			}

			switch g.Board[x][y] {
			case game.PieceEmpty:
				fmt.Print("_")
			case game.PieceA:
				ct.Foreground(ct.Red, true)
				fmt.Print("X")
				ct.ResetColor()
			case game.PieceB:
				ct.Foreground(ct.Green, true)
				fmt.Print("O")
				ct.ResetColor()
			}
		}
		fmt.Println("")
	}
}
