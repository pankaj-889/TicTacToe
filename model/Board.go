package model

import (
	"fmt"
)

type Board struct {
	Size int
	Grid [][]PlayingPiece
}

func NewBoard(size int) *Board {
	Grid := make([][]PlayingPiece, size)
	for i := range Grid {
		Grid[i] = make([]PlayingPiece, size)
		for j := range Grid[i] {
			Grid[i][j] = PlayingPiece{}
		}
	}
	return &Board{
		Size: size,
		Grid: Grid,
	}
}

func (b *Board) PrintBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.Grid[i][j].PieceType == nil {
				fmt.Print("_ ")
			} else {
				fmt.Printf("%s ", b.Grid[i][j])
			}
		}
		fmt.Println()
	}
}
