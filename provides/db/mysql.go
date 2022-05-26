package db

import (
	"github.com/facebookgo/inject"
	"github.com/lanceryou/micro/db"
	"github.com/lanceryou/msa"
	"go.uber.org/config"
)

func init() {
	msa.RegisterProvider(&mysqlFactory{})
}

type mysqlFactory struct{}

func (n *mysqlFactory) NewProvider(conf config.Provider) msa.Provider {
	var cv config.Value
	if cv = conf.Get("db"); !cv.HasValue() {
		return nil
	}

	var opts map[string]*db.Options
	if err := cv.Populate(&opts); err != nil {
		panic(err)
	}

	return msa.ProvideFunc(func() []*inject.Object {
		var objects []*inject.Object
		for k, v := range opts {
			conn := db.Open(v)
			name := "db." + k

			objects = append(objects, &inject.Object{Name: name, Value: conn})
		}
		return objects
	})
}
