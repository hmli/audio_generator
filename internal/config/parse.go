package config

import "github.com/BurntSushi/toml"

func ParseConfig(b []byte) (*Config, error) {
	config := new(Config)
	_, err := toml.Decode(b, config)
	if err != nil {
		return nil, err
	}
	return config, err
}