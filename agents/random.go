package agents

import (
	"log"
	"math/rand"

	"github.com/t73liu/connectfourai/game"
)

type RandomAgent struct {
}

func (ra *RandomAgent) GetMove(g *game.Game) int32 {
	validMoves := g.ListValidMoves()
	if len(validMoves) == 0 {
		log.Fatalln("no valid moves left")
	}
	moveIndex := rand.Intn(len(validMoves))
	return validMoves[moveIndex]
}
