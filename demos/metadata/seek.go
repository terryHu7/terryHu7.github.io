package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func seek(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// ��λ��moov box��λ��
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}
	var size uint32
	var boxType uint32
	for {
		err = binary.Read(file, binary.BigEndian, &size)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		err = binary.Read(file, binary.BigEndian, &boxType)
		if err != nil {
			panic(err)
		}
		// �ж��Ƿ�Ϊmoov box
		if boxType == 0x6d6f6f76 {
			break
		}
		// ��������box
		_, err = file.Seek(int64(size-8), io.SeekCurrent)
		if err != nil {
			panic(err)
		}
	}

	// ��ȡmvhd box������
	_, err = file.Seek(12, io.SeekCurrent) // �����汾�źͱ�־λ
	if err != nil {
		panic(err)
	}
	var creationTime uint64
	var modificationTime uint64
	var timeScale uint32
	var duration uint64
	err = binary.Read(file, binary.BigEndian, &creationTime)
	if err != nil {
		panic(err)
	}
	err = binary.Read(file, binary.BigEndian, &modificationTime)
	if err != nil {
		panic(err)
	}
	err = binary.Read(file, binary.BigEndian, &timeScale)
	if err != nil {
		panic(err)
	}
	err = binary.Read(file, binary.BigEndian, &duration)
	if err != nil {
		panic(err)
	}

	// �������ʱ��
	durationInSeconds := float64(duration) / float64(timeScale)

	fmt.Printf("Duration: %f seconds\n", durationInSeconds)
}
