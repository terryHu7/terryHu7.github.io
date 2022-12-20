package main

import (
	"fmt"
	"golangdefault/defaults"
	"io/ioutil"
	"testing"
)

func TestMain(t *testing.T) {
	// recordSpec := []RecordSpec{}
	filepath := "/home/hutianyu/Record.json"
	// fullRecordPath := "/home/hutianyu/tmp1.mp4"
	// ejson := Ejson.NewEJsonInit(filepath)
	// ejson.Set([]byte(strconv.Quote(fullRecordPath)), "[0]", "Path")
	mcsbyte, _ := ioutil.ReadFile(filepath)
	// if err != nil {
	// 	return
	// }

	// err = json.Unmarshal(mcsbyte, &recordSpec)
	// if err != nil {
	// 	return
	// }

	tmp := []defaults.RecordSpec{}
	defaults.SetDefaults(&tmp, mcsbyte)
	fmt.Printf("%+v", tmp)

}
