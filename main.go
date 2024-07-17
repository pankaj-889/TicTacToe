package main

import (
	. "tictactoe/service"
)

func main() {

	game := Game{}
	game.TicTacToeGame()
	game.StartGame()

	/*

		Player
		piece
		turn


		Board
		size
		grid of PlayingPiece

		PieceType
		enum

		PlayingPiece
		PieceType

		Game
		board
		players
		turn
	*/
}
