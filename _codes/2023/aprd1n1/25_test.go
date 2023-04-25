package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestTickets(t *testing.T) {
	tickets := [][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
	}
	fmt.Println(findItinerary(tickets))
}

type dstWithTag struct {
	dst     string
	visited bool
}

type dstsWithTag []*dstWithTag

func (ds dstsWithTag) Len() int {
	return len(ds)
}
func (ds dstsWithTag) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}
func (ds dstsWithTag) Less(i, j int) bool {
	return ds[i].dst < ds[j].dst
}

func findItinerary(tickets [][]string) []string {
	from_toMap := map[string]dstsWithTag{}

	for _, ticket := range tickets {
		if _, ok := from_toMap[ticket[0]]; !ok {
			from_toMap[ticket[0]] = dstsWithTag{}
		}

		from_toMap[ticket[0]] = append(from_toMap[ticket[0]], &dstWithTag{
			ticket[1],
			false,
		})
	}

	for _, from_to := range from_toMap {
		sort.Sort(from_to)
	}

	res := []string{"JFK"}

	var backtrack func() bool
	backtrack = func() bool {
		if len(res) == len(tickets)+1 {
			return true
		}

		for _, destination := range from_toMap[res[len(res)-1]] {
			if !destination.visited {
				res = append(res, destination.dst)
				destination.visited = true
				if backtrack() {
					return true
				}
				res = res[:len(res)-1]
				destination.visited = false
			}
		}
		return false
	}
	backtrack()
	return res
}
