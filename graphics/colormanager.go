package graphics

import (
	"image/color"
)

type Colours struct {
	Board       color.RGBA
	Whitesquare color.RGBA
	Blacksquare color.RGBA
}

func NewColours() *Colours {
	newCol := Colours{}
	newCol.Board = color.RGBA{35, 27, 26, 255}
	newCol.Whitesquare = color.RGBA{239, 248, 226, 255}
	newCol.Blacksquare = color.RGBA{71, 121, 152, 255}

	return &newCol
}
