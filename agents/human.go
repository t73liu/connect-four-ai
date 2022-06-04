package agents

import (
	"fmt"

	"github.com/t73liu/connectfourai/game"
)

type HumanAgent struct {
}

func (ha *HumanAgent) GetMove(g *game.Game) int32 {
	var move int32
	for {
		fmt.Println("Make your move (0-indexed): ")
		if _, err := fmt.Scanln(&move); err != nil {
			fmt.Printf("Invalid input: %s\n", err)
			continue
		}
		if !g.IsValidMove(move) {
			fmt.Printf("%d is an invalid move\n", move)
			continue
		}
		break
	}
	return move
}
