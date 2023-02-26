package chess

import (
	"16bchess/graphics"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	NONE = iota
	PAWN
	KNIGHT
	ROOK
	BISHOP
	KING
	QUEEN
)

const (
	WHITE = iota
	BLACK
)

type Piece struct {
	sprite    ebiten.Image
	side      int
	Piecetype int
}

func (p Piece) Render(image *ebiten.Image, sheets graphics.SpritesheetManager, squareSize int, square Square) {

	pieceImage := ebiten.Image{}

	if square.Piece.side == BLACK {
		pieceImage = sheets.Sheets["blackpieces"].Sprites[square.Piece.Piecetype-1]
	} else {
		pieceImage = sheets.Sheets["whitepieces"].Sprites[square.Piece.Piecetype-1]
	}

	pieceSizeX := pieceImage.Bounds().Max.X
	pieceSizeY := pieceImage.Bounds().Max.Y
	pieceTranslation := ebiten.GeoM{}

	xScale := (float64(squareSize-4) / float64(pieceImage.Bounds().Max.X))
	yScale := (float64(squareSize-4) / float64(pieceImage.Bounds().Max.Y))

	pieceTranslation.Scale(xScale, yScale)
	pieceTranslation.Translate((float64(squareSize)-float64(pieceSizeX))/2, (float64(squareSize)-float64(pieceSizeY))/2)

	image.DrawImage(&pieceImage, &ebiten.DrawImageOptions{GeoM: pieceTranslation})
}
