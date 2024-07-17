package service

import (
	"errors"
	"fmt"
	. "tictactoe/constants"
	. "tictactoe/model"
)

type Game struct {
	Board       *Board
	Player      []Player
	CurrentTurn int
}

func (g *Game) TicTacToeGame() {
	g.initializeGame()
}

func (g *Game) initializeGame() {
	valX := StringToPiece("X")
	pX := PlayingPiece{
		PieceType: &valX,
	}

	valO := StringToPiece("O")
	pO := PlayingPiece{
		PieceType: &valO,
	}

	p1 := Player{
		Name:  "player1",
		Piece: pX,
	}

	p2 := Player{
		Name:  "player2",
		Piece: pO,
	}

	g.Player = append(g.Player, p1)
	g.Player = append(g.Player, p2)
	g.Board = NewBoard(3)
	g.CurrentTurn = 0
}

func (g *Game) StartGame() {
	for !g.isGameOver() {
		g.Board.PrintBoard()
		player := g.Player[g.CurrentTurn]
		fmt.Printf("%s's turn. Enter row and column (0, 1, or 2): ", player.Name)
		var row, col int
		fmt.Scanf("%d %d", &row, &col)

		err := g.makeMove(row, col, player.Piece)
		if err != nil {
			fmt.Println("Invalid move:", err)
			continue
		}

		if g.isWinningMove(row, col, player.Piece) {
			g.Board.PrintBoard()
			fmt.Printf("%s wins!\n", player.Name)
			return
		}

		g.CurrentTurn = (g.CurrentTurn + 1) % 2
	}
	g.Board.PrintBoard()
	fmt.Println("It's a draw!")
}

func (g *Game) isGameOver() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Board.Grid[i][j].PieceType == nil {
				return false
			}
		}
	}
	return true
}

func (g *Game) makeMove(row, col int, piece PlayingPiece) error {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return errors.New("move out of bounds")
	}
	if g.Board.Grid[row][col].PieceType != nil {
		return errors.New("cell already occupied")
	}

	g.Board.Grid[row][col] = piece
	return nil
}

func (g *Game) isWinningMove(row, col int, piece PlayingPiece) bool {
	// Check row
	if g.Board.Grid[row][0].PieceType == piece.PieceType && g.Board.Grid[row][1].PieceType == piece.PieceType && g.Board.Grid[row][2].PieceType == piece.PieceType {
		return true
	}
	// Check column
	if g.Board.Grid[0][col].PieceType == piece.PieceType && g.Board.Grid[1][col].PieceType == piece.PieceType && g.Board.Grid[2][col].PieceType == piece.PieceType {
		return true
	}
	// Check diagonals
	if row == col && g.Board.Grid[0][0].PieceType == piece.PieceType && g.Board.Grid[1][1].PieceType == piece.PieceType && g.Board.Grid[2][2].PieceType == piece.PieceType {
		return true
	}
	if row+col == 2 && g.Board.Grid[0][2].PieceType == piece.PieceType && g.Board.Grid[1][1].PieceType == piece.PieceType && g.Board.Grid[2][0].PieceType == piece.PieceType {
		return true
	}
	return false
}
