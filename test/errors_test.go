package main

import (
	"errors"
	"fmt"
	"testing"
)

var ErrNotFound = errors.New("not found")

func GetErr() error  {
	return fmt.Errorf("err = %w, uid = %+v", ErrNotFound, 1)
}

func TestError(t *testing.T) {
	err := GetErr()
	fmt.Println(errors.Is(err, ErrNotFound))
}

type ISay interface {
	Say()
}

type Father struct {

}

func (f *Father)Say()  {
	fmt.Println("father")
}

type Son struct {
	Father
}

/*
func (s *Son)Say()  {
	fmt.Println("son")
}

 */

func TestPoly(t *testing.T)  {
	var i ISay = &Son{}
	i.Say()
}