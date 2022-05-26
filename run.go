package msa

import (
	"flag"
	"fmt"
	"github.com/lanceryou/msa/provides/log"
	"os"
	"os/signal"
	"syscall"
)

var (
	waitSignal = flag.Bool("wait", true, "wait signal")
)

type initialization interface {
	Init()
}

type starter interface {
	Start()
}

// 主流程
func Run(objs ...interface{}) {
	flag.Parse()
	initConfig()
	log.InitLogger(conf)

	RegisterProvider(objs...)
	p := NewProvider(injects.Vals...)
	if err := injects.Provide(p.Provide()...); err != nil {
		panic(err)
	}

	if err := injects.Populate(); err != nil {
		panic(err)
	}

	for _, v := range injects.Objects() {
		if o, ok := v.Value.(initialization); ok {
			o.Init()
		}
	}

	for _, v := range injects.Objects() {
		if o, ok := v.Value.(starter); ok {
			o.Start()
		}
	}

	if *waitSignal {
		wait()
	}
}

func wait() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-ch:
		fmt.Printf("received signal:%v\n", sig)
	}
}
