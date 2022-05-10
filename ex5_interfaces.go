package main

import (
	"fmt"
)

type Any interface{}
type M map[Any]map[Any]Any
type TypeField string

const (
	ObjectF TypeField = "object"
	ObjectArrayF TypeField = "object_array"
)

func method(val Any) Any {
	switch val.(type){
	case int:
		return val.(int) + 1
	case float64:
		return val.(float64) + 1.0
	case string:
		return val.(string) + "..."
	case []string:
		v := make([]Any, 0)
		for _, i := range val.([]string){
			v = append(v, "<"+i+">")
		}
		return v
	}
	return nil
}

func anyRun(){
	inter := []Any{"a",22,[]string{"golang", "python"}}
	for _, val := range inter{
		fmt.Printf("[INFO]--[%+v]\n",val)
		fmt.Printf("[INFO]--[%+v]\n", method(val))
	}
}

func INFO(val... interface{}){
	fmt.Println(val...)
}

func validateRange (val interface{}) {
	val, ok := val.(float64)
	fmt.Printf("[info]--[%+v]--[%+v]", val, ok)

}

var testData1 = M{M{"a":1, "b":2,;p[.,]},M{"d":3}}


func parsingFields(typeField TypeField, valueField interface{}) (error) {
	if typeField == ObjectF{
		switch valueField.(type){
		case int, int8, int16, int32, int64, float32, float64:
			return nil
		case []Any:
			return nil
		}
		return nil
	}
	if typeField == ObjectArrayF{
		return nil
	}	
	return nil
}


func main(){
	// anyRun()
	parsingFields(ObjectF, testData1)
}