package main

import "testing"

func TestFFmpegsplit(t *testing.T) {
	_ = SplitMedia("./test.fmp4", "Size", 10485760*2, "mp4")
}
