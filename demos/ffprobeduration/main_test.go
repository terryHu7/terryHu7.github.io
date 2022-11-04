package main

import (
	"testing"
)

func TestFFprobeDuration(t *testing.T) {
	for i := 0; i < 10000; i++ {
		path := "YuWen_20221101_152236_3_0.fmp4"
		_, _ = getDurationV2(path)
	}
}
