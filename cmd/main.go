package main

import (
	"audio/internal/api"
	"audio/internal/config"
	"audio/internal/file_make"
	"audio/internal/texter"
	"flag"
	"fmt"
)

func main() {
	conf := flag.String("conf", "config.toml", "config file path")
	flag.Parse()
	cfg, err := config.ParseConfig(*conf)
	if err != nil {
		panic("parse config err:"+err.Error())
	}
	fmt.Printf("cfg:%+v\n", *cfg)
	app := api.NewApp(cfg)
	text := `
	伴随着二次元文化而生的A站，从诞生之初就充斥着亚文化的气质：“ACFun”取意自“Anime Comic Fun”，以及“天下漫友是一家”，一小群屡遭误解、排挤的二次元文化爱好者聚拢在一起，用小圈层的自娱自乐反叛而不羁地解构主流文化	
`
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
