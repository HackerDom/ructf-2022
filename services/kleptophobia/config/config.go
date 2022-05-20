package config

import (
	"io/ioutil"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"kleptophobia/utils"
)

func InitConfig[T proto.Message](filename string, config T) {
	rawConfig, err := ioutil.ReadFile(filename)
	utils.FailOnError(err)

	err = protojson.Unmarshal(rawConfig, config)
	utils.FailOnError(err)
}
