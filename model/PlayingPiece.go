package model

import . "tictactoe/constants"

type PlayingPiece struct {
	PieceType *PieceType
}

func (p *PlayingPiece) GetPiece() *PlayingPiece {
	return &PlayingPiece{
		PieceType: p.PieceType,
	}
}
