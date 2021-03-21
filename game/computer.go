package game

import (
	"log"
	"math/rand"
	"time"
)

const simulationCount = 7500
const ComputerPiece = PieceB
const HumanPiece = PieceA

func init() {
	log.SetPrefix("computer: ")
	rand.Seed(time.Now().UnixNano())
}

func randomMove() int {
	return rand.Intn(BoardX)
}

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

func monteCarlo(g Game) int {
	bestRatio := 0.0
	bestMove := -1

	for move := 0; move < BoardX; move++ {
		if !g.IsValidMove(move) {
			continue
		}

		wins := 0
		losses := 0

		for sim := 0; sim < simulationCount; sim++ {
			gameSim := g

			gameSim.PlacePiece(move, ComputerPiece)
			gameSim.CurrentTurn = HumanPiece
			runSimulation(&gameSim)

			if gameSim.IsWinner(ComputerPiece) {
				wins++
			} else {
				losses++
			}
		}

		var ratio float64 = 0

		if losses > 0 {
			ratio = float64(wins) / float64(losses)
		}

		log.Printf("Move %v has success %v", move, ratio)

		if ratio > bestRatio {
			bestRatio = ratio
			bestMove = move
		}
	}

	log.Printf("Monte Carlo output is move %v", bestMove)

	return bestMove
}

func immediateWins(g Game, piece int) int {
	for move := 0; move < BoardX; move++ {
		if !g.IsValidMove(move) {
			continue
		}

		boardCopy := g
		boardCopy.PlacePiece(move, piece)

		if boardCopy.IsWinner(piece) {
			return move
		}
	}

	return -1
}

func DoComputerMove(g Game) int {
	move := -1

	log.Println("Finding winning moves...")
	move = immediateWins(g, ComputerPiece)

	if move == -1 {
		log.Println("Finding blocking moves...")
		move = immediateWins(g, HumanPiece)
	}

	if move == -1 {
		log.Println("Using Monte Carlo...")
		move = monteCarlo(g)
	}

	if move == -1 {
		log.Println("Falling back to random move...")
		move = randomMove()
	}

	log.Printf("Calculated best move is %v", move)
	return move
}
