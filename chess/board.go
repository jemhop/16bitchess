package chess

import (
	"16bchess/graphics"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PAWN = iota
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

// need to implement move history etc
type Board struct {
	Squares    [8][8]Square
	Size       int
	whiteColor color.Color
	blackColor color.Color
}

type Square struct {
	piece       int
	attacked_by []int
}

type Piece struct {
	team      int
	piecetype int
}

func NewBoard(size int, colours graphics.Colours) *Board {
	board := Board{}

	board.Size = size
	board.Squares = [8][8]Square{}
	board.blackColor = colours.Blacksquare
	board.whiteColor = colours.Whitesquare

	return &board
}

func (board Board) Render(screen *ebiten.Image, position graphics.Position) {
	boardImage := ebiten.NewImage(board.Size, board.Size)

	squareSize := board.Size / 8

	for y := range board.Squares {
		for x := range board.Squares[y] {
			//Find a cleaner way to do this

			squareCol := board.whiteColor
			if !board.determineSquareColour(x, y) {
				squareCol = board.blackColor
			}

			graphics.DrawSquare(squareSize, graphics.Position{X: x * squareSize, Y: y * squareSize}, squareCol, boardImage)

		}
	}

	translation := ebiten.GeoM{}
	translation.Translate(float64(position.X), float64(position.Y))

	screen.DrawImage(boardImage, &ebiten.DrawImageOptions{GeoM: translation})
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
