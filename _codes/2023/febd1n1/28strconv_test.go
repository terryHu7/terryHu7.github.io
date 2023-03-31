package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

const RECORD_TMP_SUFFIX string = ".rectemp"

func TestStrconv(t *testing.T) {

	path := "/home/user/Videos/eswin/xxx_xxx.fmp4"

	fmt.Println(hiddenPathGenerator(path))
}

func hiddenPathGenerator(path string) string {
	dirname, filename := filepath.Dir(path), filepath.Base(path)

	norectemp := strings.TrimSuffix(filename, RECORD_TMP_SUFFIX) //f s doesn't end with ".rectemp", s is returned unchanged.
	nosuffix := strings.TrimSuffix(norectemp, filepath.Ext(norectemp))

	hiddenPath := filepath.Join(dirname, "."+nosuffix)
	return hiddenPath
}
