package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestReg(t *testing.T) {
	completePath := "/home/user/emedia/xxx/%d.mp4"
	dir := filepath.Dir(completePath)
	filename := filepath.Base(completePath)

	if strings.Contains(filename, "%") {
		r, _ := regexp.Compile(`%(0\d)?d`)
		if ok := r.MatchString(filename); ok {
			match := r.FindString(filename) // %d %02d %03d...
			if match != "" {
				digit, _ := strconv.Atoi(match[1 : len(match)-1])
				zeros := "0"
				for i := digit; i > 1; i-- {
					zeros += "0"
				}
				filename = r.ReplaceAllString(filename, zeros)
				completePath = dir + "/" + filename
			}
		}
	}
	fmt.Println(completePath)
}
