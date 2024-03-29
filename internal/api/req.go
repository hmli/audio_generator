package api

import (
	"audio/internal/config"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type App struct {
	client *http.Client
	config *config.Config
}

func NewApp(config *config.Config) *App {
	return &App{
		config: config,
		client: http.DefaultClient,
	}
}


type Param struct {
	Aue string
	Auf string
	VoiceName string
}

const (
	AUE_RAW = "raw"
	AUE_MP3 = "mp3"
	AUF_16k = "audio/L16;rate=16000"
	AUF_8k = "audio/L16;rate=8000"
)



func (app *App) postReq(link string, text string, param map[string]string) (respBody []byte, err error){
	fmt.Println("开始生成")
	defer fmt.Println("生成结束")
	curtime := strconv.FormatInt(time.Now().Unix(), 10)
	b, _ := json.Marshal(param)
	b64param := base64.StdEncoding.EncodeToString(b)
	w := md5.New()
	io.WriteString(w, app.config.Key+curtime+b64param)

	checksum := fmt.Sprintf("%x", w.Sum(nil))
	data := url.Values{}
	data.Set("text", text)
	body := strings.NewReader(data.Encode())
	req, _ := http.NewRequest("POST", link, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-CurTime", curtime)
	req.Header.Set("X-Appid", app.config.Appid)
	req.Header.Set("X-Param", b64param)
	req.Header.Set("X-CheckSum", checksum)
	resp, err := app.client.Do(req)
	if err != nil {
		return respBody, err
	}
	if resp == nil || resp.Body == nil {
		return respBody, errors.New("resp or body nil")
	}
	defer resp.Body.Close()
	respBody, _ = ioutil.ReadAll(resp.Body)
	if resp.Header.Get("Content-type") != "audio/mpeg" {
		fmt.Println("出错：", string(respBody))
		return
	}
	return
}

