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
	piece       Piece
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

func (g *Game) IsGameOver() bool {
	return g.State == PlayerOneWin || g.State == PlayerTwoWin
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
				piece:       piece,
				RowIndex:    int32(rowIndex),
				ColumnIndex: move,
			})
			g.updateState()
			break
		}
	}
	return nil
}

func (g *Game) UndoMove() {
	// TODO
}

func (g *Game) updateState() {
	currentState := g.State
	if g.isWinningMove() {
		switch currentState {
		case PlayerOneTurn:
			g.State = PlayerOneWin
		case PlayerTwoTurn:
			g.State = PlayerTwoWin
		}
	} else {
		switch currentState {
		case PlayerOneTurn:
			g.State = PlayerTwoTurn
		case PlayerTwoTurn:
			g.State = PlayerOneTurn
		}
	}
}

var directions = [][]int{
	// Up
	{1, 0},
	// Up-Right
	{1, 1},
	// Right
	{0, 1},
	// Down-Right
	{-1, 1},
	// Down
	{-1, 0},
	// Down-Left
	{-1, -1},
	// Left
	{0, -1},
	// Up-Left
	{1, -1},
}

func (g *Game) isWinningMove() bool {
	lastMove := g.MoveHistory[len(g.MoveHistory)-1]
	for _, direction := range directions {
		var count int
		rowIndex := int(lastMove.RowIndex)
		colIndex := int(lastMove.ColumnIndex)
		for isWithinBounds(rowIndex, colIndex) {
			if g.Board[rowIndex][colIndex] == lastMove.piece {
				count++
				if count == 4 {
					return true
				}
			} else {
				break
			}
			colIndex += direction[0]
			rowIndex += direction[1]
		}
	}
	return false
}

func isWithinBounds(rowIndex, colIndex int) bool {
	return rowIndex >= 0 && rowIndex < numOfRows && colIndex >= 0 && colIndex < numOfColumns
}

func NewGame() *Game {
	return &Game{}
}
