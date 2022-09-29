package main

import "testing"

func TestMd2Html(t *testing.T) {
	md2html("./tmp.txt", "./outtmp.txt")
}
