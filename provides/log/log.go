package log

import (
	"fmt"
	"github.com/lanceryou/micro/log"
	"go.uber.org/config"
)

func InitLogger(conf config.Provider) {
	var cv config.Value
	if cv = conf.Get("log"); !cv.HasValue() {
		return
	}

	var cfg log.Option
	if err := cv.Populate(&cfg); err != nil {
		panic(err)
	}
	fmt.Printf("cfg:%v monitor\n", cfg)
	log.InitLogger(&cfg)
}
