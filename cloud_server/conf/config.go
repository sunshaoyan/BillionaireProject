package conf

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var configFile []byte
var Configure ServiceConfig

func ConfInit() {
	var err error
	configPath := "conf/config.yaml"
	configFile, err = ioutil.ReadFile(configPath)
	if err != nil {
		log.Panicf("config file not found: %s", configPath)
		return
	}
	err = GetServiceConfig()
	if err != nil {
		log.Panicf("get services Configure error", err.Error())
	}
}

func GetServiceConfig() (err error) {
	err = yaml.Unmarshal(configFile, &Configure)
	return err
}

type ServiceConfig struct {
	Port          string      `yaml:"port"`
	Logger        Logger      `yaml:"logger"`
	MongodbURL    string      `yaml:"mongodburl"`
	MongodbMode   int         `yaml:"mongodbmode"`
	Certification string      `yaml:"certification"`
	Restrict      Restriction `yaml:"restriction"`
}

type Restriction struct {
	Pub    []string `yaml:"publish"`
	Sub    []string `yaml:"subscribe"`
	PubSub []string `yaml:"pubsub"`
}

type Logger struct {
	Logdir   string `yaml:"logdir"`
	FileName string `yaml:"filename"`
	Level    string `yaml:"level"`
	Format   string `yaml:"format"` // text or json
}
