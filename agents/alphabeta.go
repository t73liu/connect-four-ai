package agents

import (
	"fmt"
	"log"

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
	bestMove, moveVal := aba.minimax(
		g,
		true,
		0,
		utils.NegativeInfinity,
		utils.PositiveInfinity,
	)
	fmt.Printf("Alpha-Beta chose move %d (value: %.2f)\n", bestMove, moveVal)
	return bestMove
}

func (aba *AlphaBetaAgent) minimax(
	g *game.Game,
	isMaximizer bool,
	depth int32,
	alpha float64,
	beta float64,
) (int32, float64) {
	if depth == aba.MaxDepth || g.IsGameOver() {
		previousMove := g.GetPreviousMove()
		if previousMove == nil {
			log.Fatalln("Unable to get previous move")
		}
		return previousMove.RowIndex, game.Evaluate(g, depth, aba.MaximizerPiece)
	}
	var bestVal float64
	bestMove := int32(-1)
	if isMaximizer {
		bestVal = utils.NegativeInfinity
		for _, move := range g.ListValidMoves() {
			err := g.MakeMove(move)
			if err != nil {
				log.Fatalln("Unexpected move error", err)
			}
			_, moveVal := aba.minimax(g, false, depth+1, alpha, beta)
			if utils.GreaterThanFloat64(moveVal, bestVal) {
				bestVal = moveVal
				alpha = moveVal
				bestMove = move
			}
			g.UndoMove()
			if utils.LessThanOrEqualFloat64(beta, alpha) {
				break
			}
		}
	} else {
		bestVal = utils.PositiveInfinity
		for _, move := range g.ListValidMoves() {
			err := g.MakeMove(move)
			if err != nil {
				log.Fatalln("Unexpected move error", err)
			}
			_, moveVal := aba.minimax(g, true, depth+1, alpha, beta)
			if utils.LessThanFloat64(moveVal, bestVal) {
				bestVal = moveVal
				beta = moveVal
				bestMove = move
			}
			g.UndoMove()
			if utils.LessThanOrEqualFloat64(beta, alpha) {
				break
			}
		}
	}
	return bestMove, bestVal
}
