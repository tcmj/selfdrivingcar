package pugagui

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
)

func init() {
	whiteImage.Fill(color.White)
}

func NewLerper() *Player {
	asdf := &Player{}
	return asdf
}

func (p *Player) DrawSimple(g *MyGame, screen *ebiten.Image) {

	//fmt.Println("generatePlayerImage...")

	A := POS{200, 150}
	B := POS{150, 250}
	C := POS{50, 100}
	D := POS{250, 200}

	op := &vector.StrokeOptions{}
	op.Width = 2
	op.LineJoin = vector.LineJoinMiter
	
	var path vector.Path

	path.MoveTo(A.x, A.y)
	path.LineTo(B.x, B.y)
	path.MoveTo(C.x, C.y)
	path.LineTo(D.x, D.y)
	path.QuadTo(150, 57.5, 100, 45)
 

	vector.DrawFilledCircle( screen,A.x,A.y,8.0,color.White,g.aa)
	vector.DrawFilledCircle( screen,B.x,B.y,8.0,color.White,g.aa)
	vector.DrawFilledCircle( screen,C.x,C.y,8.0,color.White,g.aa)
	vector.DrawFilledCircle( screen,D.x,D.y,8.0,color.White,g.aa)


 
	ebitenutil.DebugPrintAt (screen, "A", int (A.x), int(A.y))
	ebitenutil.DebugPrintAt (screen, "B", int (B.x), int(B.y))
	ebitenutil.DebugPrintAt (screen, "C", int (C.x), int(C.y))
	ebitenutil.DebugPrintAt (screen, "D", int (D.x), int(D.y))
	

	offset := float32(12.0)
	path.MoveTo(0+offset, 50-offset)
	path.LineTo(25, 0+offset)
	path.LineTo(50-offset, 50-offset)
	path.LineTo(0+offset, 50-offset)
	path.Close()

	vs, is := path.AppendVerticesAndIndicesForStroke(nil, nil, op)

	genPlayerImage := ebiten.NewImage(150, 150)
	genPlayerImage.Fill(color.Transparent)

	genPlayerImage.Fill(color.Black)
	triOp := &ebiten.DrawTrianglesOptions{}
	triOp.AntiAlias = true

	//var redImage = ebiten.NewImage(3, 3)
	//var redSubImage = redImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

	screen.DrawTriangles(vs, is, whiteSubImage, &ebiten.DrawTrianglesOptions{
		AntiAlias: g.aa,
	})

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
	if g.showCenter {
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
