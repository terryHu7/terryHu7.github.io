package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestFilepath(t *testing.T) {
	tpin := time.Now()
	file := "hh_电影模式_视频_2022-11-10-18-02-38.mp4"
	// dir := filepath.Dir(file)
	// basePath := filepath.Base(file)

	// files, _ := ioutil.ReadDir(dir)
	// for _, f := range files {
	// 	if f.IsDir() {
	// 		continue
	// 	}
	// 	if strings.Contains(f.Name(), strings.TrimSuffix(basePath, ".mp4")) && f.Name() != basePath {
	// 		fmt.Println("ok")
	// 	}

	// }
	dir := filepath.Dir(file)
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if strings.Contains(f.Name(), strings.TrimSuffix(file, ".mp4")+"_") {
			fmt.Println(dir + "/" + f.Name())
		}
	}
	time.Sleep(1 * time.Second)
	costtime := time.Since(tpin)
	fmt.Printf("SplitbyPipeline costtime: %v\n", costtime)

	fmt.Printf("max-size-time=%d", 10*int(math.Pow10(9)))
}
