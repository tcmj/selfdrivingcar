package pugagui

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)


/*
     Main GUI part!
	 Here it comes together!


*/





const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600
	DEBUG         = false
)

var redImage = ebiten.NewImage(3, 3)

type MyGame struct {
	count int

	vertices []ebiten.Vertex
	indices  []uint16

	aa         bool
	showCenter bool
}

func init() {
	log.Printf("Initializing...")
	redImage.Fill(color.RGBA{255, 0, 0, 255})
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
}

func NewGame(amount int, colorful bool) *MyGame {
	g := &MyGame{}
	return g
}

func (g *MyGame) Update() error {
	g.count++

	g.count++
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.aa = !g.aa
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.showCenter = !g.showCenter
	}
	return nil

}

func (g *MyGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	//fmt.Println("Game.Layout... ",outsideWidth, outsideHeight,screenWidth, screenHeight)
	//return outsideWidth /2, outsideHeight/2
	return outsideWidth, outsideHeight
}

func (g *MyGame) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x35, 0x15, 0xf5, 0xff})
 
	//lerper := &Player{}
	//NewLerper().DrawIT(g, screen)
	NewLerper().DrawSimple(g, screen)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()), 10, SCREEN_HEIGHT-20)

}

// BestMasco returns some Text
func BestMasco() string {
	return "Pug"

}
