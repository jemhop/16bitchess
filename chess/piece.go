package chess

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
	side      int
	piecetype int
}
