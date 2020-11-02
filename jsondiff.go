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
	result = diff("", json1Map, json2Map)

	for _, item := range result {
		fmt.Println(item.status, item.code, item.message)
	}
	return nil, nil
}

func diff(keyPrefix string, json1Map map[string]interface{}, json2Map map[string]interface{}) []DiffInfo {
	var result = make([]DiffInfo, 0)
	for json1Key, json1Val := range json1Map {
		json2Val, ok := json2Map[json1Key]
		if !ok {
			result = append(result, DiffInfo{status: "error", code: "key_not_exists", message: fmt.Sprintf("%s.%s\t%v\t%v", keyPrefix, json1Key, json1Val, nil)})
			continue
		}
		switch value := json1Val.(type) {
		case int, float64, string: //可以断言到
			if json1Val != json2Val {
				result = append(result, DiffInfo{status: "error", code: "val_not_equal", message: fmt.Sprintf("%s.%s\t%v\t%v", keyPrefix, json1Key, json1Val, json2Val)})
				continue
			}
			break
		case map[string]interface{}: //可以断言到
			value2, ok := json2Val.(map[string]interface{})
			if !ok {
				result = append(result, DiffInfo{status: "error", code: "val_type_not_equal", message: fmt.Sprintf("%s.%s\t%v\t%v", keyPrefix, json1Key, json1Val, json2Val)})
				continue
			}
			childResult := diff(fmt.Sprintf("%s.%s", keyPrefix, json1Key), value, value2)
			if childResult != nil {
				result = append(result, childResult...)
			}
			break
		case interface{}:
			rt := reflect.TypeOf(json1Val)
			switch rt.Kind() {
			case reflect.Slice:
				v1, _ := json1Val.([]interface{})
				v2, ok := json2Val.([]interface{})
				if !ok {
					result = append(result, DiffInfo{status: "error", code: "val_type_not_equal", message: fmt.Sprintf("%s.%s\t%v\t%v", keyPrefix, json1Key, json1Val, json2Val)})
					continue
				}
				if len(v1) != len(v2) {
					result = append(result, DiffInfo{status: "error", code: "val_ary_len_not_equal", message: fmt.Sprintf("%s.%s\t%v\t%v", keyPrefix, json1Key, json1Val, json2Val)})
					continue
				}
				for v1k, v1v := range v1 {
					if v1v != v2[v1k] {
						result = append(result, DiffInfo{status: "error", code: "val_ary_index_val_not_equal", message: fmt.Sprintf("%s.%s[%d]\t%v\t%v", keyPrefix, json1Key, v1k, v1v, v2[v1k])})
					}
				}
				fmt.Println("Slice", v1)
			default:
				fmt.Println("interface{}")
			}
		}
	}
	return result
}
