package texter

import (
	"io/ioutil"
	"strings"
)

const (
	_ = iota
	AutoMerge // 自动按行合并成尽可能接近1000字节的
	EveryLine // 一行一个文件
)


func readFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}


func TextAutoMerge(raw []byte) []string {
	lines := TextEveryLine(raw)
	newLines := make([]string, 0, len(lines))
	temp := ""
	for __, line := range lines {
		if len(line) + len(temp) < 990 {
			temp = temp + "\n" + line
		} else {
			newLines = append(newLines, temp)
			temp = ""
		}

	}
}

func TextEveryLine(raw []byte) []string {

	lines := strings.Split(string(raw), "\n")
	for i :=0; i<len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
		if len(lines[i]) == 0 {
			lines = append(lines[0:i], lines[i+1:]...)
		}
	}
	return lines
}
