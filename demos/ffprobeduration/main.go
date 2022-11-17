package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	shell "github.com/kballard/go-shellquote"
	//"github.com/kr/pty"
	"github.com/creack/pty"
)

var pass int

func main() {
	for i := 0; i < 10; i++ {
		path := "YuWen_20221101_152236_3_0.fmp4"
		_, _ = getDurationV2(path)
	}
}

func getDurationV2(path string) (int, error) {
	cmdStrs := strings.Split(fmt.Sprintf("-i %s -show_entries format=duration -v quiet -of default=noprint_wrappers=1:nokey=1", path), " ")
	ffprobecmd := exec.Command("ffprobe", cmdStrs...)
	var out bytes.Buffer
	ffprobecmd.Stdout = &out
	err := ffprobecmd.Run()
	if err != nil {
		return 0, err
	}
	fmt.Printf("Duration String of file: %s, duration: %s", path, out.String())
	duration := strings.Split(out.String(), ".")[0]
	value, _ := strconv.ParseInt(duration, 10, 64)
	return int(value), nil
}

func getDuration(path string) (int, error) {
	durationStr, _ := RunCmd(fmt.Sprintf("ffprobe -i %s -show_entries format=duration -v quiet -of default=noprint_wrappers=1:nokey=1", path))
	fmt.Printf("Duration String of file: %s, duration: %s", path, durationStr)
	// pass++
	// fmt.Println(pass)
	duration := strings.Split(durationStr, ".")[0]
	value, _ := strconv.ParseInt(duration, 10, 64)
	return int(value), nil
}

var (
	ErrEmptySearch = errors.New("empty search string")
)

type ExpectSubprocess struct {
	Cmd          *exec.Cmd
	buf          *buffer
	outputBuffer []byte
}
type buffer struct {
	f       *os.File
	b       bytes.Buffer
	collect bool

	collection bytes.Buffer
}

func (buf *buffer) Read(chunk []byte) (int, error) {
	nread := 0
	if buf.b.Len() > 0 {
		n, err := buf.b.Read(chunk)
		if err != nil {
			return n, err
		}
		if n == len(chunk) {
			return n, nil
		}
		nread = n
	}
	fn, err := buf.f.Read(chunk[nread:])
	return fn + nread, err
}

func (buf *buffer) ReadRune() (r rune, size int, err error) {
	l := buf.b.Len()

	chunk := make([]byte, utf8.UTFMax)
	if l > 0 {
		n, err := buf.b.Read(chunk)
		if err != nil {
			return 0, 0, err
		}
		if utf8.FullRune(chunk[:n]) {
			r, rL := utf8.DecodeRune(chunk)
			if n > rL {
				buf.PutBack(chunk[rL:n])
			}
			if buf.collect {
				buf.collection.WriteRune(r)
			}
			return r, rL, nil
		}
	}
	// else add bytes from the file, then try that
	for l < utf8.UTFMax {
		fn, err := buf.f.Read(chunk[l : l+1])
		if err != nil {
			return 0, 0, err
		}
		l = l + fn

		if utf8.FullRune(chunk[:l]) {
			r, rL := utf8.DecodeRune(chunk)
			if buf.collect {
				buf.collection.WriteRune(r)
			}
			return r, rL, nil
		}
	}
	return 0, 0, errors.New("File is not a valid UTF=8 encoding")
}

func (buf *buffer) PutBack(chunk []byte) {
	if len(chunk) == 0 {
		return
	}
	if buf.b.Len() == 0 {
		buf.b.Write(chunk)
		return
	}
	d := make([]byte, 0, len(chunk)+buf.b.Len())
	d = append(d, chunk...)
	d = append(d, buf.b.Bytes()...)
	buf.b.Reset()
	buf.b.Write(d)
}

func (expect *ExpectSubprocess) Start() error {
	_, err := _start(expect)
	return err
}
func Spawn(command string) (*ExpectSubprocess, error) {
	expect, err := _spawn(command)
	if err != nil {
		panic("serr")
	}
	res, serr := _start(expect)
	if serr != nil {
		panic("serr")
	}
	return res, serr
}

func (expect *ExpectSubprocess) Close() error {
	if err := expect.Cmd.Process.Kill(); err != nil {
		return err
	}
	if err := expect.buf.f.Close(); err != nil {
		return err
	}
	return nil
}

