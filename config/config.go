package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Database string
	Server   string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("Tweetstack/db.conf", &c); err != nil {
		log.Fatal(err)
	}

}
