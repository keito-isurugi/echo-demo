package mono

import (
	"testing"
)

//test demo
func Abs(i int) int{
	return 1
}
func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
			t.Errorf("Abs(-1) = %d; want 1", got)
	}
}