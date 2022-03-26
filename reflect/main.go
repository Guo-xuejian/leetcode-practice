package main

import (
	"fmt"
	"reflect"
)

type J struct {
	a string
	b string `json:"B"`
	C string
	D string `json:"DD" otherTag:"good"`
}

func printTag(stru interface{}) {
	t := reflect.TypeOf(stru).Elem()
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("结构体内第%v个字段 %v 对应的json tag是 %v , 还有otherTag？ = %v \n", i+1, t.Field(i).Name, t.Field(i).Tag.Get("json"), t.Field(i).Tag.Get("otherTag"))
	}
}

func main() {
	j := J{
		a: "1",
		b: "2",
		C: "3",
		D: "4",
	}
	printTag(&j)
}
