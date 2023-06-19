package main

import (
	"fmt"
	"math"
	"os"
	"testing"
	"time"
	"unsafe"
)

func TestMKV(t *testing.T) {
	fd, _ := os.Stat("./wrong.mkv")
	size := fd.Size()
	duration, _ := GetMkvData("./wrong.mkv", size)
	fmt.Println(duration)
}

//MKV
func GetMkvData(filename string, size int64) (duration int64, err error) {

	if EmblLenMapTable[0] == 9 {
		return
	}
	EmblLenMapTable[0] = 9
	for ulLoop := 1; ulLoop < 255; ulLoop++ {
		for index := 1; index <= 8; index++ {
			if ulLoop>>index == 0 {
				EmblLenMapTable[ulLoop] = 8 - index + 1
				break
			}
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	if size == 0 {
		fileinfo, _ := file.Stat()
		size = fileinfo.Size()
	}
	duration = 0
	if size == 0 {
		return
	}

	content := make([]byte, 1024*1024)

	totallen, err := file.Read(content)
	if err != nil {
		return
	}

	analylen := 0
	ebmlmng := &EbmlManager{}
	ebmlmng.OperateFlag = ebmlmng.OperateFlag | EbmlDataGetDuration
	ebmlmng.OperateFlag = ebmlmng.OperateFlag | EbmlDataGetModTime
	ebmlmng.OperateFlag = ebmlmng.OperateFlag | EbmlDataGetTimecodeScale

	for {
		ebmlData := content[analylen:totallen]
		ebmlctl := getEbmlInfo(&ebmlData, totallen-analylen)
		analylen = analylen + ebmlctl.DataLen + ebmlctl.ContrlLen
		if analylen > totallen {
			ebmlctl.DataLen = ebmlctl.DataLen - (analylen - totallen)
			analylen = totallen
		}

		segdata := content[analylen-ebmlctl.DataLen : analylen]
		switch ebmlctl.ElementId {
		case 0x18538067:
			ebmlmng.EbmlSegmentProc(&segdata, ebmlctl.DataLen)
		}

		if ebmlmng.OperateFlag == 0 {
			break
		}

		if analylen >= totallen {
			break
		}
	}

	if ebmlmng.TimecodeScale == 0 {
		ebmlmng.TimecodeScale = 1
	}

	var createtime time.Time
	if ebmlmng.OperateFlag&EbmlDataGetModTime == 0 {
		// lastmodtime := ebmlmng.CreateTime + int64(ebmlmng.Duration*float64(ebmlmng.TimecodeScale))
		// modtime = time.Unix(lastmodtime/int64(time.Second), lastmodtime%int64(time.Second)).Format("2006-01-02 15:04:05")
		createtime = time.Unix(ebmlmng.CreateTime/int64(time.Second), ebmlmng.CreateTime%int64(time.Second))
	}

	if ebmlmng.OperateFlag&EbmlDataGetDuration == 0 {
		durSec := int64(math.Floor(float64(ebmlmng.Duration)*float64(ebmlmng.TimecodeScale)/float64(time.Second) + 0.5))
		duration = durSec
	}

	if duration == 0 {
		now := time.Now().Unix()
		duration = now - createtime.Unix()
	}

	return
}

type EbmlDataCtlInfo struct {
	ElementId int64
	DataLen   int
	ContrlLen int
}

var EmblLenMapTable [255]int

type EbmlManager struct {
	FileName      string
	Duration      float64
	CreateTime    int64
	TimecodeScale uint64
	OperateFlag   uint64
	BufLeftLen    uint64
}

const (
	EbmlDataGetDuration      = uint64(1)
	EbmlDataGetModTime       = uint64(1 << 1)
	EbmlDataGetTimecodeScale = uint64(1 << 2)
)

func (ebmlmng *EbmlManager) EbmlSegmentProc(data *[]byte, totallen int) {

	analylen := 0
	content := (*data)[0:totallen]
	for {
		ebmlData := content[analylen:totallen]
		ebmlctl := getEbmlInfo(&ebmlData, totallen-analylen)
		analylen = analylen + ebmlctl.DataLen + ebmlctl.ContrlLen
		if analylen > totallen {
			ebmlctl.DataLen = ebmlctl.DataLen - (analylen - totallen)
			analylen = totallen
		}

		segdata := content[analylen-ebmlctl.DataLen : analylen]
		switch ebmlctl.ElementId {
		case 0x1549A966:
			ebmlmng.EbmlSegmentInformationProc(&segdata, ebmlctl.DataLen)
		case 0x1F43B675:
			//ebmlmng.EbmlClusterProc(&segdata, ebmlctl.DataLen)
		}

		if ebmlmng.OperateFlag == 0 {
			break
		}

		if analylen >= totallen {
			break
		}
	}

	return
}

func (ebmlmng *EbmlManager) EbmlSegmentInformationProc(data *[]byte, totallen int) {

	analylen := 0
	content := (*data)[0:totallen]
	for {
		ebmlData := content[analylen:totallen]
		ebmlctl := getEbmlInfo(&ebmlData, totallen-analylen)
		analylen = analylen + ebmlctl.DataLen + ebmlctl.ContrlLen
		if analylen > totallen {
			ebmlctl.DataLen = ebmlctl.DataLen - (analylen - totallen)
			analylen = totallen
		}
		segdata := content[analylen-int(ebmlctl.DataLen) : analylen]
		switch ebmlctl.ElementId {
		case 0x2AD7B1:
			ebmlmng.TimecodeScale = getElmbDataInt(&segdata, ebmlctl.DataLen)
			ebmlmng.OperateFlag = ebmlmng.OperateFlag & (^EbmlDataGetTimecodeScale)
		case 0x2AD7B2:
		case 0x4489:
			duration := getElmbDataInt(&segdata, ebmlctl.DataLen)
			ebmlmng.Duration = math.Float64frombits(duration)
			ebmlmng.OperateFlag = ebmlmng.OperateFlag & (^EbmlDataGetDuration)
		case 0x4461:
			timeMod := getElmbDataInt(&segdata, ebmlctl.DataLen)
			//go default,nust be 2006-01-02 15:04:05
			timeLayout := "2006-01-02 15:04:05"

			//UTC start time
			timeUtcStr := "2001-01-01 00:00:00"
			timeUtc, _ := time.Parse(timeLayout, timeUtcStr)
			ebmlmng.CreateTime = timeUtc.UnixNano() + int64(timeMod)
			ebmlmng.OperateFlag = ebmlmng.OperateFlag & (^EbmlDataGetModTime)
		case 0x4d80:

		}

		if ebmlmng.OperateFlag == 0 {
			break
		}

		if analylen >= totallen {
			break
		}
	}

	return
}

func GetNetWorkBigEndianFlag() bool {
	data := uint16(1)
	dataByte := (*[2]byte)(unsafe.Pointer(&data))

	return (*dataByte)[0] != 1
}

func getEbmlElementId(data *[]byte) (elementid int64, idLen int) {
	idLen = EmblLenMapTable[int((*data)[0])]
	if idLen == 9 {
		idLen = 8
	}

	//mkv big endian save
	bigflag := GetNetWorkBigEndianFlag()

	elementbyte := (*[8]byte)(unsafe.Pointer(&elementid))
	for ulLoop := 0; ulLoop < idLen; ulLoop++ {
		if bigflag {
			(*elementbyte)[8-idLen+ulLoop] = (*data)[ulLoop]
		} else {
			(*elementbyte)[7-(8-idLen)-ulLoop] = (*data)[ulLoop]
		}
	}

	return
}

func getEbmlDataSize(data *[]byte) (datasize int64, sizeLen int) {
	sizeLen = EmblLenMapTable[int((*data)[0])]
	if sizeLen == 9 {
		sizeLen = 8
	}

	//mkv big endian save
	bigflag := GetNetWorkBigEndianFlag()

	(*data)[0] = (0xFF >> sizeLen) & ((*data)[0])

	dataseizebyte := (*[8]byte)(unsafe.Pointer(&datasize))
	for ulLoop := 0; ulLoop < sizeLen; ulLoop++ {
		if bigflag {
			(*dataseizebyte)[8-sizeLen+ulLoop] = (*data)[ulLoop]
		} else {
			(*dataseizebyte)[7-(8-sizeLen)-ulLoop] = (*data)[ulLoop]
		}
	}

	return
}

func getEbmlInfo(content *[]byte, totallen int) (ebmlctl *EbmlDataCtlInfo) {

	ebmlctl = &EbmlDataCtlInfo{}
	leftLen := totallen
	if leftLen > 10 {
		leftLen = 10
	}

	iddata := (*content)[0:leftLen]
	id, idlen := getEbmlElementId(&iddata)
	ebmlctl.ElementId = id

	leftLen = totallen - idlen
	if leftLen > 10 {
		leftLen = 10
	}

	data := (*content)[idlen : idlen+leftLen]
	datasize, datasizelen := getEbmlDataSize(&data)
	ebmlctl.DataLen = int(datasize)
	ebmlctl.ContrlLen = idlen + datasizelen
	return
}

func getElmbDataInt(data *[]byte, datalen int) (dataint uint64) {

	//mkv big endian save
	bigflag := GetNetWorkBigEndianFlag()

	elementbyte := (*[8]byte)(unsafe.Pointer(&dataint))
	for ulLoop := 0; ulLoop < datalen; ulLoop++ {
		if bigflag {
			(*elementbyte)[8-datalen+ulLoop] = (*data)[ulLoop]
		} else {
			(*elementbyte)[7-(8-datalen)-ulLoop] = (*data)[ulLoop]
		}
	}

	return
}
