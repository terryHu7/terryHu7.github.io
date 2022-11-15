package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

var cpulimitpid int

func main() {
	defer func() {
		killcpulimit := exec.Command("killall", "cpulimit")
		if err := killcpulimit.Run(); err != nil {
			fmt.Printf("killall cpulimit err: %v\n", err.Error())
		}
	}()
	fmt.Println("mark 1-----------------------")
	cpulimitcmd := exec.Command("/usr/bin/eswin/cpulimit", "-e", "/usr/bin/ffmpeg", "-l", "500", "-b")
	fmt.Println("cpulimit.Run starting--------------------------------")
	if cpuerr := cpulimitcmd.Run(); cpuerr != nil {
		fmt.Println(cpuerr.Error())
	}
	if cpulimitcmd.Process != nil && cpulimitcmd.Process.Pid != 0 {
		cpulimitpid = cpulimitcmd.Process.Pid
		fmt.Printf("cpulimit pid : %v\n", cpulimitpid)
	}

	fmt.Println("mark 2-----------------------")
	ffmpegCmdStr := fmt.Sprintf("-i %s %s -y", "./test.fmp4", "./tmp.mp4")
	executeFfmpeg(ffmpegCmdStr)
	fmt.Println("mark 5-----------------------")
	time.Sleep(20 * time.Second)
	fmt.Println("mark 6-----------------------")
	executeFfmpeg(ffmpegCmdStr)
	fmt.Println("mark 7-----------------------")

}

func executeFfmpeg(ffmpegCmdStr string) error {
	fmt.Printf("Execute: ffmpeg %s\n", ffmpegCmdStr)
	cmdStrs := strings.Split(ffmpegCmdStr, " ")
	fmt.Println("--------------------Calculating time...")
	tpin := time.Now()
	cmd := exec.Command("ffmpeg", cmdStrs...)
	err := cmd.Start()
	if err != nil {
		panic("ffmpeg")
	}

	if cmd.Process != nil && cmd.Process.Pid != 0 {
		toptrace := exec.Command("bash", "/home/hutianyu/toptrace.sh", fmt.Sprint(cpulimitpid), fmt.Sprint(cmd.Process.Pid))
		fmt.Println("toptrace.Run starting--------------------------------")
		toptrace.Start()
		fmt.Println("mark 3-----------------------")
	}

	err = WaitTimeout(cmd, 1000*time.Second) //block until timeout
	if err != nil {
		panic("ffmpeg time out")
	}
	fmt.Println("mark 4-----------------------")
	costtime := time.Since(tpin)
	fmt.Printf("--------------------ffmpeg %s costtime: %v\n", ffmpegCmdStr, costtime)
	return nil
}
