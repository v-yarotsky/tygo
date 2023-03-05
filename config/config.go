package config

import (
	"io/ioutil"
	"log"

	"github.com/v-yarotsky/tygo/tygo"
	"gopkg.in/yaml.v2"
)

const defaultFallbackType = "any"

func ReadFromFilepath(cfgFilepath string) tygo.Config {
	b, err := ioutil.ReadFile(cfgFilepath)
	if err != nil {
		log.Fatalf("Could not read config file from %s: %v", cfgFilepath, err)
	}
	conf := tygo.Config{}
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		log.Fatalf("Could not parse config file froms: %v", err)
	}

	// apply defaults
	for _, packageConf := range conf.Packages {
		if packageConf.FallbackType == "" {
			packageConf.FallbackType = defaultFallbackType
		}
	}

	return conf
}
