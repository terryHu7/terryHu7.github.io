package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ParamConfig struct {
	RecordWithCodec          bool   `json:"RecordWithCodec"`
	CloseupWithTee           bool   `json:"CloseupWithTee"`
	RtcApiControl            bool   `json:"RtcApiControl"`
	AudioMinThresholdTime    string `json:"AudioMinThresholdTime"`
	Streamsync               string `json:"Streamsync"`
	TransStreamsync          string `json:"TransStreamsync"`
	PreviewStreamsync        string `json:"PreviewStreamsync"`
	ComposePreviewStreamsync string `json:"ComposePreviewStreamsync"`
	ComposeStreamsync        string `json:"ComposeStreamsync"`
	DirectorStreamsync       string `json:"DirectorStreamsync"`
	CloseupLatency           string `json:"CloseupLatency"`
	PipeMutiSend             bool   `json:"PipeMutiSend"`
	EnableAudioJitter        bool   `json:"EnableAudioJitter"`
	HoldOnASecond            int    `json:"HoldOnASecond"`
}

func main() {
	paramConfig := &ParamConfig{}
	if content, err := ioutil.ReadFile("/home/hutianyu/mediaconfig.json"); err == nil {
		fmt.Println(string(content))
		json.Unmarshal(content, paramConfig)
	}
	fmt.Printf("paramcfg:%+v", paramConfig)
}
