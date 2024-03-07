package pugagui

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/tcmj/selfdrivingcar/pugutils"
)

type Player struct {
	PosX, PosY float64
}

type POS struct {
	x, y float32
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

func DrawDot(g *MyGame, screen *ebiten.Image, point POS, label string, isRed ...bool) {
	vector.DrawFilledCircle(screen, point.x, point.y, 8.0, color.White, g.aa)
	vector.StrokeCircle(screen, point.x, point.y, 8.0, 1.0, color.Black, g.aa)
	dO := &text.DrawOptions{}
	dO.GeoM.Translate(float64(point.x), float64(point.y))
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

func DrawLine(g *MyGame, screen *ebiten.Image, pA POS, pB POS) {
	op := &vector.StrokeOptions{}
	op.Width = 2
	op.LineJoin = vector.LineJoinMiter

	var path vector.Path

	path.MoveTo(pA.x, pA.y)
	path.LineTo(pB.x, pB.y)

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

	A := POS{200, 150}
	B := POS{150, 250}
	C := POS{50, 100}
	D := POS{250, 200}

	DrawLine(g, screen, A, B)
	DrawLine(g, screen, C, D)

	//DrawDot(g, screen, POS{100, 150}, "A")
	DrawDot(g, screen, A, "A")
	DrawDot(g, screen, B, "B")
	DrawDot(g, screen, C, "C")
	DrawDot(g, screen, D, "D")

	var t float32 = g.count

	M := POS{
		pugutils.Lerp32(A.x, B.x, t),
		pugutils.Lerp32(A.y, B.y, t),
	}
	DrawDot(g, screen, M, "M", t < 0 || t > 1)
	N := POS{
		pugutils.Lerp32(C.x, D.x, t),
		pugutils.Lerp32(C.y, D.y, t),
	}
	DrawDot(g, screen, N, "N", t < 0 || t > 1)

	msg := fmt.Sprintf(`FPS: %0.2f, TPS: %0.2f
Press A to switch anti-aliasing : %t.
Press C to switch color to draw the strokes : %t
Press Up/Down to adjust Speed : %f
Lerp - Percentage: %0.2f`, ebiten.ActualFPS(), ebiten.ActualTPS(), g.aa, g.strokeLineFlag, g.addendum, t)
	ebitenutil.DebugPrint(screen, msg)

}

func (p *Player) DrawIT(g *MyGame, screen *ebiten.Image) {
	target := screen

	joins := []vector.LineJoin{
		vector.LineJoinMiter,
		vector.LineJoinMiter,
		vector.LineJoinBevel,
		vector.LineJoinRound,
	}
	caps := []vector.LineCap{
		vector.LineCapButt,
		vector.LineCapRound,
		vector.LineCapSquare,
	}

	ow, oh := target.Bounds().Dx(), target.Bounds().Dy()
	size := min(ow/(len(joins)+1), oh/(len(caps)+1))
	offsetX, offsetY := (ow-size*len(joins))/2, (oh-size*len(caps))/2

	// Render the lines on the target.
	for j, cap := range caps {
		for i, join := range joins {
			r := image.Rect(i*size+offsetX, j*size+offsetY, (i+1)*size+offsetX, (j+1)*size+offsetY)
			miterLimit := float32(5)
			if i == 1 {
				miterLimit = 10
			}
			p.drawLine(g, target, r, cap, join, miterLimit)
		}
	}

	msg := fmt.Sprintf(`FPS: %0.2f, TPS: %0.2f
Press A to switch anti-aliasing.
Press C to switch to draw the center lines.`, ebiten.ActualFPS(), ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (p *Player) drawLine(g *MyGame, screen *ebiten.Image, region image.Rectangle, cap vector.LineCap, join vector.LineJoin, miterLimit float32) {
	c0x := float64(region.Min.X + region.Dx()/4)
	c0y := float64(region.Min.Y + region.Dy()/4)
	c1x := float64(region.Max.X - region.Dx()/4)
	c1y := float64(region.Max.Y - region.Dy()/4)
	r := float64(min(region.Dx(), region.Dy()) / 4)
	a0 := 2 * math.Pi * float64(g.count) / (16 * ebiten.DefaultTPS)
	a1 := 2 * math.Pi * float64(g.count) / (9 * ebiten.DefaultTPS)

	var path vector.Path
	sin0, cos0 := math.Sincos(a0)
	sin1, cos1 := math.Sincos(a1)
	path.MoveTo(float32(r*cos0+c0x), float32(r*sin0+c0y))
	path.LineTo(float32(-r*cos0+c0x), float32(-r*sin0+c0y))
	path.LineTo(float32(r*cos1+c1x), float32(r*sin1+c1y))
	path.LineTo(float32(-r*cos1+c1x), float32(-r*sin1+c1y))

	// Draw the main line in white.
	op := &vector.StrokeOptions{}
	op.LineCap = cap
	op.LineJoin = join
	op.MiterLimit = miterLimit
	op.Width = float32(r / 2)
	vs, is := path.AppendVerticesAndIndicesForStroke(g.vertices[:0], g.indices[:0], op)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 1
		vs[i].ColorG = 1
		vs[i].ColorB = 1
		vs[i].ColorA = 1
	}
	screen.DrawTriangles(vs, is, whiteSubImage, &ebiten.DrawTrianglesOptions{
		AntiAlias: g.aa,
	})

	// Draw the center line in red.
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
	}
}
