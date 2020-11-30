package mongodemo

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"

	//"gopkg.in/mgo.v2/bson"
	"testing"
)

type Person struct {
	Name  string
	Phone string `bson:"phone,omitempty"`
}

type ClueInfo struct {
	OwnerId   uint32 //`bson:"ownerid"`
	Status    int32  //`bson:"status,omitempty"`    //状态
	SubStatus int32  //`bson:"substatus,omitempty"` //子状态
	Name      string
}

type TestStruct struct {
	Name string
	ID   int32
}

func TestBson(t *testing.T) {
	data, err := bson.Marshal(&Person{Name: "Bob"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", string(data))
}

func TestMarshal(t *testing.T) {
	/*
		sourceClue := ClueInfo{
			OwnerId: 50959560,
		}

	*/

	data, err := bson.Marshal(&ClueInfo{OwnerId: 123})
	fmt.Printf("value = %q, error = %+v\n", data, err)
	value := ClueInfo{}
	err = bson.Unmarshal(data, &value)
	fmt.Printf("error = %+v, value = %+v\n", err, value)

	data, err = bson.Marshal(&ClueInfo{OwnerId: 123})
	fmt.Printf(" value = %q, error = %+v\n", data, err)
	value = ClueInfo{}
	err = bson.Unmarshal(data, &value)
	fmt.Printf("error = %+v, value = %+v\n", err, value)
}

func TestStructStruct(t *testing.T) {
	fmt.Println("start")
	data, err := bson.Marshal(&ClueInfo{OwnerId: 123})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", data)

	value := ClueInfo{}
	err2 := bson.Unmarshal(data, &value)
	if err2 != nil {
		panic(err)
	}
	fmt.Println("value:", value)

	mmap := bson.M{}
	err3 := bson.Unmarshal(data, mmap)
	if err3 != nil {
		panic(err)
	}
	fmt.Println("mmap:", mmap)

}
