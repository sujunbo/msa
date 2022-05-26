package http

import "github.com/lanceryou/msa"

func init() {
	msa.RegisterProvider(&serverFactory{})
}
