package chess

import (
	"strings"
	"unicode"
)

const defaultFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

// Example FEN string
// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
// Applies a fen string to the given board (if fen=="", applies default position)
// Thanks to: https://github.com/SebLague/Chess-AI/blob/main/Assets/Scripts/Core/FenUtility.cs
func applyFenString(fen string, board *Board) {

	if fen == "" {
		fen = defaultFen
	}

	sections := strings.Split(fen, " ")

	rank := 7
	file := 0

	for _, r := range sections[0] {
		if r == '/' {
			file = 0
			rank--
		} else {
			if unicode.IsDigit(r) {
				file += int(r)
			} else {
				pieceType, side := runeToPiece(r)

				piece := Piece{Piecetype: pieceType, side: side}

				board.Squares[file][rank].Piece = piece

				file++
			}
		}
	}

	board.whiteToMove = (sections[1] == "w")

	//implement castling rights, en passant, and clocks later

}

func runeToPiece(r rune) (piece int, side int) {
	lowercase := unicode.ToLower(r)

	switch lowercase {
	case 'p':
		piece = PAWN
	case 'n':
		piece = KNIGHT
	case 'b':
		piece = BISHOP
	case 'r':
		piece = ROOK
	case 'q':
		piece = QUEEN
	case 'k':
		piece = KING
	}

	if unicode.IsLower(r) {
		side = BLACK
	} else {
		side = WHITE
	}

	return piece, side
}
