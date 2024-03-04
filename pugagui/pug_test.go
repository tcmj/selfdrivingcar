package pugagui_test

import (
	"testing"

	"github.com/tcmj/selfdrivingcar/pugagui"
)

func TestPug(t *testing.T) {

	if pugagui.BestMasco() != "Pug" {
		t.Fatal("Wrong Pug")
	}
}
