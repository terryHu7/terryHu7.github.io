package main

import (
	"fmt"
	"sort"
	"testing"
)

func Test452(t *testing.T) {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[j][0] < points[i][0]
	})
	segments := [][]int{}

	for i := 0; i < len(points); i++ {
		if len(segments) == 0 {
			segments = append(segments, points[i])
			continue
		}

		lastSegment := segments[len(segments)-1]
		ca, cb := points[i][0], points[i][1]
		la, lb := lastSegment[0], lastSegment[1]

		if cb < la || ca > lb {
			segments = append(segments, points[i])
			continue
		}
		var aa, bb int
		if ca <= la {
			aa = la
		} else {
			aa = ca
		}

		if cb >= lb {
			bb = lb
		} else {
			bb = cb
		}

		segments[len(segments)-1] = []int{aa, bb}
	}

	return len(segments)
}
