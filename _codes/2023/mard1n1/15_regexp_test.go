package mard1n1

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var CommonPath = "/home/hutianyu/"

func TestRegexp(t *testing.T) {
	re := regexp.MustCompile(`_%(0\d)?d`)

	str := "test_%02d_%d"
	s := re.FindStringSubmatch(str)
	fmt.Println(s)
	fmt.Println(strconv.Atoi(s[len(s)-1]))
}

func reg() {
	for _, path := range []string{`tmp_yyyy-mm-dd-hh-mm-ss.mp4`, `/home/hutianyu/tmp_%02d.mp4`, `tmp_yyyy-mm-dd-hh-mm-ss_%d.mp4`} {
		// fullpath
		fullpath := ""
		if !strings.HasPrefix(path, "/") {
			format := filepath.Ext(path)
			if format == ".mp4" || format == ".mkv" || format == ".flv" {
				fullpath = CommonPath + path
			} else if format == ".ogg" || format == ".aac" || format == ".adts" || format == ".wav" {
				fullpath = CommonPath + path
			} else {
				continue
			}
		} else {
			fullpath = path
		}

		dir, filename := filepath.Split(fullpath)

		//need regular expression or not
		pattern := filename
		if strings.Contains(pattern, "yyyy-mm-dd-hh-mm-ss") {
			pattern = strings.ReplaceAll(filename, `yyyy-mm-dd-hh-mm-ss`, `\d{4}-\d{2}-\d{2}-\d{2}-\d{2}-\d{2}`)
		}
		if strings.Contains(pattern, "_%") {
			r, _ := regexp.Compile(`_%(0\d)?d`)
			if matched := r.MatchString(pattern); matched {
				pattern = r.ReplaceAllString(pattern, `_\d+`)
			}
		}

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println("fail to read dir:", err)
			return
		}

		for _, f := range files {
			if !f.IsDir() {
				matched, _ := regexp.MatchString(pattern, f.Name())
				if matched {
					fmt.Println("File found:", dir+"/"+f.Name())
				}
			}
		}
	}
}
