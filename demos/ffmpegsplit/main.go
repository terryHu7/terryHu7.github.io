package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	_ = SplitMedia("./test.fmp4", "Size", 10485760*2, "mp4") //52428800
}

/*
func Shellcmd() error {
	cmd := exec.Command("ls")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}
*/

func SplitMedia(filePath string, splitMethod string, splitUnit int, format string) []string {
	suffix := filepath.Ext(filePath)
	newSuffix := suffix
	if suffix == ".fmp4" {
		newSuffix = ".mp4"
	}
	dstPath := strings.TrimSuffix(filePath, suffix) + "_%02d" + newSuffix

	fmt.Printf("Split Media file %s to %s\n", filePath, dstPath)

	resFiles := []string{}
	if splitMethod == "Time" && splitUnit != 0 {
		fmt.Printf("Split Media file by Time.")

		ffmpegCmdStr := fmt.Sprintf("-i %s -c copy -f segment -segment_time %d -reset_timestamps 1 %s", filePath, splitUnit, dstPath)
		if format == "mp4" {
			ffmpegCmdStr = fmt.Sprintf("-i %s -c copy -f segment -segment_time %d -segment_format_options movflags=+faststart -reset_timestamps 1 %s", filePath, splitUnit, dstPath)
		}
		err := executeFfmpeg(ffmpegCmdStr)
		if err != nil {
			return []string{}
		}

		// store splited files
		dir := filepath.Dir(filePath)
		basePath := filepath.Base(filePath)

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Printf("Read dir fail: %v", err)
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}
			if strings.Contains(f.Name(), strings.TrimSuffix(basePath, suffix)) && f.Name() != basePath {
				resFiles = append(resFiles, dir+"/"+f.Name())
			}
		}
	} else if splitMethod == "Size" && splitUnit != 0 {
		fmt.Println("Split Media file by Size.")

		// Duration of the source media
		duration, err := getDuration(filePath)
		if err != nil {
			return resFiles
		}
		if duration == 0 {
			fmt.Printf("No need to split media file %s (duration=0)", filePath)
			return resFiles
		}
		// Duration that has been encoded so far
		curDuration := 0
		// Time nodes to split video
		segmentTimes := []string{}
		// Filename of the source media (without extension)
		basePath := strings.TrimSuffix(filePath, suffix)
		// Number of the current video part
		outputIndex := 0
		// Filename of the next video part
		nextFilename := fmt.Sprintf("%s_%02d%s", basePath, outputIndex, newSuffix)
		for curDuration < duration {

			if curDuration != 0 {
				segmentTimes = append(segmentTimes, strconv.FormatInt(int64(curDuration), 10))
			}
			// split the next part
			ffmpegCmdStr := fmt.Sprintf("-i %s -ss %d -fs %d -c copy %s", filePath, curDuration, splitUnit, nextFilename)
			if format == "mp4" {
				ffmpegCmdStr = fmt.Sprintf("-i %s -ss %d -fs %d -c copy -movflags +faststart %s", filePath, curDuration, splitUnit, nextFilename)
			}
			err := executeFfmpeg(ffmpegCmdStr)
			if err != nil {
				return resFiles
			}

			// duration of the new part
			newDuration, err := getDuration(nextFilename)
			if err != nil {
				return resFiles
			}
			if newDuration == 0 {
				if outputIndex == 0 {
					return []string{filePath}
				} else {
					fmt.Printf("Remove media file without content: %s", nextFilename)
					err = os.Remove(nextFilename)
					if err != nil {
						fmt.Printf("Delete file failed %v " + err.Error())
					}
					break
				}
			}

			// store splited files
			resFiles = append(resFiles, nextFilename)

			// Total duration splited so far
			curDuration = curDuration + newDuration
			outputIndex = outputIndex + 1

			nextFilename = fmt.Sprintf("%s_%02d%s", basePath, outputIndex, newSuffix)
		}

		if len(segmentTimes) != 0 {
			ffmpegCmdStr := fmt.Sprintf("-i %s -f segment -segment_times %s -c copy -reset_timestamps 1 %s -y", filePath, strings.Join(segmentTimes, ","), dstPath)
			if format == "mp4" {
				ffmpegCmdStr = fmt.Sprintf("-i %s -f segment -segment_times %s -c copy -segment_format_options movflags=+faststart -reset_timestamps 1 %s -y", filePath, strings.Join(segmentTimes, ","), dstPath)
			}
			err = executeFfmpeg(ffmpegCmdStr)
			if err != nil {
				return []string{}
			}
		}
	} else {
		fmt.Printf("Unsupported Split method. SplitMethod: %s, SplitUnit: %d", splitMethod, splitUnit)
		return resFiles
	}
	fmt.Printf("Split finished: %s", filePath)

	// fmt.Printf("Remove unsplited file: %s", filePath)
	// err := os.Remove(filePath)
	// if err != nil {
	// 	fmt.Printf("Delete file failed %v" + err.Error() + filePath)
	// }
	return resFiles
}

////////////////////////////////////////////////////////
func executeFfmpeg(ffmpegCmdStr string) error {
	cmdStrs := strings.Split(ffmpegCmdStr, " ")
	cmd := exec.Command("ffmpeg", cmdStrs...)
	err := cmd.Start()
	if err != nil {
		return err
	}
	err = WaitTimeout(cmd, 1000*time.Second)
	if err != nil {
		return err
	}
	return nil
}

const KillGrace = 5 * time.Second

func WaitTimeout(c *exec.Cmd, timeout time.Duration) error {
	var kill *time.Timer
	term := time.AfterFunc(timeout, func() {
		err := c.Process.Signal(syscall.SIGTERM)
		if err != nil {
			//log.Printf("E! [agent] Error terminating process: %s", err)
			return
		}

		kill = time.AfterFunc(KillGrace, func() {
			err := c.Process.Kill()
			if err != nil {
				//log.Printf("E! [agent] Error killing process: %s", err)
				return
			}
		})
	})

	err := c.Wait()

	// Shutdown all timers
	if kill != nil {
		kill.Stop()
	}
	termSent := !term.Stop()

	// If the process exited without error treat it as success.  This allows a
	// process to do a clean shutdown on signal.
	if err == nil {
		return nil
	}

	// If SIGTERM was sent then treat any process error as a timeout.
	if termSent {
		//return TimeoutErr
	}

	// Otherwise there was an error unrelated to termination.
	return err
}

func getDuration(path string) (int, error) {
	durationcmd := exec.Command("bash", "-c", fmt.Sprintf("ffprobe -i %s -show_entries format=duration -v quiet -of default=noprint_wrappers=1:nokey=1", path))
	output, err := durationcmd.CombinedOutput()
	if err != nil {
		return 0, err
	}
	durationStr := string(output)
	fmt.Printf("Duration String of file: %s, duration: %s", path, durationStr)
	duration := strings.Split(durationStr, ".")[0]
	value, _ := strconv.ParseInt(duration, 10, 64)
	return int(value), nil
}
