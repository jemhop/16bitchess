package chess

import (
	"16bchess/graphics"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// need to implement move history etc
type Board struct {
	Squares  [8][8]Square
	Size     int
	whiteCol color.Color
	blackCol color.Color
}

type Square struct {
	piece       Piece
	col         color.Color
	attacked_by []int
}

func NewBoard(size int, colours graphics.Colours) *Board {
	board := Board{}

	board.Size = size
	board.blackCol = colours.Blacksquare
	board.whiteCol = colours.Whitesquare

	board = *setSquareColours(board)

	return &board
}

func setSquareColours(board Board) *Board {
	newBoard := board

	for y := range board.Squares {
		for x := range board.Squares[y] {
			//Find a cleaner way to do this

			newBoard.Squares[y][x].col = board.whiteCol
			if !board.determineSquareColour(x, y) {
				newBoard.Squares[y][x].col = board.blackCol
			}

		}
	}

	return &newBoard
}

/* func (board Board) FenStringLoader(FEN string) {

} */

func (board Board) Render(screen *ebiten.Image, position graphics.Position) {
	boardImage := ebiten.NewImage(board.Size, board.Size)

	squareSize := board.Size / 8

	for y := range board.Squares {
		for x := range board.Squares[y] {
			position := graphics.Position{X: x * squareSize, Y: y * squareSize}
			board.renderSquare(boardImage, position, board.Squares[x][y])
		}
	}

	translation := ebiten.GeoM{}
	translation.Translate(float64(position.X), float64(position.Y))

	screen.DrawImage(boardImage, &ebiten.DrawImageOptions{GeoM: translation})
}

// a bunch of placeholder code for drawing pieces before sprites are implemented
func (board Board) renderSquare(screen *ebiten.Image, position graphics.Position, square Square) {
	squareSize := board.Size / len(board.Squares)
	squareImage := ebiten.NewImage(squareSize, squareSize)
	squareImage.Fill(square.col)

	pieceImage := ebiten.NewImage(squareSize, squareSize)

	pieceChar := getPieceCharacter(square.piece)

	if pieceChar != "" {
		ebitenutil.DebugPrint(pieceImage, pieceChar)
	}

	translation := ebiten.GeoM{}
	translation.Translate(float64(position.X), float64(position.Y))

	squareImage.DrawImage(pieceImage, &ebiten.DrawImageOptions{})

	screen.DrawImage(squareImage, &ebiten.DrawImageOptions{GeoM: translation})
}

func getPieceCharacter(piece Piece) string {
	pieceChar := ""

	switch piece.piecetype {
	case PAWN:
		pieceChar = "p"
	case KNIGHT:
		pieceChar = "n"
	case BISHOP:
		pieceChar = "b"
	case ROOK:
		pieceChar = "r"
	case QUEEN:
		pieceChar = "q"
	case KING:
		pieceChar = "k"
	}

	if piece.team == WHITE {
		pieceChar = strings.ToUpper(pieceChar)
	}

	return pieceChar
}

// Function to determine the colour of a square based on the x and y value. Returns true for white and false for black
func (board Board) determineSquareColour(x int, y int) bool {
	if y%2 != 0 {
		if x%2 == 0 {
			return false
		} else {
			return true
		}
	} else {
		if x%2 == 0 {
			return true
		} else {
			return false
		}
	}
}
