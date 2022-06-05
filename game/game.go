package game

import "fmt"

const (
	numOfRows    = 6
	numOfColumns = 7
)

type State int32

const (
	PlayerOneTurn State = iota
	PlayerTwoTurn
	PlayerOneWin
	PlayerTwoWin
	Draw
)

func (s State) String() string {
	switch s {
	case PlayerOneTurn:
		return "Player One Turn"
	case PlayerTwoTurn:
		return "Player Two Turn"
	case PlayerOneWin:
		return "Player One Win"
	case PlayerTwoWin:
		return "Player Two Win"
	case Draw:
		return "Draw"
	}
	return "Unknown GameState"
}

type Piece int32

const (
	Empty Piece = iota
	PlayerOnePiece
	PlayerTwoPiece
)

func (p Piece) String() string {
	switch p {
	case Empty:
		return " "
	case PlayerOnePiece:
		return "X"
	case PlayerTwoPiece:
		return "O"
	}
	return "Unknown Piece"
}

type Move struct {
	Piece       Piece
	RowIndex    int32
	ColumnIndex int32
}

type Game struct {
	State       State
	Board       [numOfRows][numOfColumns]Piece
	MoveHistory []*Move
}

func (g *Game) PrintBoard() {
	for _, row := range g.Board {
		for i, cell := range row {
			fmt.Printf("| %s ", cell)
			if i == numOfColumns-1 {
				fmt.Println("|")
			}
		}
		fmt.Println("-----------------------------")
	}
	fmt.Printf("State: %s\n", g.State)
}

func (g *Game) ListValidMoves() []int32 {
	switch g.State {
	case PlayerOneTurn, PlayerTwoTurn:
		validMoves := make([]int32, 0, numOfColumns)
		for i, cell := range g.Board[0] {
			if cell == Empty {
				validMoves = append(validMoves, int32(i))
			}
		}
		return validMoves
	}
	return nil
}

func (g *Game) IsValidMove(move int32) bool {
	if move >= numOfColumns || move < 0 {
		return false
	}
	return g.Board[0][move] == Empty
}

func (g *Game) GetPreviousMove() *Move {
	moveCount := len(g.MoveHistory)
	if moveCount == 0 {
		return nil
	}
	return g.MoveHistory[moveCount-1]
}

func (g *Game) IsGameOver() bool {
	return g.State == PlayerOneWin || g.State == PlayerTwoWin || g.State == Draw
}

func (g *Game) UndoMove() {
	previousMove := g.GetPreviousMove()
	if previousMove == nil {
		return
	}
	g.MoveHistory = g.MoveHistory[:len(g.MoveHistory)-1]
	g.Board[previousMove.RowIndex][previousMove.ColumnIndex] = Empty
	switch previousMove.Piece {
	case PlayerOnePiece:
		g.State = PlayerOneTurn
	case PlayerTwoPiece:
		g.State = PlayerTwoTurn
	}
}

func (g *Game) MakeMove(move int32) error {
	piece := Empty
	switch g.State {
	case PlayerOneTurn:
		piece = PlayerOnePiece
	case PlayerTwoTurn:
		piece = PlayerTwoPiece
	}
	if !g.IsValidMove(move) {
		return fmt.Errorf("%d is not a valid move", move)
	}
	for i := range g.Board {
		rowIndex := numOfRows - i - 1
		if g.Board[rowIndex][move] == Empty {
			g.Board[rowIndex][move] = piece
			g.MoveHistory = append(g.MoveHistory, &Move{
				Piece:       piece,
				RowIndex:    int32(rowIndex),
				ColumnIndex: move,
			})
			g.updateState()
			break
		}
	}
	return nil
}

func (g *Game) updateState() {
	currentState := g.State
	if g.IsWinningMove() {
		switch currentState {
		case PlayerOneTurn:
			g.State = PlayerOneWin
		case PlayerTwoTurn:
			g.State = PlayerTwoWin
		}
	} else if len(g.ListValidMoves()) != 0 {
		switch currentState {
		case PlayerOneTurn:
			g.State = PlayerTwoTurn
		case PlayerTwoTurn:
			g.State = PlayerOneTurn
		}
	} else {
		g.State = Draw
	}
}

var directions = [][]int{
	// Up
	{1, 0},
	// Right
	{0, 1},
	// Up-Right
	{1, 1},
	// Up-Left
	{1, -1},
}

var oppositeDirections = [][]int{
	// Down
	{-1, 0},
	// Left
	{0, -1},
	// Down-Left
	{-1, -1},
	// Down Right
	{-1, 1},
}

func (g *Game) IsWinningMove() bool {
	previousMove := g.GetPreviousMove()
	if previousMove == nil {
		return false
	}
	for i, direction := range directions {
		count := 1
		rowIndex := int(previousMove.RowIndex) + direction[0]
		colIndex := int(previousMove.ColumnIndex) + direction[1]
		for IsWithinBounds(rowIndex, colIndex) {
			if g.Board[rowIndex][colIndex] == previousMove.Piece {
				count++
				if count == 4 {
					return true
				}
			} else {
				break
			}
			rowIndex += direction[0]
			colIndex += direction[1]
		}
		oppositeDirection := oppositeDirections[i]
		rowIndex = int(previousMove.RowIndex) + oppositeDirection[0]
		colIndex = int(previousMove.ColumnIndex) + oppositeDirection[1]
		for IsWithinBounds(rowIndex, colIndex) {
			if g.Board[rowIndex][colIndex] == previousMove.Piece {
				count++
				if count == 4 {
					return true
				}
			} else {
				break
			}
			rowIndex += oppositeDirection[0]
			colIndex += oppositeDirection[1]
		}
	}
	return false
}

func IsWithinBounds(rowIndex, colIndex int) bool {
	return rowIndex >= 0 && rowIndex < numOfRows && colIndex >= 0 && colIndex < numOfColumns
}

func NewGame() *Game {
	return &Game{}
}
