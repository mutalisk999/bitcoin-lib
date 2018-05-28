package utility

import "testing"

func TestAssert(t *testing.T) {
	Assert(1 == 1, "1==1")
	Assert(1 == 2, "1==2")
}
