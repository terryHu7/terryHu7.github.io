package main

import (
	"fmt"
	"myserver/database"
	"os/exec"
	"time"
)

func main() {
	database.Dbinit()
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C: //block
				{
					fmt.Println("tik tik...")
				}
			case <-done:
				{
					return
				}
			}
		}
	}()
	time.Sleep(1 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("tick finished")

	cmd := exec.Command("bash", "-c", "ls -rt")
	// cmd := exec.Command("echo", "Hi")
	out, _ := cmd.Output()
	fmt.Println(string(out))

}
