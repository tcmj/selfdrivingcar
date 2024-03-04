package pugutils_test

import (
	"testing"
	"github.com/tcmj/selfdrivingcar/pugutils"
)

func TestPug(t *testing.T) {
	myresult := pugutils.Lerp(2.0, 3.5, 2.0)
	if myresult != 3.2 {
		t.Fatalf("Wrong Result of pugutils.Lerp:  %f", myresult)
	}
}
