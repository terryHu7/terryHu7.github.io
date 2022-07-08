package main

import (
	"encoding/json"
	"fmt"
)

type BackgroundCfg struct {
	DeviceId   string `json:"DeviceId"`
	Background string `json:"Background"`
}

func main() {
	body := `{"Background":"/home/user/Desktop/wp_app_20220628_1805/wp_app_20220628_1805/wplivedevice/linux_arm64/v1.0.0_20220627_0/wplivedevice-linux-arm64-v1.0.0/./layout/default.png","DeviceId":"0800-0000"}`

	jsonbytes := []byte(body)

	backgroundcfg := BackgroundCfg{}
	_ = json.Unmarshal(jsonbytes, &backgroundcfg)
	fmt.Printf("%+v,%T,%T\n", backgroundcfg, backgroundcfg.Background, backgroundcfg.DeviceId)
}
