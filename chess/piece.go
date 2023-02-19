package chess

import "16bchess/graphics"

const (
	NONE = iota
	PAWN
	KNIGHT
	BISHOP
	ROOK
	KING
	QUEEN
)

const (
	WHITE = iota
	BLACK
)

type Piece struct {
	sprite    graphics.Sprite
	team      int
	piecetype int
}
