package test_abstract

import (
	"fmt"
	"testing"
)

type IAbstract interface{
	SayHi()
	SayHello()
}

type Father struct {
}

func (f *Father)SayHi() {
	fmt.Println("father")
}

type TestSayHi struct {
	*Father
}

func (t *TestSayHi)TestSayHi(){
	t.SayHi()
}

func (t *Father)SayHello()  {
	fmt.Println("father hello")
}

type Sun struct {
	*Father
}

func NewSun(father *Father) *Sun {
	return &Sun{Father: father}
}

func (s *Sun)SayHi()  {
	fmt.Println("sun")
}

func (s *Sun)SayHello()  {
	fmt.Println("sun hello")
}

type GrandSun struct {
	*Sun
}

func NewGrandSun(sun *Sun) *GrandSun {
	return &GrandSun{Sun: sun}
}

func (g *GrandSun) SayHi() {
	fmt.Println("grand sun")
}

func TestAbstract(t *testing.T)  {
	grandSun := NewGrandSun(NewSun(&Father{}))
	grandSun.SayHi()
	grandSun.Father.SayHello()
}

func GetValueFromMap(values map[int]string, value int)  string{
	return values[value]
}

func TestGetvaluefrommap(t *testing.T)  {
	fmt.Printf("value = %+v", GetValueFromMap(nil, 10))
}

func TestFor(t *testing.T){
	v := []int{1,2,3,4,5}
	for index, value := range v {
		if value == 6 {
			v = append(v[:index], v[index+1:]...)
		}
	}
	fmt.Println(v)
}


