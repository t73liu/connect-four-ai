package agents

import (
	"log"
	"math"

	"github.com/t73liu/connectfourai/game"
)

type MiniMaxAgent struct {
	MaximizerPiece game.Piece
	MaxDepth       int32
}

func (mma *MiniMaxAgent) GetMove(g *game.Game) int32 {
	validMoves := g.ListValidMoves()
	if len(validMoves) == 0 {
		log.Fatalln("no valid moves left")
	}
	var bestMove int32
	bestVal := math.Inf(-1)
	for _, move := range validMoves {
		err := g.MakeMove(move)
		if err != nil {
			log.Fatalln("Unexpected move error", err)
		}
		moveVal := mma.minimax(g, false, 1)
		g.UndoMove()
		if moveVal > bestVal {
			bestMove = move
			bestVal = moveVal
		}
	}
	return bestMove
}

func (mma *MiniMaxAgent) minimax(g *game.Game, isMaximizer bool, depth int32) float64 {
	if depth == mma.MaxDepth || g.IsGameOver() {
		return mma.evaluate(g, depth)
	}
	if isMaximizer {
		bestVal := math.Inf(-1)
		for _, move := range g.ListValidMoves() {
			err := g.MakeMove(move)
			if err != nil {
				log.Fatalln("Unexpected move error", err)
			}
			bestVal = math.Max(bestVal, mma.minimax(g, false, depth+1))
			g.UndoMove()
		}
		return bestVal
	} else {
		bestVal := math.Inf(1)
		for _, move := range g.ListValidMoves() {
			err := g.MakeMove(move)
			if err != nil {
				log.Fatalln("Unexpected move error", err)
			}
			bestVal = math.Min(bestVal, mma.minimax(g, true, depth+1))
			g.UndoMove()
		}
		return bestVal
	}
}

var scoreDirections = [][]int{
	// Right
	{0, 1},
	// Down
	{1, 0},
	// Down-Right
	{1, 1},
	// Down-Left
	{1, -1},
}

var potentialDirections = [][]int{
	// Right
	{0, 1},
	// Left
	{0, -1},
	// Up
	{-1, 0},
	// Up-Right
	{-1, 1},
	// Up-Left
	{-1, -1},
	// Down-Right
	{1, 1},
	// Down-Left
	{1, -1},
}

func (mma *MiniMaxAgent) evaluate(g *game.Game, depth int32) float64 {
	if g.IsGameOver() && g.State != game.Draw {
		previousMove := g.GetPreviousMove()
		if previousMove == nil {
			log.Fatalln("Unable to get previous move in game over state")
		}
		if previousMove.Piece == mma.MaximizerPiece {
			return float64(1000 - depth)
		} else {
			return float64(-1000 + depth)
		}
	}
	var score float64
	for rowIndex, row := range g.Board {
		for colIndex, cell := range row {
			if cell != mma.MaximizerPiece {
				continue
			}
			// Favour moves that connect rows, columns or diagonals.
			for _, direction := range scoreDirections {
				adjacentRowIndex := rowIndex + direction[0]
				adjacentColIndex := colIndex + direction[1]
				if game.IsWithinBounds(adjacentRowIndex, adjacentColIndex) {
					adjacentPiece := g.Board[adjacentRowIndex][adjacentColIndex]
					if adjacentPiece == mma.MaximizerPiece {
						score++
					}
				}
			}
			// Favour moves that have potential for connect four with empty
			// spaces.
			for _, direction := range potentialDirections {
				adjacentRowIndex := rowIndex + direction[0]
				adjacentColIndex := colIndex + direction[1]
				for game.IsWithinBounds(adjacentRowIndex, adjacentColIndex) {
					adjacentPiece := g.Board[adjacentRowIndex][adjacentColIndex]
					if adjacentPiece == mma.MaximizerPiece {
						score += 0.2
					} else if adjacentPiece == game.Empty {
						score += 0.1
					} else {
						break
					}
					adjacentRowIndex += direction[0]
					adjacentColIndex += direction[1]
				}
			}
		}
	}
	return score
}
