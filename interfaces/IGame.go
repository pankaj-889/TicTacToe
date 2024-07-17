package interfaces

import (
	. "tictactoe/model"
)

type IGame interface {
	AddPlayers(player Player)
	StartGame()
}

type IPlayingPiece interface {
	GetPiece() *PlayingPiece
}
