package msa

import (
	"github.com/facebookgo/inject"
	"go.uber.org/config"
)

type Inject struct {
	inject.Graph
	Vals []interface{}
}

var injects Inject

type Provider interface {
	Provide() []*inject.Object
}

type ProvideFactory interface {
	NewProvider(provider config.Provider) Provider
}

type ProvideFunc func() []*inject.Object

func (f ProvideFunc) Provide() []*inject.Object {
	return f()
}

func RegisterProvider(objs ...interface{}) {
	injects.Vals = append(injects.Vals, objs...)
}

func NewProvider(vals ...interface{}) ProvideFunc {
	return func() []*inject.Object {
		var objects []*inject.Object
		for _, val := range vals {
			if val == nil {
				continue
			}

			if v, ok := val.(ProvideFactory); ok {
				if p := v.NewProvider(conf); p == nil {
					continue
				} else {
					objects = append(objects, p.Provide()...)
				}
			}

			if v, ok := val.(Provider); ok {
				objects = append(objects, v.Provide()...)
			} else {
				objects = append(objects, &inject.Object{Value: val})
			}
		}

		return objects
	}
}
