package main

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tcmj/selfdrivingcar/pug"
	"image"
	"image/color"
	"log"
)

const APP_TITLE = "app#1::Lerper"

const (
	SCREEN_WIDTH  = 888
	SCREEN_HEIGHT = 777
)

var (
	whiteImage = ebiten.NewImage(3, 3)

	// whiteSubImage is an internal sub image of whiteImage.
	// Use whiteSubImage at DrawTriangles instead of whiteImage in order to avoid bleeding edges.
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

	strokeImage = ebiten.NewImage(4, 4)

	uiFaceSource *text.GoTextFaceSource
)

type MyGameApp1 struct {
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
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		panic(err)
	}
	uiFaceSource = s

	whiteImage.Fill(color.White)

	strokeImage.Fill(color.RGBA{5, 24, 54, 255})
}

func NewGameApp1(amount int, colorful bool) *MyGameApp1 {
	g := &MyGameApp1{
		count:    -1.1,
		addendum: 0.001,
		leftVal:  20,
		rghtVal:  120,
		zoom:     2,
	}
	return g
}

func (g *MyGameApp1) Update() error {
	g.count += g.addendum
	if g.count > 2.0 {
		fmt.Println("resetting", g.count)
		g.count = -1.0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.aa = !g.aa
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.strokeLineFlag = !g.strokeLineFlag
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.addendum += 0.001
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.addendum -= 0.001
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		g.leftVal += 11.0
		g.rghtVal -= 11.0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		g.rghtVal += 11.0
		g.leftVal -= 11.0
	}
	_, dy := ebiten.Wheel()
	if dy > 0 {
		g.zoom += 1
	} else if dy < 0 && g.zoom != 1 {
		g.zoom -= 1
	}
	mx, my := ebiten.CursorPosition()
	g.mouseX = mx
	g.mouseY = my
	return nil
}

func (g *MyGameApp1) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / g.zoom, outsideHeight / g.zoom
}

func (g *MyGameApp1) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x35, 0x15, 0xf5, 0xff})

	//lerper := &Player{}
	//NewLerper().DrawIT(g, screen)
	DrawSimple(g, screen)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()), 10, SCREEN_HEIGHT-20)

}

func DrawSimple(g *MyGameApp1, screen *ebiten.Image) {

	//fmt.Println("generatePlayerImage...")

	screen.Fill(color.RGBA{113, 143, 191, 255})

	A := pug.POS{X: 250, Y: g.rghtVal}
	B := pug.POS{X: g.leftVal, Y: 260}
	C := pug.POS{X: 100, Y: 150}
	D := pug.POS{X: 300, Y: 250}

	DrawLine(g, screen, A, B)
	DrawLine(g, screen, C, D)

	//DrawDot(g, screen, POS{100, 150}, "A")
	DrawDot(g, screen, A, "A")
	DrawDot(g, screen, B, "B")
	DrawDot(g, screen, C, "C")
	DrawDot(g, screen, D, "D")

	var t float64 = g.count

	M := pug.POS{
		X: pug.Lerp(A.X, B.X, t),
		Y: pug.Lerp(A.Y, B.Y, t),
	}
	DrawDot(g, screen, M, "M", t < 0 || t > 1)
	N := pug.POS{
		X: pug.Lerp(C.X, D.X, t),
		Y: pug.Lerp(C.Y, D.Y, t),
	}
	DrawDot(g, screen, N, "N", t < 0 || t > 1)

	msg := fmt.Sprintf(
		`FPS: %0.2f, TPS: %0.2f, [A]ntiAliasing: %t, [C]olor-Switch: %t, [MouseWheel] to Zoom: %d
Press [Up]/[Down] to adjust Speed : %f
Press [Left]/[Right] to move Points (Ay/Bx) : %.2f 
Lerp - Percentage: %0.2f`,
		ebiten.ActualFPS(), ebiten.ActualTPS(), g.aa, g.strokeLineFlag, g.zoom,
		g.addendum, g.leftVal, t)
	ebitenutil.DebugPrint(screen, msg)
	intersec, offset := pug.GetIntersection(A, B, C, D)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%v (offset=%.2f)", intersec, offset), 70, 300)

	if intersec != nil {

		DrawDot(g, screen, *intersec, "I")
	}
}

func DrawDot(g *MyGameApp1, screen *ebiten.Image, point pug.POS, label string, isRed ...bool) {
	vector.DrawFilledCircle(screen, float32(point.X), float32(point.Y), 8.0, color.White, g.aa)
	vector.StrokeCircle(screen, float32(point.X), float32(point.Y), 8.0, 1.0, color.Black, g.aa)
	dO := &text.DrawOptions{}
	dO.GeoM.Translate(float64(point.X), float64(point.Y))
	if len(isRed) > 0 && isRed[0] {

		dO.ColorScale.ScaleWithColor(color.RGBA{225, 24, 54, 255})
	} else {

		dO.ColorScale.ScaleWithColor(color.Black)
	}
	dO.LineSpacing = 1
	dO.PrimaryAlign = text.AlignCenter
	dO.SecondaryAlign = text.AlignCenter
	text.Draw(screen, label, &text.GoTextFace{
		Source: uiFaceSource,
		Size:   12,
	}, dO)

}

func DrawLine(g *MyGameApp1, screen *ebiten.Image, pA pug.POS, pB pug.POS) {
	op := &vector.StrokeOptions{}
	op.Width = 2
	op.LineJoin = vector.LineJoinMiter
	var path vector.Path
	path.MoveTo(float32(pA.X), float32(pA.Y))
	path.LineTo(float32(pB.X), float32(pB.Y))
	path.Close()
	vs, is := path.AppendVerticesAndIndicesForStroke(nil, nil, op)

	if g.strokeLineFlag {
		op.Width = 1
		vs, is := path.AppendVerticesAndIndicesForStroke(g.vertices[:0], g.indices[:0], op)
		for i := range vs {
			vs[i].SrcX = 1
			vs[i].SrcY = 1
			vs[i].ColorR = 1
			vs[i].ColorG = 0
			vs[i].ColorB = 0
			vs[i].ColorA = 1
		}
		screen.DrawTriangles(vs, is, whiteSubImage, &ebiten.DrawTrianglesOptions{
			AntiAlias: g.aa,
		})
	} else {
		var cm colorm.ColorM
		colorm.DrawTriangles(screen, vs, is, strokeImage, cm, &colorm.DrawTrianglesOptions{
			AntiAlias: g.aa,
		})
	}
}

func main() {
	fmt.Println(APP_TITLE)
	ebiten.SetWindowIcon(pug.GetIconGopher())
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle(APP_TITLE)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(NewGameApp1(100, false)); err != nil {
		log.Fatal(err)
	}
}
