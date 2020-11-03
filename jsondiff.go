package jsondiff

import (
	"bytes"
	"fmt"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// DiffInfo .
type DiffInfo struct {
	Status  string
	Code    string
	Field   string
	Message string
}

// diffMap 递归比较
func diffMap(fieldPrefix string, json1Map map[string]interface{}, json2Map map[string]interface{}) (result []DiffInfo) {
	result = make([]DiffInfo, 0)
	for json1Key, json1Val := range json1Map {
		json2Val, ok := json2Map[json1Key]
		if !ok {
			result = append(result, DiffInfo{Status: "error", Code: "key_not_exists", Field: fmt.Sprintf("%s.%s", fieldPrefix, json1Key), Message: fmt.Sprintf("%v\t%v", json1Val, nil)})
			continue
		}
		diffInterfaceResult := diffInterface(fmt.Sprintf("%s.%s", fieldPrefix, json1Key), json1Val, json2Val)
		result = append(result, diffInterfaceResult...)
	}
	return result
}

// diffInterface 比较对象
func diffInterface(fieldPrefix string, json1Val interface{}, json2Val interface{}) (result []DiffInfo) {
	result = make([]DiffInfo, 0)
	switch value := json1Val.(type) {
	case int, float64, string: //可以断言到
		if json1Val != json2Val {
			result = append(result, DiffInfo{Status: "error", Code: "val_not_equal", Field: fieldPrefix, Message: fmt.Sprintf("%v\t%v", json1Val, json2Val)})
			return
		}
		break
	case map[string]interface{}: //可以断言到
		value2, ok := json2Val.(map[string]interface{})
		if !ok {
			result = append(result, DiffInfo{Status: "error", Code: "val_type_not_equal", Field: fieldPrefix, Message: fmt.Sprintf("%v\t%v", json1Val, json2Val)})
			return
		}
		childResult := diffMap(fieldPrefix, value, value2)
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
				result = append(result, DiffInfo{Status: "error", Code: "val_type_not_equal", Field: fieldPrefix, Message: fmt.Sprintf("%v\t%v", json1Val, json2Val)})
				return
			}
			if len(v1) != len(v2) {
				result = append(result, DiffInfo{Status: "error", Code: "val_ary_len_not_equal", Field: fieldPrefix, Message: fmt.Sprintf("%v\t%v", json1Val, json2Val)})
				return
			}
			//需要判断此时数组内的元素类型
			for v1k, v1v := range v1 {
				diffInterfaceResult := diffInterface(fmt.Sprintf("%s.[%d]", fieldPrefix, v1k), v1v, v2[v1k])
				result = append(result, diffInterfaceResult...)
			}
			// fmt.Println("Slice", v1)
		default:
			fmt.Println("interface{}")
		}
	}
	return
}

// Diff .
func Diff(json1 string, json2 string, ignoreCase bool) (result []DiffInfo, err error) {
	return DiffBytes([]byte(json1), []byte(json2), ignoreCase)
}

// DiffIgnoreCase 比较忽略大小写
func DiffIgnoreCase(json1 string, json2 string) ([]DiffInfo, error) {
	return Diff(json1, json2, true)
}

// DiffBytesIgnoreCase 比较忽略大小写
func DiffBytesIgnoreCase(json1 []byte, json2 []byte) ([]DiffInfo, error) {
	return DiffBytes(json1, json2, true)
}

// DiffBytes .
func DiffBytes(json1 []byte, json2 []byte, ignoreCase bool) (result []DiffInfo, err error) {
	json1Map := make(map[string]interface{})
	json2Map := make(map[string]interface{})

	result = make([]DiffInfo, 0)

	if ignoreCase {
		json1 = bytes.ToLower(json1)
		json2 = bytes.ToLower(json2)
	}

	err = json.Unmarshal(json1, &json1Map)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(json2, &json2Map)
	if err != nil {
		return nil, err
	}
	result = diffMap("", json1Map, json2Map)
	return result, nil
}