func buildKMPTable(searchString string) []int {
	pos := 2
	cnd := 0
	length := len(searchString)

	var table []int
	if length < 2 {
		length = 2
	}

	table = make([]int, length)
	table[0] = -1
	table[1] = 0

	for pos < len(searchString) {
		if searchString[pos-1] == searchString[cnd] {
			cnd += 1
			table[pos] = cnd
			pos += 1
		} else if cnd > 0 {
			cnd = table[cnd]
		} else {
			table[pos] = 0
			pos += 1
		}
	}
	return table
}

func (expect *ExpectSubprocess) Expect(searchString string) (e error) {
	target := len(searchString)
	if target < 1 {
		return ErrEmptySearch
	}
	chunk := make([]byte, target*2)
	if expect.outputBuffer != nil {
		expect.outputBuffer = expect.outputBuffer[0:]
	}
	m := 0
	i := 0
	// Build KMP Table
	table := buildKMPTable(searchString)

	for {
		n, err := expect.buf.Read(chunk)
		if n == 0 && err != nil {
			return err
		}
		if expect.outputBuffer != nil {
			expect.outputBuffer = append(expect.outputBuffer, chunk[:n]...)
		}
		offset := m + i
		for m+i-offset < n {
			if searchString[i] == chunk[m+i-offset] {
				i += 1
				if i == target {
					unreadIndex := m + i - offset
					if len(chunk) > unreadIndex {
						expect.buf.PutBack(chunk[unreadIndex:n])
					}
					return nil
				}
			} else {
				m += i - table[i]
				if table[i] > -1 {
					i = table[i]
				} else {
					i = 0
				}
			}
		}
	}
}

func (expect *ExpectSubprocess) SendLine(command string) error {
	_, err := io.WriteString(expect.buf.f, command+"\r\n")
	return err
}

func (expect *ExpectSubprocess) InteractWithReturn() (outPut bytes.Buffer) {
	var b2 bytes.Buffer
	defer expect.Cmd.Wait()
	//io.Copy(os.Stdout, &expect.buf.b)
	//go io.Copy(os.Stdout, expect.buf.f)
	//go io.Copy(expect.buf.f, os.Stdin)
	io.Copy(&b2, &expect.buf.b)
	go io.Copy(&b2, expect.buf.f)
	go io.Copy(expect.buf.f, os.Stdin)
	//time.Sleep(10*time.Second)
	fmt.Println(b2)
	return b2
}

func (expect *ExpectSubprocess) Wait() error {
	return expect.Cmd.Wait()
}

func _start(expect *ExpectSubprocess) (*ExpectSubprocess, error) {
	f, err := pty.Start(expect.Cmd)
	if err != nil {
		panic("pty.Start err")
	}
	expect.buf.f = f
	return expect, nil
}

func _spawn(command string) (*ExpectSubprocess, error) {
	wrapper := new(ExpectSubprocess)

	wrapper.outputBuffer = nil

	splitArgs, err := shell.Split(command)
	if err != nil {
		panic("fail to Split")
	}
	numArguments := len(splitArgs) - 1
	if numArguments < 0 {
		panic("gexpect: No command given to spawn")
	}
	path, err := exec.LookPath(splitArgs[0])
	if err != nil {
		panic("fail to LookPath")
	}

	if numArguments >= 1 {
		wrapper.Cmd = exec.Command(path, splitArgs[1:]...)
	} else {
		wrapper.Cmd = exec.Command(path)
	}
	wrapper.buf = new(buffer)

	return wrapper, nil
}

func RunCmd(cmd string) (string, error) {
	child, _ := Spawn(cmd)

	var b3 bytes.Buffer

	io.Copy(&b3, &child.buf.b)
	go io.Copy(&b3, child.buf.f)
	go io.Copy(child.buf.f, os.Stdin)
	WaitTimeout(child.Cmd, 10000*time.Second)
	s := b3.String()
	if s == "" {
		panic("something wrong")
	}
	// return "tmp", nil
	return s, nil
}

const KillGrace = 5 * time.Second

// WaitTimeout waits for the given command to finish with a timeout.
// It assumes the command has already been started.
// If the command times out, it attempts to kill the process.
func WaitTimeout(c *exec.Cmd, timeout time.Duration) error {
	var kill *time.Timer
	term := time.AfterFunc(timeout, func() {
		err := c.Process.Signal(syscall.SIGTERM)
		if err != nil {
			//log.Printf("E! [agent] Error terminating process: %s", err)
			panic("SIGTERM")
		}

		kill = time.AfterFunc(KillGrace, func() {
			err := c.Process.Kill()
			if err != nil {
				//log.Printf("E! [agent] Error killing process: %s", err)
				panic("KILL")
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
