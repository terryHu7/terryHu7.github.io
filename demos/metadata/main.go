package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	duration, _ := GetMP4Data("./test_metadata.mp4", 0, time.Time{})
	fmt.Println(duration)
}

type BoxHeader struct {
	Size       uint32
	FourccType [4]byte
	Size64     uint64
}

func GetCreatetimeFromFile(filename string) (createtime time.Time) {
	name := strings.TrimSuffix(path.Base(filename), path.Ext(path.Base(filename)))
	data := strings.Split(name, "_")
	datestrings := strings.Split(data[len(data)-1], "-")
	if len(datestrings) < 6 {
		datestrings = strings.Split(data[len(data)-2], "-")
	}
	date := datestrings[0] + "-" + datestrings[1] + "-" + datestrings[2] + " " + datestrings[3] + ":" + datestrings[4] + ":" + datestrings[5]
	createtime, _ = time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
	return
}

//MP4
func GetMP4Data(filename string, size int64, modtime time.Time) (duration int64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	if size == 0 {
		fileinfo, _ := file.Stat()
		size = fileinfo.Size()
		modtime = fileinfo.ModTime()
	}
	duration = 0
	if size == 0 {
		return
	}
	var info = make([]byte, 0x10)
	var boxHeader BoxHeader
	var moovOffset int64
	var moovSize uint32
	var offset int64 = 0
	var timeScale, Duration uint32
	var baseMediaDecodeTime, sampleCount, dataOffset uint32
	timeout := time.After(time.Second * 300)
	for {
		select {
		case <-timeout:
			createtime := GetCreatetimeFromFile(filename)
			duration = modtime.Unix() - createtime.Unix()
			return
		default:
		}
		_, err = file.ReadAt(info, offset)
		if err != nil {
			break
		}
		buf := bytes.NewBuffer(info)
		binary.Read(buf, binary.BigEndian, &boxHeader)
		fourccType := string(boxHeader.FourccType[:])
		if fourccType == "moov" {
			moovOffset = offset
			moovSize = boxHeader.Size
		}
		if fourccType == "moof" {
			moofinfo := make([]byte, boxHeader.Size)
			_, err = file.ReadAt(moofinfo, offset)
			if err != nil {
				break
			}
			boxtype := string(moofinfo[32+4 : 32+8])
			if boxtype == "tfhd" {
				tfdtboxoffset := binary.BigEndian.Uint32(moofinfo[32:32+4]) + 32
				default_sample_flags := binary.BigEndian.Uint32(moofinfo[tfdtboxoffset-4 : tfdtboxoffset])
				if default_sample_flags == 65728 {
					//tfdt
					versionbyte := moofinfo[tfdtboxoffset+8 : tfdtboxoffset+9]
					trunboxoffset := tfdtboxoffset
					if versionbyte[0] == 1 {
						baseMediaDecodeTime = uint32(binary.BigEndian.Uint64(moofinfo[tfdtboxoffset+12 : tfdtboxoffset+20]))
						trunboxoffset = tfdtboxoffset + 20
					} else {
						baseMediaDecodeTime = binary.BigEndian.Uint32(moofinfo[tfdtboxoffset+12 : tfdtboxoffset+16])
						trunboxoffset = tfdtboxoffset + 16
					}
					//trun
					sampleCount = binary.BigEndian.Uint32(moofinfo[trunboxoffset+12 : trunboxoffset+16])
					dataOffset = binary.BigEndian.Uint32(moofinfo[trunboxoffset+16 : trunboxoffset+20])
				}
			}
		}
		if fourccType == "mdat" {
			if boxHeader.Size == 1 {
				offset += int64(boxHeader.Size64)
				if offset == size {
					break
				} else if boxHeader.Size64 == 0 {

					return
				}
				continue
			}
		}
		if boxHeader.Size == 0 {

			return
		}
		offset += int64(boxHeader.Size)
		if offset == size || fourccType == "\x00\x00\x00\x00" {
			break
		}
	}
	moovStartBytes := make([]byte, moovSize)
	_, err = file.ReadAt(moovStartBytes, moovOffset)
	if err != nil {
		return
	}
	if len(moovStartBytes) != 0 {
		versionbyte := moovStartBytes[0x10:0x11]
		if versionbyte[0] == 1 {
			// modificationTimeOffset := 0x18
			timeScaleOffset := 0x20
			durationOffset := 0x24
			// modificationTime = uint32(binary.BigEndian.Uint64(moovStartBytes[modificationTimeOffset : modificationTimeOffset+8]))
			timeScale = binary.BigEndian.Uint32(moovStartBytes[timeScaleOffset : timeScaleOffset+4])
			Duration = uint32(binary.BigEndian.Uint64(moovStartBytes[durationOffset : durationOffset+8]))
		} else {
			// modificationTimeOffset := 0x18
			timeScaleOffset := 0x1C
			durationOffset := 0x20
			// modificationTime = binary.BigEndian.Uint32(moovStartBytes[modificationTimeOffset : modificationTimeOffset+4])
			timeScale = binary.BigEndian.Uint32(moovStartBytes[timeScaleOffset : timeScaleOffset+4])
			Duration = binary.BigEndian.Uint32(moovStartBytes[durationOffset : durationOffset+4])
		}
		//mvhd have no duration
		if Duration == 0 {
			boxsize := binary.BigEndian.Uint32(moovStartBytes[0x08 : 0x08+4])
			otherBoxOffset := 0x08 + int64(boxsize)
			var boxtype string
			for {
				boxtype = string(moovStartBytes[otherBoxOffset+4 : otherBoxOffset+8])
				if boxtype == "mvex" {
					versionbyte := moovStartBytes[otherBoxOffset+0x08+8 : otherBoxOffset+0x08+9]
					var fragmentDuration uint32
					if versionbyte[0] == 1 {
						fragmentDuration = uint32(binary.BigEndian.Uint64(moovStartBytes[otherBoxOffset+0x08+12 : otherBoxOffset+0x08+20]))
					} else {
						fragmentDuration = binary.BigEndian.Uint32(moovStartBytes[otherBoxOffset+0x08+12 : otherBoxOffset+0x08+16])
					}
					Duration = fragmentDuration
					break
				} else {
					boxsize = binary.BigEndian.Uint32(moovStartBytes[otherBoxOffset+0 : otherBoxOffset+4])
					otherBoxOffset += int64(boxsize)
					if otherBoxOffset >= int64(len(moovStartBytes)) {
						break
					}
				}
			}
		}
		//mvex have no fragment_duration
		if Duration == 0 {
			Duration = baseMediaDecodeTime + sampleCount*dataOffset
		}
		lengthOfTime := math.Floor(float64(Duration)/float64(timeScale) + 0.5)
		duration = int64(lengthOfTime)
		// modificationtime = utils.ChangeDate(modificationTime)
	}
	if duration == 0 {
	}
	return
}
