package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func genPwd(phone string) string {
	r := rand.Intn(1000000)
	s := phone + "_" + strconv.Itoa(int(r))
	has := md5.Sum([]byte(s))
	return fmt.Sprintf("%X", has)
}

type TestOmit struct {
	Uid int32 `json:"uid, omitempty"`
}

func testomitempty() {
	//str := "{\"uid\":20}"
	str := ""
	v := &TestOmit{}
	err := json.Unmarshal([]byte(str), &v)
	fmt.Printf("err = %+v, value = %+v", err, v)
}

type JsonTest struct {
	DataValue string //`json:"data_value"`
}

func ErrTest(err *error) {
	if nil == *err {
		fmt.Println("nil.....")
	}
	*err = errors.New("test")
}

func sliceTest(values *[]int) {
	*values = append(*values, 20)
}

func main() {

	/*
		http.HandleFunc("/health-check", httpTest.HealthCheckHandler)
		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}


		v := JsonTest{
			DataValue: "20",
		}
		str, _ := json.Marshal(v)
		fmt.Printf("%+v",string(str))
	*/
	rand.Seed(time.Now().Unix())
	var v uint64 = rand.Uint64()
	fmt.Printf("%+v\n", v)
	fmt.Println("hello world")
}
