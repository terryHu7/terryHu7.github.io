package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/bwmarrin/snowflake"
)

func TestSnowFlake(t *testing.T) {
	n, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id", id)
	}
}
