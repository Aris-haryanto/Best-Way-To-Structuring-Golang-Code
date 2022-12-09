package services

import "fmt"

type Configs struct {
	BaseUrl string `yaml:"BASE_URL"`
}

// define global var with default config
var (
	config *Configs
)

// set this from server.go
func InitConfig(conf *Configs) {
	config = conf

	fmt.Println(config.BaseUrl)
}
