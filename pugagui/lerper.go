package pugagui

import (
	"bytes"
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tcmj/selfdrivingcar/pug"
)

type Player struct {
	PosX, PosY float64
}

var (
	whiteImage = ebiten.NewImage(3, 3)

	// whiteSubImage is an internal sub image of whiteImage.
	// Use whiteSubImage at DrawTriangles instead of whiteImage in order to avoid bleeding edges.
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

	strokeImage = ebiten.NewImage(4, 4)

	uiFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		panic(err)
	}
	uiFaceSource = s

	whiteImage.Fill(color.White)

	strokeImage.Fill(color.RGBA{5, 24, 54, 255})
}

func NewLerper() *Player {
	asdf := &Player{}
	return asdf
}

func DrawDot(g *MyGame, screen *ebiten.Image, point pug.POS, label string, isRed ...bool) {
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

func DrawLine(g *MyGame, screen *ebiten.Image, pA pug.POS, pB pug.POS) {
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
		//cm.Translate(255,2,2,255)
		//cm.Scale(0.2, 0.5, 0.3, 1.0)

		colorm.DrawTriangles(screen, vs, is, strokeImage, cm, &colorm.DrawTrianglesOptions{
			AntiAlias: g.aa,
		})

	}

}
func (p *Player) DrawSimple(g *MyGame, screen *ebiten.Image) {

	//fmt.Println("generatePlayerImage...")

	screen.Fill(color.RGBA{113, 143, 191, 255})

	A := pug.POS{X: 200, Y: 150}
	B := pug.POS{X: 150, Y: 210}
	C := pug.POS{X: 50, Y: 100}
	D := pug.POS{X: 250, Y: 200}

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

	msg := fmt.Sprintf(`FPS: %0.2f, TPS: %0.2f
Press A to switch anti-aliasing : %t.
Press C to switch color to draw the strokes : %t
Press Up/Down to adjust Speed : %f
Lerp - Percentage: %0.2f`, ebiten.ActualFPS(), ebiten.ActualTPS(), g.aa, g.strokeLineFlag, g.addendum, t)
	ebitenutil.DebugPrint(screen, msg)
	intersec, offset := pug.GetIntersection(A,B,C,D)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%v (offset=%.2f)",intersec,offset), 111,111)
	DrawDot(g, screen, *intersec, "I")
}
