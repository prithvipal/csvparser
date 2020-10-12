package csvparser

import (
	"fmt"
	"reflect"
)

// GetTagName ...
func GetTagName(i interface{}) {
	e := reflect.ValueOf(i) //.Elem()
	for i := 0; i < e.NumField(); i++ {
		tag := e.Type().Field(i).Tag
		tagName, _ := tag.Lookup("csv")
		fmt.Printf("%v \n", tagName)
	}
}

// GetFieldNameTypeAndValue ...
func GetFieldNameTypeAndValue(i interface{}) {
	e := reflect.ValueOf(i) //.Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
	}
}
