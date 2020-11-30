package options_demo

import (
	"fmt"
	"testing"
)

type ServerOptions struct {
	option1 string
	option2 string
	option3 string
}

var defaultServerOptions = &ServerOptions{
	option1: "option1",
}

type Services struct {
	Opts *ServerOptions
}

type IServerOption interface {
	apply(o *ServerOptions)
}

func NewServices(opt ... IServerOption) *Services {
	opts := defaultServerOptions
	for _, o := range opt {
		o.apply(opts)
	}

	return &Services{
		Opts: opts,
	}
}

type FuncServerOption struct {
	f func(options *ServerOptions)
}

func (f *FuncServerOption)apply(o *ServerOptions)  {
	f.f(o)
}

func NewServerFuncOption(f func(options *ServerOptions)) *FuncServerOption {
	return &FuncServerOption{
		f: f,
	}
}

func Option2(option2 string) IServerOption {
	return NewServerFuncOption(func(options *ServerOptions) {
		options.option2 = option2
	})
}

func Option3(option3 string) IServerOption {
	return NewServerFuncOption(func(options *ServerOptions) {
		options.option3 = option3
	})
}

func NilInterface(req interface{}) error  {

	if nil == req {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	return nil
}

func TestInterface(t *testing.T)  {
	var service *Services
	NilInterface(service)
}

func TestOptions(t *testing.T) {
	options := []IServerOption{
		Option2("hello"),
		Option3("option3"),
	}
	s := NewServices(options...)
	fmt.Println(s.Opts)
}

