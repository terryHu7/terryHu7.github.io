package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type Group struct {
	Name string `asdibasdhaod:"Name"`
}

func TestMain(t *testing.T) {
	GetCompileParamCfg()
}

func GetCompileParamCfg() (params *Group) {
	params = &Group{}
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return
	}
	json.Unmarshal(content, params)
	fmt.Println(params)
	return
}
