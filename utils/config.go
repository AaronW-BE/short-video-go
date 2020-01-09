package utils

import (
	"MiniVideo/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadConfig() types.Config  {
	config := types.Config{}
	file, err := ioutil.ReadFile(".env.yaml")
	if err != nil {
		panic(err)
	}
	configErr := yaml.Unmarshal(file, &config)
	if configErr != nil {
		panic(configErr)
	}
	return config
}