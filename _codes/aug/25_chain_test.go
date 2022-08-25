package aug

import (
	"fmt"
	"testing"
)

const (
	WP_PROB  = "wp problem"
	MA_PROB  = "ma problem"
	EMS_PROB = "ems problem"
)

type Group interface {
	Analyze(log string) string
}

type WP struct {
	Next Group
}

func (wp WP) Analyze(log string) (res string) {
	fmt.Println("WP Analyzing...")
	if log != WP_PROB {
		res = wp.Next.Analyze(log)
	} else {
		res = "WP would fix."
	}
	return res
}

type MA struct {
	Next Group
}

func (ma MA) Analyze(log string) (res string) {
	fmt.Println("MA Analyzing...")
	if log != MA_PROB {
		res = ma.Next.Analyze(log)
	} else {
		res = "MA would fix."
	}
	return res
}

type EMS struct {
	Next Group
}

func (ems EMS) Analyze(log string) (res string) {
	fmt.Println("EMS Analyzing...")
	if log != EMS_PROB {
		res = "Suggest to hold this issue."
	} else {
		res = "EMS would fix."
	}
	return res
}

func TestChain(t *testing.T) {
	team_chain := &WP{Next: &MA{Next: &EMS{}}}
	fmt.Println(team_chain.Analyze(EMS_PROB))
}
