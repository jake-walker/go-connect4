package game

import (
	"log"
	"math/rand"
	"time"
)

// The number of simulations for the computer to run for each valid move
const simulationCount = 7500

// Aliases for pieces
const ComputerPiece = PieceB
const HumanPiece = PieceA

func init() {
	log.SetPrefix("computer: ")
	rand.Seed(time.Now().UnixNano())
}

// Get a random move (i.e. a random column on the board)
func randomMove() int {
	return rand.Intn(BoardX)
}

// Fill a given board with random moves until somebody wins or there is a tie
func runSimulation(g *Game) {
	for !g.IsFinished() {
		col := randomMove()

		if !g.IsValidMove(col) {
			continue
		}

		g.PlacePiece(col, g.CurrentTurn)
		g.SwapTurn()
	}
}

// Monte Carlo algorithm for choosing a best move
func monteCarlo(g Game) int {
	// The ratio of the best move
	bestRatio := 0.0
	// The column of the best move
	bestMove := -1

	for move := 0; move < BoardX; move++ {
		if !g.IsValidMove(move) {
			continue
		}

		wins := 0
		losses := 0

		for sim := 0; sim < simulationCount; sim++ {
			// Create a copy of the board
			gameSim := g

			// Place a piece in the current move
			gameSim.PlacePiece(move, ComputerPiece)
			// Start the simulation with the human going first
			gameSim.CurrentTurn = HumanPiece
			runSimulation(&gameSim)

			// If the computer has won, increment the win counter
			if gameSim.IsWinner(ComputerPiece) {
				wins++
			} else {
				losses++
			}
		}

		var ratio float64 = 0

		// Calculate the win to loss ratio
		if losses > 0 {
			ratio = float64(wins) / float64(losses)
		}

		log.Printf("Move %v has success %v", move, ratio)

		// If this is the best ratio so far, set the current best ratio and best move
		if ratio > bestRatio {
			bestRatio = ratio
			bestMove = move
		}
	}

	log.Printf("Monte Carlo output is move %v", bestMove)

	return bestMove
}

// Check each column for a move that would win
func immediateWins(g Game, piece int) int {
	for move := 0; move < BoardX; move++ {
		if !g.IsValidMove(move) {
			continue
		}

		// Create a copy of the board
		boardCopy := g
		// Place the given piece in that column
		boardCopy.PlacePiece(move, piece)

		// Does placing that piece result in a win?
		if boardCopy.IsWinner(piece) {
			return move
		}
	}

	return -1
}

func DoComputerMove(g Game) int {
	move := -1

	// Find wins that would make the computer win this turn
	log.Println("Finding winning moves...")
	move = immediateWins(g, ComputerPiece)

	if move == -1 {
		// Otherwise, find wins that would make the human win next turn
		log.Println("Finding blocking moves...")
		move = immediateWins(g, HumanPiece)
	}

	if move == -1 {
		// Otherwise, if nobody can win, use Monte Carlo to find the best move
		log.Println("Using Monte Carlo...")
		move = monteCarlo(g)
	}

	if move == -1 {
		// If everything else fails (which is very unlikely), choose a random place
		log.Println("Falling back to random move...")
		move = randomMove()
	}

	log.Printf("Calculated best move is %v", move)
	return move
}
