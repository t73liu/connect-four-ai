package agents

import (
	"log"
	"math"

	"github.com/t73liu/connectfourai/game"
	"github.com/t73liu/connectfourai/utils"
)

type AlphaBetaAgent struct {
	MaximizerPiece game.Piece
	MaxDepth       int32
}

func (aba *AlphaBetaAgent) GetMove(g *game.Game) int32 {
	validMoves := g.ListValidMoves()
	if len(validMoves) == 0 {
		log.Fatalln("no valid moves left")
	}
	var bestMove int32
	bestVal := utils.NegativeInfinity
	alpha := utils.NegativeInfinity
	beta := utils.PositiveInfinity
	for _, move := range validMoves {
		err := g.MakeMove(move)
		if err != nil {
			log.Fatalln("Unexpected move error", err)
		}
		moveVal := aba.minimax(g, false, 1, alpha, beta)
		g.UndoMove()
		if utils.GreaterThanFloat64(moveVal, bestVal) {
			bestMove = move
			bestVal = moveVal
			alpha = moveVal
		}
	}
	return bestMove
}

func (aba *AlphaBetaAgent) minimax(
	g *game.Game,
	isMaximizer bool,
	depth int32,
	alpha float64,
	beta float64,
) float64 {
	if depth == aba.MaxDepth || g.IsGameOver() {
		return game.Evaluate(g, depth, aba.MaximizerPiece)
	}
	if isMaximizer {
		bestVal := utils.NegativeInfinity
		for _, move := range g.ListValidMoves() {
			err := g.MakeMove(move)
			if err != nil {
				log.Fatalln("Unexpected move error", err)
			}
			bestVal = math.Max(bestVal, aba.minimax(g, false, depth+1, alpha, beta))
			alpha = math.Max(alpha, bestVal)
			g.UndoMove()
			if utils.LessThanOrEqualFloat64(beta, alpha) {
				break
			}
		}
		return bestVal
	} else {
		bestVal := utils.PositiveInfinity
		for _, move := range g.ListValidMoves() {
			err := g.MakeMove(move)
			if err != nil {
				log.Fatalln("Unexpected move error", err)
			}
			bestVal = math.Min(bestVal, aba.minimax(g, true, depth+1, alpha, beta))
			beta = math.Min(beta, bestVal)
			g.UndoMove()
			if utils.LessThanOrEqualFloat64(beta, alpha) {
				break
			}
		}
		return bestVal
	}
}
