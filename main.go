package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/t73liu/connectfourai/agents"
	"github.com/t73liu/connectfourai/game"
)

type agent interface {
	GetMove(g *game.Game) int32
}

func getAgent(agentType string, piece game.Piece) agent {
	switch agentType {
	case "human":
		return &agents.HumanAgent{}
	case "random":
		return &agents.RandomAgent{}
	case "minimax":
		return &agents.MiniMaxAgent{MaxDepth: 4, MaximizerPiece: piece}
	case "alphabeta":
		return &agents.AlphaBetaAgent{MaxDepth: 4, MaximizerPiece: piece}
	case "negamax":
		return nil
	case "minimax-pruning":
		return nil
	default:
		log.Fatalf("unrecognized agent type: %s", agentType)
	}
	return nil
}

func main() {
	playerOne := flag.String(
		"p1",
		"human",
		"Agent to use for player 1",
	)
	playerTwo := flag.String(
		"p2",
		"alphabeta",
		"Agent to use for player 2",
	)
	flag.Parse()

	agentOne := getAgent(*playerOne, game.PlayerOnePiece)
	agentTwo := getAgent(*playerTwo, game.PlayerTwoPiece)
	isPlayerOneTurn := true

	g := game.NewGame()
	for !g.IsGameOver() {
		g.PrintBoard()
		var move int32
		if isPlayerOneTurn {
			move = agentOne.GetMove(g)
		} else {
			move = agentTwo.GetMove(g)
		}
		if err := g.MakeMove(move); err != nil {
			fmt.Printf("Try again: %s\n", err)
			continue
		}
		isPlayerOneTurn = !isPlayerOneTurn
	}
	g.PrintBoard()
}
