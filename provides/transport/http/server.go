package http

import (
	"github.com/lanceryou/micro/transport/http"
	"github.com/lanceryou/msa"
	"go.uber.org/config"
	"os"
)

type serverFactory struct{}

func (s *serverFactory) NewProvider(conf config.Provider) msa.Provider {
	if cfg := getServerConfig(conf); cfg != nil {
		return msa.NewProvider(newHTTPServer(cfg))
	}
	return nil
}

func newHTTPServer(cfg *serverConfig) *http.HTTPServer {
	return http.NewServer(
		http.Addr(cfg.Addr),
	)
}

type serverConfig struct {
	Addr string
}

func getServerConfig(conf config.Provider) *serverConfig {
	var cv config.Value

	if cv = conf.Get("httpserver"); !cv.HasValue() {
		return nil
	}

	addrMap := make(map[string]string)
	if err := cv.Populate(&addrMap); err != nil {
		return nil
	}

	var cfg serverConfig
	cfg.Addr = addrMap["addr"]
	port := os.Getenv("PORT")
	if port != "" {
		cfg.Addr = ":" + port
	}
	return &cfg
}
