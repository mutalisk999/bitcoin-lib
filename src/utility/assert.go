package utility

func Assert(cond bool, errMsg string) {
	if !cond {
		panic("Assert: " + errMsg)
	}
}
