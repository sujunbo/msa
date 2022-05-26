package msa

import (
	"go.uber.org/config"
)

var conf config.Provider

func initConfig() {
	var err error
	if conf, err = config.NewYAML(
		config.File("./config/service.yml"),
	); err != nil {
		panic(err)
	}
	return
}
