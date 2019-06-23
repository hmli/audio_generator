package config


type Config struct {
	Appid string `toml:"appid"`
	Key string `toml:"key"`
	Api Api `toml:"api"`
}


type Api struct {
	Tts string `toml:"tts"`
}
