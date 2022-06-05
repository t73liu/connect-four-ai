package agents

import (
	"fmt"
	"log"

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
	bestMove, moveVal := mma.minimax(g, true, 0)
	fmt.Printf("MiniMax chose move %d (value: %.2f)\n", bestMove, moveVal)
	return bestMove
}

func (mma *MiniMaxAgent) minimax(g *game.Game, isMaximizer bool, depth int32) (int32, float64) {
	if depth == mma.MaxDepth || g.IsGameOver() {
		return g.GetPreviousMove().RowIndex, game.Evaluate(g, depth, mma.MaximizerPiece)
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
			_, moveVal := mma.minimax(g, false, depth+1)
			if utils.GreaterThanFloat64(moveVal, bestVal) {
				bestVal = moveVal
				bestMove = move
			}
			g.UndoMove()
		}
	} else {
		bestVal = utils.PositiveInfinity
		for _, move := range g.ListValidMoves() {
			err := g.MakeMove(move)
			if err != nil {
				log.Fatalln("Unexpected move error", err)
			}
			_, moveVal := mma.minimax(g, true, depth+1)
			if utils.LessThanFloat64(moveVal, bestVal) {
				bestVal = moveVal
				bestMove = move
			}
			g.UndoMove()
		}
	}
	return bestMove, bestVal
}
