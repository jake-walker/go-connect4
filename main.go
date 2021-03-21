package main

import (
	"fmt"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/jake-walker/go-connect4/game"
)

func main() {
	// Make a new game instance
	myGame := game.Game{
		CurrentTurn: game.PieceA,
	}

	printBoard(myGame)

	// Loop through alternating turns until the game is won
	for !myGame.IsFinished() {
		if myGame.CurrentTurn == game.HumanPiece {
			// For the human's turn, ask them where to go
			fmt.Println("It is your turn!")
			askPlace(&myGame)
		} else {
			// For the computer's turn run the computer turn function and place their piece
			fmt.Println("It is computer's turn!")
			move := game.DoComputerMove(myGame)
			myGame.PlacePiece(move, game.ComputerPiece)
			// Finish the computer's turn by setting the human to be next
			myGame.CurrentTurn = game.HumanPiece
		}

		printBoard(myGame)
	}

	// Print a finishing message
	if myGame.IsWinner(game.HumanPiece) {
		fmt.Println("Congrats! You beat the computer")
	} else if myGame.IsWinner(game.ComputerPiece) {
		fmt.Println("Unlucky, the computer beat you this time")
	} else {
		fmt.Println("It was a tie")
	}
}

func askPlace(g *game.Game) {
	var i = 0

	// Keep asking where to go until a valid selection is made
	for i < 1 || i > game.BoardX {
		fmt.Print("Where would you like to go? ")
		fmt.Scan(&i)
	}

	// Place the piece
	err := g.PlacePiece(i-1, g.CurrentTurn)
	if err != nil {
		// If that was an invalid turn, ask again
		fmt.Println("Whoops! That was an invalid move")
		askPlace(g)
		return
	}

	// Swap the turn to the other player
	g.SwapTurn()
}

func printBoard(g game.Game) {
	// Set win coordinates to all -1 so that nothing is highlighted when there are no winners
	var winX1, winY1, winX2, winY2, winX3, winY3, winX4, winY4 int = -1, -1, -1, -1, -1, -1, -1, -1

	// Load in the winning coordinates if anyone has won
	if g.IsWinner(game.PieceA) {
		winX1, winY1, winX2, winY2, winX3, winY3, winX4, winY4 = g.WinningMove(game.PieceA)
	} else if g.IsWinner(game.PieceB) {
		winX1, winY1, winX2, winY2, winX3, winY3, winX4, winY4 = g.WinningMove(game.PieceB)
	}

	fmt.Print("\n")
	// Print out the column headings
	for col := 1; col < game.BoardX+1; col++ {
		fmt.Printf(" %v", col)
	}
	fmt.Print("\n")

	for y := 0; y < game.BoardY; y++ {
		for x := 0; x < game.BoardX; x++ {
			fmt.Print(" ")

			// If this square is a winning one, highlight it
			if (x == winX1 && y == winY1) || (x == winX2 && y == winY2) || (x == winX3 && y == winY3) || (x == winX4 && y == winY4) {
				ct.Background(ct.Yellow, false)
			}

			switch g.Board[x][y] {
			case game.PieceEmpty:
				fmt.Print("_")
			case game.PieceA:
				ct.Foreground(ct.Red, true)
				fmt.Print("X")
			case game.PieceB:
				ct.Foreground(ct.Green, true)
				fmt.Print("O")
			}

			ct.ResetColor()
		}
		fmt.Println("")
	}
}
