package interfacedemo

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	UID   int `name:"UserId"`
	UserName string
}

func TestTypeOf(test *testing.T) {

	u := &User{10, "123"}
	t := reflect.TypeOf(u)

	fmt.Println(t.Kind())

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fmt.Println(t.Kind())

	fmt.Println(t.Name())
	fmt.Println(t.NumField())

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, t.Field(i).Type)
	}

	fmt.Println()

	v := reflect.ValueOf(u)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		fmt.Println(f.Interface())

		if f.Kind() == reflect.Int {
			fmt.Println(f.Int())
			f.Set(reflect.ValueOf(20))
		}
		if f.Kind() == reflect.String {
			fmt.Println(f.String())
			f.Set(reflect.ValueOf("234"))
		}
	}

	fmt.Printf("user = %+v\n", u)

	s := []interface{}{50, "50"}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.ValueOf(s[i]).Kind() {
			f.Set(reflect.ValueOf(s[i]))
		}
	}

	fmt.Printf("user = %+v\n", u)

}

func findValue(m map[string]interface{}, key, tag string) interface{} {
	fmt.Printf("eky = %+v, tag = %+v\n", key, tag)
	for k, v := range m {
		if k == tag || k == key {
			return v
		}
	}
	return nil
}

func MapToStruct(m map[string]interface{}, u interface{}) {
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()

		if v.Kind() != reflect.Struct {
			panic("need struct")
		}

		for i := 0; i < v.NumField(); i++ {
			name := t.Field(i).Name
			value := findValue(m, name, v.Type().Field(i).Tag.Get("name"))
			if nil != value {
				if reflect.ValueOf(value).Kind() == v.Field(i).Kind() {
					v.Field(i).Set(reflect.ValueOf(value))
				}
			}
		}

	} else {
		panic("need ptr")
	}

}

func TestMap(test *testing.T) {
	m := map[string]interface{}{
		"UserId":   100,
		"UserName": "100",
		"Age":80,
	}

	u := &User{}

	MapToStruct(m, u)
	fmt.Printf("u = %+v\n", u)

	v := reflect.ValueOf(u)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Type().Field(i).Tag.Get("name"))
	}

}
