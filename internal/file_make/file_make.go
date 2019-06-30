package file_make

import (
	"crypto/md5"
	"fmt"
	"os"
	"time"
)

func MakeAudioFile(body []byte, extension string) {
	now := time.Now()
	hash := fmt.Sprintf("%x", md5.Sum(body))
	filename := now.Format("060102-150405-") + hash[:4] + "." + extension
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("file create err:", err.Error())
		return
	}
	f.Write(body)
	f.Close()
}
