package graphics

import (
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Spritesheet struct {
	Sprites                   []ebiten.Image
	spriteWidth, spriteHeight int
	path                      string
}

type SpritesheetManager struct {
	Sheets map[string]Spritesheet
}

type sheetMetadata struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// Initializes the spritesheet managers map of sheets
func (s *SpritesheetManager) LoadSpriteSheets(paths []string) {
	s.Sheets = make(map[string]Spritesheet)
	for _, path := range paths {
		fileName := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
		jsonPath := strings.TrimSuffix(path, filepath.Ext(path)) + ".json"
		metadata := loadSpritesheetMetadata(jsonPath)

		s.Sheets[fileName] = loadSpriteSheet(path, metadata.Width, metadata.Height)
	}
}

func loadSpritesheetMetadata(path string) sheetMetadata {
	bytes, _ := os.ReadFile(path)

	var metadata sheetMetadata
	err := json.Unmarshal(bytes, &metadata)

	if err != nil {
		log.Fatal(err)
	}

	return metadata
}

func loadSpriteSheet(path string, width int, height int) Spritesheet {
	spriteSheet := Spritesheet{}

	spriteSheet.path = path
	spriteSheet.spriteWidth = width
	spriteSheet.spriteHeight = height

	spriteSheetImage := openImage(path)

	sliced := sliceImage1D(spriteSheetImage, width)

	fmt.Println(len(sliced))

	for _, s := range sliced {
		spriteSheet.Sprites = append(spriteSheet.Sprites, *ebiten.NewImageFromImage(s))
	}

	return spriteSheet
}

// Heavily modified ChatGPT code
// Assumes the spritesheet is 1d. I'll change that later
func sliceImage1D(img image.Image, gridSize int) []image.Image {
	bounds := img.Bounds()
	width, _ := bounds.Max.X, bounds.Max.Y

	// Calculate the size of each sub-image
	subWidth := width / gridSize

	// Create a slice to hold the sub-images
	subImages := make([]image.Image, subWidth)

	for col := 0; col < subWidth; col++ {
		subImg := img.(interface {
			SubImage(r image.Rectangle) image.Image
		}).SubImage(image.Rect(col*gridSize, 0, col*gridSize+gridSize, gridSize))

		subImages[col] = subImg
	}

	return subImages
}

func openImage(filePath string) image.Image {
	imgFile, err := os.Open(filePath)
	if err != nil {
		log.Println("Cannot read file:", err)
	}

	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}
	return img
}
