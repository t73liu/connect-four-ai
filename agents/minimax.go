package agents

import (
	"log"
	"math"

	"github.com/t73liu/connectfourai/game"
	"github.com/t73liu/connectfourai/utils"
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
		if utils.GreaterThanFloat64(moveVal, bestVal) {
			bestMove = move
			bestVal = moveVal
		}
	}
	return bestMove
}

func (mma *MiniMaxAgent) minimax(g *game.Game, isMaximizer bool, depth int32) float64 {
	if depth == mma.MaxDepth || g.IsGameOver() {
		return game.Evaluate(g, depth, mma.MaximizerPiece)
	}
	if isMaximizer {
		bestVal := utils.NegativeInfinity
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
		bestVal := utils.PositiveInfinity
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
