package config

import "github.com/BurntSushi/toml"

func ParseConfig(path string) (*Config, error) {
	config := new(Config)
	_, err := toml.DecodeFile(path, config)
	if err != nil {
		return nil, err
	}
	return config, err
}

