package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/containerd/cgroups"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func main() {
	// fmt.Println("Open file...")
	// filepath := "/home/hutianyu/0823.txt"

	// for count, max := 0, 20; count < max; count++ {
	// 	fd, err := os.OpenFile(filepath, os.O_RDWR, 0)
	// 	if err != nil {
	// 		fmt.Printf("fail to open %s", filepath)
	// 	}
	// 	defer fd.Close()
	// 	content, _ := ioutil.ReadAll(fd)
	// 	fmt.Println(string(content))
	// 	fd.WriteString(fmt.Sprint(count))
	// 	time.Sleep(time.Second)
	// }
	ffmpegCmdStr := "-i /home/hutianyu/test.mp4 -c copy ./output.mp4 -y"

	pid, _ := executeFfmpeg(ffmpegCmdStr)
	if pid != 0 {
		addTasksIntoCgroups("ffmpeg_group", 150000, 1000000, pid)
	}

}

func executeFfmpeg(ffmpegCmdStr string) (pid int, err error) {
	cmdStrs := strings.Split(ffmpegCmdStr, " ")
	cmd := exec.Command("ffmpeg", cmdStrs...)
	err = cmd.Start()
	if err != nil {
		return 0, err
	}

	if cmd.Process != nil {
		return cmd.Process.Pid, nil
	} else {
		return 0, nil
	}
}

func addTasksIntoCgroups(cgPath string, quota int64, period uint64, pid int) {
	control, err := cgroups.New(cgroups.V1, cgroups.StaticPath(cgPath), &specs.LinuxResources{
		CPU: &specs.LinuxCPU{
			Quota:  &quota,
			Period: &period,
		},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	defer control.Delete()
	if err = control.AddTask(cgroups.Process{Pid: pid}, "cpu", "ffmpeg_group"); err != nil {
		log.Fatal(err)
		return
	}
	if tasks, err := control.Tasks(cgroups.Cpu, false); err == nil {
		log.Printf("Current tasks: %v", tasks)
	}
}
