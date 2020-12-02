package sort_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type SortDemo struct {
	ID     int32  `json:"id"`
	Status *int32 `json:"status,omitempty"`
}

type Father struct {
}

func (f *Father) Print() string {
	return fmt.Sprintf("hello")
}

type Son struct {
	Father
}

func (s *Son) Print1() {
	fmt.Println(s.Print())
}

func TestPrint(t *testing.T) {
	s := Son{}
	s.Print1()
}

func TestSort(t *testing.T) {
	value := SortDemo{}
	valueJson, _ := json.Marshal(value)
	fmt.Println(string(valueJson))
	fmt.Printf("%+v\n", value)
	value1 := SortDemo{}
	json.Unmarshal(valueJson, &value1)

	fmt.Printf("%+v\n", value1)
}

func TestSliceCopy(t *testing.T) {
	src := []int{1, 2, 3, 4}
	dst := make([]int, len(src), len(src))
	copy(dst, src[:len(src):len(src)])
	fmt.Println(dst)
	fmt.Println(src)
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)
	fmt.Println(slice2)
}
