package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

func IntToBool(a int) bool {
	switch a {
	case 0, -1:
		return false
	}
	return true
}

func BookmarkType(a int64) string {
	switch a {
	case 1:
		return "url"
	default:
		return "folder"
	}
}

func TimeStampFormat(stamp int64) time.Time {
	s1 := time.Unix(stamp, 0)
	if s1.Local().Year() > 9999 {
		return time.Date(9999, 12, 13, 23, 59, 59, 0, time.Local)
	}
	return s1
}

func TimeEpochFormat(epoch int64) time.Time {
	maxTime := int64(99633311740000000)
	if epoch > maxTime {
		return time.Date(2049, 1, 1, 1, 1, 1, 1, time.Local)
	}
	t := time.Date(1601, 1, 1, 0, 0, 0, 0, time.UTC)
	d := time.Duration(epoch)
	for i := 0; i < 1000; i++ {
		t = t.Add(d)
	}
	return t
}

func WriteFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return nil
	}
	return err
}

func FormatFilename(dir, browser, filename, format string) string {
	r := strings.Replace(strings.TrimSpace(strings.ToLower(browser)), " ", "_", -1)
	p := path.Join(dir, fmt.Sprintf("%s_%s.%s", r, filename, format))
	return p
}

func MakeDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		return os.Mkdir(dirName, 0700)
	}
	return nil
}
