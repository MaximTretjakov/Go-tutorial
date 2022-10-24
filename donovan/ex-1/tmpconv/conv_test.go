package tmpconv_test

import (
	"testing"

	"github.com/MaximTretjakov/Go-tutorial/donovan/ex-1/tmpconv"
)

func TestCToK(t *testing.T) {
	k := tmpconv.CToK(30)
	if k != -243.14999999999998 {
		t.Errorf("-273,15 + 30 = -243,15 got = %f", k)
	}
}

func TestKToC(t *testing.T) {
	c := tmpconv.KToC(-243.15)
	if c != 29.99999999999997 {
		t.Errorf("273,15 - 243,15 = 30 got = %f", c)
	}
}
