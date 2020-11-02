package jsondiff

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// DiffInfo .
type DiffInfo struct {
	status  string
	code    string
	message string
}

// Diff .
func Diff(json1 string, json2 string) ([]DiffInfo, error) {
	json1Map := make(map[string]interface{})
	json2Map := make(map[string]interface{})
	var err error
	var result = make([]DiffInfo, 0)

	err = json.Unmarshal([]byte(json1), &json1Map)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(json2), &json2Map)
	if err != nil {
		return nil, err
	}
	for json1Key, json1Val := range json1Map {
		json2Val, ok := json2Map[json1Key]
		if !ok {
			result = append(result, DiffInfo{status: "error", code: "key_not_exists", message: fmt.Sprintf("%s\t%v\t%v", json1Key, json1Val, nil)})
			continue
		}
		switch value := json1Val.(type) {
		case int, float64, string: //可以断言到
			if json1Val != json2Val {
				result = append(result, DiffInfo{status: "error", code: "val_not_equal", message: fmt.Sprintf("%s\t%v\t%v", json1Key, json1Val, json2Val)})
			}
			break
		case map[string]interface{}: //可以断言到
			fmt.Println("map", value)
			break
		case interface{}:
			rt := reflect.TypeOf(json1Val)
			switch rt.Kind() {
			case reflect.Slice:
				v, _ := json1Val.([]interface{})
				fmt.Println("Slice", v)
			default:
				fmt.Println("interface{}")
			}
		}
	}

	for _, item := range result {
		fmt.Println(item.status, item.code, item.message)
	}
	return nil, nil
}
