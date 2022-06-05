package game

import (
	"log"
)

var scoreDirections = []direction{
	down,
	right,
	downRight,
	downLeft,
}

var potentialDirections = []direction{
	up,
	left,
	right,
	upRight,
	upLeft,
	downRight,
	downLeft,
}

// Evaluate returns a score for the current board given a player piece and
// depth (number of moves required).
func Evaluate(g *Game, depth int32, playerPiece Piece) float64 {
	if g.IsGameOver() {
		if g.State == Draw {
			return 0
		}
		previousMove := g.GetPreviousMove()
		if previousMove == nil {
			log.Fatalln("Unable to get previous move in game over state")
		}
		if previousMove.Piece == playerPiece {
			return float64(1000 - depth)
		} else {
			return float64(-1000 + depth)
		}
	}
	var score float64
	for rowIndex, row := range g.Board {
		for colIndex, cell := range row {
			if cell != playerPiece {
				continue
			}
			// Favour moves that connect rows, columns or diagonals.
			for _, d := range scoreDirections {
				adjacentRowIndex := rowIndex + d.rowDirection
				adjacentColIndex := colIndex + d.colDirection
				if IsWithinBounds(adjacentRowIndex, adjacentColIndex) {
					adjacentPiece := g.Board[adjacentRowIndex][adjacentColIndex]
					if adjacentPiece == playerPiece {
						score++
					}
				}
			}
			// Favour moves that have potential for connect four with empty
			// spaces.
			for _, d := range potentialDirections {
				adjacentRowIndex := rowIndex + d.rowDirection
				adjacentColIndex := colIndex + d.colDirection
				var counter = 0
				for IsWithinBounds(adjacentRowIndex, adjacentColIndex) {
					adjacentPiece := g.Board[adjacentRowIndex][adjacentColIndex]
					if counter < 4 && (adjacentPiece == playerPiece || adjacentPiece == Empty) {
						score += 0.1
					} else {
						break
					}
					counter++
					adjacentRowIndex += d.rowDirection
					adjacentColIndex += d.colDirection
				}
			}
		}
	}
	return score
}
