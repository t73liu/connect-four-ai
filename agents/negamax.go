package agents

import (
	"fmt"
	"log"

	"github.com/t73liu/connectfourai/game"
	"github.com/t73liu/connectfourai/utils"
)

type NegaMaxAgent struct {
	PlayerPiece game.Piece
	MaxDepth    int32
}

func (nma *NegaMaxAgent) GetMove(g *game.Game) int32 {
	validMoves := g.ListValidMoves()
	if len(validMoves) == 0 {
		log.Fatalln("no valid moves left")
	}
	bestMove, moveVal := nma.negamax(g, 0, 1)
	fmt.Printf("NegaMax chose move %d (value: %.2f)\n", bestMove, moveVal)
	return bestMove
}

func (nma *NegaMaxAgent) negamax(g *game.Game, depth int32, color int32) (int32, float64) {
	if depth == nma.MaxDepth || g.IsGameOver() {
		previousMove := g.GetPreviousMove()
		if previousMove == nil {
			log.Fatalln("Unable to get previous move")
		}
		return previousMove.RowIndex, float64(color) * game.Evaluate(g, depth, nma.PlayerPiece)
	}
	bestVal := utils.NegativeInfinity
	bestMove := int32(-1)
	bestVal = utils.NegativeInfinity
	for _, move := range g.ListValidMoves() {
		err := g.MakeMove(move)
		if err != nil {
			log.Fatalln("Unexpected move error", err)
		}
		_, moveVal := nma.negamax(g, depth+1, -color)
		moveVal *= -1
		if utils.GreaterThanFloat64(moveVal, bestVal) {
			bestVal = moveVal
			bestMove = move
		}
		g.UndoMove()
	}
	return bestMove, bestVal
}
