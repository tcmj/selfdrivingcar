package main

import (
    "fmt"
    "image/color"
    "log"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/tcmj/selfdrivingcar/pug"
)

const APP_TITLE = "app#0::Template"

const (
    SCREEN_WIDTH  = 1200
    SCREEN_HEIGHT = 800
)

var (
    strokeImage = ebiten.NewImage(4, 4)
)

type MyGameTemplate struct {
    count          float64
    addendum       float64
    leftVal        float64
    rghtVal        float64
    mouseX         int
    mouseY         int
    zoom           int
    vertices       []ebiten.Vertex
    indices        []uint16
    aa             bool
    strokeLineFlag bool
}

func init() {

    strokeImage.Fill(color.RGBA{33, 44, 54, 255})
}

func NewGameApp1(amount int, colorful bool) *MyGameTemplate {
    g := &MyGameTemplate{
        count:    -1.1,
        addendum: 0.001,
        leftVal:  20,
        rghtVal:  120,
        zoom:     2,
    }
    return g
}

func (g *MyGameTemplate) Update() error {

    if inpututil.IsKeyJustPressed(ebiten.KeyA) {
        g.aa = !g.aa
    }

    return nil
}

func (g *MyGameTemplate) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return outsideWidth / g.zoom, outsideHeight / g.zoom
}

func (g *MyGameTemplate) Draw(screen *ebiten.Image) {
    screen.Fill(color.RGBA{0x35, 0x44, 0x55, 0xff})

    ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()), 10, 20)

}

func main() {
    fmt.Println(APP_TITLE)
    ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
    ebiten.SetWindowIcon(pug.GetIconGopher())
    ebiten.SetWindowTitle(APP_TITLE)
    ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
    if err := ebiten.RunGame(NewGameApp1(100, false)); err != nil {
        log.Fatal(err)
    }
}
