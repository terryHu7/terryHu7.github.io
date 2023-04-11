package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	fileName := "test.txt"

	content1 := `
	111111111111111111111111111111111111111111111111111111
	111111111111111111111111111111111111111111111111111111
	111111111111111111111111111111111111111111111111111111
	111111111111111111111111111111111111111111111111111111
	111111111111111111111111111111111111111111111111111111
	111111111111111111111111111111111111111111111111111111
	111111111111111111111111111111111111111111111111111111
	111111111111111111111111111111111111111111111111111111
	111111111111111111111111111111111111111111111111111111
	111111111111111111111111111111111111111111111111111111
	`
	content2 := `
	222222222222222222222222222222222222222222222222222222
	222222222222222222222222222222222222222222222222222222
	222222222222222222222222222222222222222222222222222222
	222222222222222222222222222222222222222222222222222222
	222222222222222222222222222222222222222222222222222222
	222222222222222222222222222222222222222222222222222222
	222222222222222222222222222222222222222222222222222222
	222222222222222222222222222222222222222222222222222222
	222222222222222222222222222222222222222222222222222222
	222222222222222222222222222222222222222222222222222222
	`
	/*
		content1 := `
		{
			"Name": "metis_live_mcs_zlb",
			"Type": "MCS",
			"Description": "Pipeline task for metis live",
			"Version": "0.3",
			"CodecStreamSpecs": [
				{
					"Name": "Exp_front",
					"URI": "file:///home/user/Videos/eswin/trim0-118_30fps_1080p.mp4",
					"VideoCodec": "h.264"
				},
				{
					"Name": "Exp_top",
					"URI": "file:///home/user/Videos/eswin/trim1.24-119_30fps_1080p.mp4",
					"VideoCodec": "h.264"
				}
			],
			"ExperimentSpecs": [
				{
					"Name": "test",
					"DeskId": "123456",
					"TopViewStream": "Exp_top",
					"FrontViewStream": "Exp_front",
					"Code": "02010001",
					"Uuid": "7896661",
					"Step": "0201000101",
					"Width": 1920,
					"Height": 1080,
					"SnapshotPath": "/home/user/intelligent_experiment/7896661"
				}
			]
		}
		`

		content2 := `
		{
			"Name": "metis_live_mcs_zlb",
			"Type": "MCS",
			"Description": "Pipeline task for metis live",
			"Version": "0.3",
			"CodecStreamSpecs": [
				{
					"Name": "Exp_front",
					"URI": "file:///home/user/Videos/eswin/������̼���ԭ����̽��-����_trim0-118_30fps_1080p.mp4",
					"VideoCodec": "h.264"
				},
				{
					"Name": "Exp_top",
					"URI": "file:///home/user/Videos/eswin/������̼���ԭ����̽��-����_trim1.24-119_30fps_1080p.mp4",
					"VideoCodec": "h.264"
				}
			],
			"ExperimentSpecs": [
				{
					"Name": "test",
					"DeskId": "123456",
					"TopViewStream": "Exp_top",
					"FrontViewStream": "Exp_front",
					"Code": "02010001",
					"Uuid": "7896661",
					"Step": "0201000101",
					"Width": 1920,
					"Height": 1080,
					"SnapshotPath": "/home/user/intelligent_experiment/7896661"
				}
			],
			"RenderSpecs": [
				{
					"Name": "Render",
					"DeviceId": "0800-0000",
					"CompositionSpec": [
						{
							"Geometry": [
								0,
								0,
								960,
								540
							],
							"SourceName": "test::Front"
						},
						{
							"Geometry": [
								960,
								0,
								960,
								540
							],
							"SourceName": "test::Top"
						},
						{
							"Geometry": [
								0,
								540,
								960,
								540
							],
							"SourceName": "Exp_front"
						},
						{
							"Geometry": [
								960,
								540,
								960,
								540
							],
							"SourceName": "Exp_top"
						}
					]
				}
			]
		}
		`
	*/
	go writeFile(fileName, content1)
	go writeFile(fileName, content2)
	time.Sleep(time.Second / 2)
}

func writeFile(fileName, content string) {
	fmt.Println(len(content))
	err := ioutil.WriteFile(fileName, []byte(content), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
}
