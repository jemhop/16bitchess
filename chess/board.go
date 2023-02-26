package chess

import (
	"16bchess/graphics"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// need to implement move history etc
type Board struct {
	Squares  [8][8]Square
	Size     int
	whiteCol color.Color
	blackCol color.Color

	whiteCastleKingside  bool
	whiteCastleQueenside bool
	blackCastleKingside  bool
	blackCastleQueenside bool

	whiteToMove bool
}

type Square struct {
	Piece       Piece
	col         color.Color
	attacked_by []int
}

func NewBoard(size int, colours graphics.Colours) *Board {
	board := Board{}

	board.Size = size
	board.blackCol = colours.Blacksquare
	board.whiteCol = colours.Whitesquare

	board = *setSquareColours(board)

	applyFenString("", &board)

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

func (board Board) Render(screen *ebiten.Image, position graphics.Position, sheets graphics.SpritesheetManager) {
	boardImage := ebiten.NewImage(board.Size, board.Size)

	squareSize := board.Size / 8

	for y := range board.Squares {
		for x := range board.Squares[y] {
			position := graphics.Position{X: x * squareSize, Y: y * squareSize}
			board.renderSquare(boardImage, position, board.Squares[x][y], sheets)
		}
	}

	translation := ebiten.GeoM{}
	translation.Translate(float64(position.X), float64(position.Y))

	screen.DrawImage(boardImage, &ebiten.DrawImageOptions{GeoM: translation})
}

func (board Board) renderSquare(screen *ebiten.Image, position graphics.Position, square Square, sheets graphics.SpritesheetManager) {
	squareSize := board.Size / len(board.Squares)
	squareImage := ebiten.NewImage(squareSize, squareSize)
	squareImage.Fill(square.col)

	translation := ebiten.GeoM{}
	translation.Translate(float64(position.X), float64(position.Y))

	if square.Piece.Piecetype != 0 {
		renderPiece(squareImage, sheets, squareSize, square)
	}

	screen.DrawImage(squareImage, &ebiten.DrawImageOptions{GeoM: translation})
}

func renderPiece(squareImage *ebiten.Image, sheets graphics.SpritesheetManager, squareSize int, square Square) {

	pieceImage := ebiten.Image{}

	if square.Piece.side == BLACK {
		pieceImage = sheets.Sheets["blackpieces"].Sprites[square.Piece.Piecetype-1]
	} else {
		pieceImage = sheets.Sheets["whitepieces"].Sprites[square.Piece.Piecetype-1]
	}

	pieceTranslation := ebiten.GeoM{}

	xScale := (float64(squareSize) / float64(pieceImage.Bounds().Max.X))
	yScale := (float64(squareSize) / float64(pieceImage.Bounds().Max.Y))

	pieceTranslation.Scale(xScale, yScale)

	squareImage.DrawImage(&pieceImage, &ebiten.DrawImageOptions{GeoM: pieceTranslation})
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
