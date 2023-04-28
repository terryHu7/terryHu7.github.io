package main

import (
	"io/ioutil"
	"os"
)

func main() {
	fileName := "./test.txt"

	content2 := `
	111	
	`

	nosync(fileName, content2)
}

func writeFile(fileName, content string) {
	err := ioutil.WriteFile(fileName, []byte(content), os.ModePerm)
	if err != nil {
		return
	}
}

func nosync(fileName, content string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	if _, err = file.Write([]byte(content)); err != nil {
		return
	}

	file.Sync()
}

// func main() {
// 	n := "~/iotoptrace.txt"
// 	ioutil.ReadFile(n)

// 	vmcmd := exec.Command("bash", "-c", fmt.Sprintf("vmtouch -e %s", n))
// 	var out bytes.Buffer
// 	vmcmd.Stdout = &out
// 	if vmerr := vmcmd.Run(); vmerr != nil {
// 		fmt.Printf("vmtouch::fail to free the cache of %s, err:%v", n, vmerr)
// 	} else {
// 		fmt.Printf("vmtouch::%s", out.String())
// 	}
// }
