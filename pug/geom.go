package pug

import (
	"errors"
	"fmt"
	"math"
	"runtime"

	"golang.org/x/exp/constraints"
)

type POS struct {
	X float64
	Y float64
}

func Lerp[T constraints.Float](valA, valB, t T) T {
	return valA + (valB-valA)*t
}

func Lerp32(valA, valB, t float32) float32 {
	return valA + (valB-valA)*t
}

// returns the intersection point if there is one (and also the offset/percentage)
func GetIntersection(A, B, C, D POS) (*POS, float64) {

	/*
			      Ix=Ax+(Bx-Ax)t=Cx+(Dx-Cx)u
			      Iy=Ay+(By-Ay)t=Cy+(Dy-Cy)u
		       ---------------------------------------------------------------------------------------------
			      Ax+(Bx-Ax)t=Cx+(Dx-Cx)u             |-Cx
			     (Ax-Cx)+(Bx-Ax)t=(Dx-Cx)u

			      Ay+(By-Ay)t=Cy+(Dy-Cy)u             |-Cy
			      (Ay-Cy)+(By-Ay)t=(Dy-Cy)u           |*(Dx-Cx)

			      (Dx-Cx)(Ay-Cy)+(Dx-Cx)(By-Ay)t  = (Dy-Cy)(Ax-Cx)+(Dy-Cy)(Bx-Ax)t   |-(Dy-Cy)(Ax-Cx)
			                                                                         |-(Dx-Cx)(By-Ay)t
			       (Dx-Cx)(Ay-Cy)-(Dy-Cy)(Ax-Cx)  = (Dy-Cy)(Bx-Ax)t-(Dx-Cx)(By-Ay)t  |Factor t
			   	(Dx-Cx)(Ay-Cy)-(Dy-Cy)(Ax-Cx)  = (Dy-Cy)(Bx-Ax)-(Dx-Cx)(By-Ay)t


			   	        (Dx-Cx)(Ay-Cy)-(Dy-Cy)(Ax-Cx)
			   	 t =   -----------------------------
			   	        (Dy-Cy)(Bx-Ax)-(Dx-Cx)(By-Ay)
	*/

	var ttop float64 = (D.X-C.X)*(A.Y-C.Y) - (D.Y-C.Y)*(A.X-C.X)
	var utop float64 = (C.Y-A.Y)*(A.X-B.X) - (C.X-A.X)*(A.Y-B.Y)
	var bott float64 = (D.Y-C.Y)*(B.X-A.X) - (D.X-C.X)*(B.Y-A.Y)

	if bott != 0.0 {
		var t float64 = ttop / bott
		var u float64 = utop / bott
		if t >= 0 && t <= 1 && u >= 0 && u < 1 { // check if values are intrapolated (not extrapolated)
			return &POS{X: Lerp(A.X, B.X, t), Y: Lerp(A.Y, B.Y, t)}, t
		}
	}
	return nil, 0
}

// var  t,_ = calculateOffsets(&A, &B, &C, &D)
// var  u ,_= calculateOffsets(&C, &D, &A, &B)
func calculateOffsets(A, B, C, D *POS) (float64, error) {
	var top = (D.Y-C.Y)*(A.X-C.X) - (D.X-C.X)*(A.Y-C.Y)
	var bottom = (D.X-C.X)*(B.Y-A.Y) - (D.Y-C.Y)*(B.X-A.X)
	if bottom == 0.0 {
		return math.NaN(), errors.New("Paralell")
	}
	return float64(top / bottom), nil
}



// Abs returns the absolute value of x.
func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}


func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc / 1024 / 1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc / 1024 / 1024)
	fmt.Printf("\tSys = %v MiB", m.Sys / 1024 / 1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
