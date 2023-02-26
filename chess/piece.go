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

func (p *Piece) assignSprite(spritesheets graphics.SpritesheetManager) {
	activeSheet := graphics.Spritesheet{}
	if p.side == BLACK {
		activeSheet = spritesheets.Sheets["blackpieces"]
	} else {
		activeSheet = spritesheets.Sheets["whitepieces"]
	}

	p.sprite = activeSheet.Sprites[p.Piecetype]
}
