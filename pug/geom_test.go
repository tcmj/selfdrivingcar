package pug_test

import (
    "testing"
    "github.com/tcmj/selfdrivingcar/pug"
)

func TestPug(t *testing.T) {
    myresult := pug.Lerp32(2.0, 3.5, 2.0)
    if myresult != 3.2 {
        t.Fatalf("Wrong Result of pugutils.Lerp:  %f", myresult)
    }
}



func TestFloatingShizzle(t *testing.T) {
    var a float64 = 0.1
    var b float64 = 0.3
    var c float64 = a+b-0.4
    if a+b-0.4 != 0.0 {
        t.Fatalf("Wrong Result of pugutils.Lerp:  %.54f", c)
    }
}

