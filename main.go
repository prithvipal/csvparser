package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	ID      int
	Title   string
	Price   float32
	Authors []string
}

type T struct {
	f string `one:"1" two:"2"blank:""`
}

func main() {

	book := Book{}
	f1(book)
}

func f1(i interface{}) {
	// typee := reflect.ValueOf(i)
	//i.(typee)
	e := reflect.ValueOf(i) //.Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
	}
}
