package csvparser

import (
	"fmt"
	"reflect"
	"strconv"
)

func getTagNames(i interface{}) []string {
	var tags []string
	s := reflect.ValueOf(i)
	e := s.Index(0)
	for i := 0; i < e.NumField(); i++ {
		tag := e.Type().Field(i).Tag
		tagName, _ := tag.Lookup("csv")
		tags = append(tags, tagName)
	}
	return tags
}

func getFieldValues(i interface{}) [][]string {
	s := reflect.ValueOf(i)
	records := make([][]string, 0)
	for i := 0; i < s.Len(); i++ {
		e := s.Index(i)
		raw := make([]string, 0)
		for j := 0; j < e.NumField(); j++ {
			var val string
			field := e.Field(j)
			kind := e.Type().Field(j).Type.Kind()
			switch kind {
			case reflect.String:
				val = field.String()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				n := field.Int()
				val = strconv.Itoa(int(n))
			case reflect.Float32, reflect.Float64:
				n := field.Float()
				val = strconv.FormatFloat(n, 'E', -1, 64)
			case reflect.Bool:
				n := field.Bool()
				val = strconv.FormatBool(n)
			}
			raw = append(raw, val)
		}
		records = append(records, raw)
	}
	return records
	//e := reflect.ValueOf(i) //.Elem()

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

// Marshal ...
func Marshal(i interface{}) [][]string {
	data := make([][]string, 0)
	tagNames := getTagNames(i)
	data = append(data, tagNames)
	records := getFieldValues(i)
	data = append(data, records...)
	return data
}
