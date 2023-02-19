package graphics

import "github.com/hajimehoshi/ebiten/v2"

type Position struct {
	X int
	Y int
}

type Sprite struct {
	imageWidth  int
	imageHeight int

	sprite ebiten.Image
}
