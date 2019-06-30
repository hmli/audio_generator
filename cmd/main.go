package main

import (
	"audio/internal/api"
	"audio/internal/config"
	"audio/internal/file_make"
	"audio/internal/texter"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	conf := flag.String("conf", "config.toml", "config file path")
	file := flag.String("f", "text.txt", "text file path")
	flag.Parse()
	cfg, err := config.ParseConfig(*conf)
	if err != nil {
		panic("parse config err:"+err.Error())
	}
	fmt.Printf("cfg:%+v\n", *cfg)
	app := api.NewApp(cfg)

	b, err := ioutil.ReadFile(*file)
	if err != nil {
		panic("open file err:"+err.Error())
	}
	text := string(b)
	ss := texter.TextEveryLine([]byte(text))
	param := api.Param{
		Aue: api.AUE_MP3,
		Auf: api.AUF_16k,
		VoiceName: "xiaoyan",
	}
	for _, s := range ss {
		body, err :=  app.DoTts(s, param)
		if err != nil {
			fmt.Println("tts err: ", err.Error())
			continue
		}
		file_make.MakeAudioFile(body, "mp3")
	}

}
