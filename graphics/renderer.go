package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// draws a square centrally at position
func DrawSquare(size int, position Position, clr color.Color, img *ebiten.Image) {
	squareImage := ebiten.NewImage(size, size)

	squareImage.Fill(clr)

	translation := ebiten.GeoM{}
	translation.Translate(float64(position.X), float64(position.Y))

	drawOptions := ebiten.DrawImageOptions{Filter: ebiten.FilterLinear, GeoM: translation}

	img.DrawImage(squareImage, &drawOptions)
}
