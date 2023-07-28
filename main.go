package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"sim/engine"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Game struct{}

var R engine.Object
var Earth engine.Object
var arrayOfObjects []*engine.Object
var (
	mplusNormalFont font.Face
	mplusBigFont    font.Face
	jaKanjis        = []rune{}
)

func (g *Game) Update() error {
	//R.ApplyForce(&engine.Vector2D{2, 2}, 1/60)
	engine.Run(&arrayOfObjects, 0.015)
	fmt.Println(R)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	text.Draw(screen, "FPS : "+strconv.Itoa(int(ebiten.ActualFPS())), mplusNormalFont, 10, 24, color.White)
	text.Draw(screen, "TPS : "+strconv.Itoa(int(ebiten.ActualTPS())), mplusNormalFont, 10, 48, color.White)

	//vector.StrokeLine(screen, x1, y1, x2, y2, 10, color.RGBA{0xff, 0, 0, 0xff}, true)
	vector.DrawFilledCircle(screen, Earth.Location.X, Earth.Location.Y, 50, color.RGBA{1, 1, 128, 0xff}, true)
	vector.DrawFilledCircle(screen, R.Location.X, R.Location.Y, 3, color.RGBA{0xff, 0, 0, 0xff}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1200, 600
}

func init() {
	pixeloidsans, err := ioutil.ReadFile("PixeloidSans.ttf")

	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(pixeloidsans)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    27, //9'un katlari olmasi idealdir
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull, // Use quantization to save glyph cache images.
	})
	if err != nil {
		log.Fatal(err)
	}

	// Adjust the line height.
	mplusBigFont = text.FaceWithLineHeight(mplusBigFont, 54)
}

func main() {

	R.Name = "test"
	R.Mass = 10
	R.Velocity.X = 15
	R.Velocity.Y = 0
	R.Location.X = 300
	R.Location.Y = 200
	R.Physical(true)

	Earth.Name = "Earth"
	Earth.Mass = 5.9722e24
	Earth.Radius = 6.3781e6
	Earth.Location.X = 300
	Earth.Location.Y = 300

	arrayOfObjects = append(arrayOfObjects, &R)
	arrayOfObjects = append(arrayOfObjects, &Earth)

	ebiten.SetWindowSize(1200, 600)
	ebiten.SetWindowTitle("simülasyon")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
