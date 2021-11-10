package jsondiff

import (
	"bytes"
	"fmt"
	"reflect"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Code 错误码类型
type Code string

// KeyNotExists 字段Key不存在
var KeyNotExists Code = "KeyNotExists"

// ValueNotEqual Value不同
var ValueNotEqual Code = "ValueNotEqual"

// BoolValueFalseOrNull bool类型值false或null类型不同
var BoolValueFalseOrNull Code = "BoolValueFalseOrNull"

// ValueTypeNotEqual Value类型不同
var ValueTypeNotEqual Code = "ValueTypeNotEqual"

// ValueArrayLengthNotEqual Value Arrry长度不同
var ValueArrayLengthNotEqual Code = "ValueArrayLengthNotEqual"

// Status .
type Status string

// StatusError 验证错误
var StatusError Status = "error"

// DiffInfo .
type DiffInfo struct {
	Status  Status
	Code    Code
	Field   string
	Message string
}

var mapPool = sync.Pool{
	New: func() interface{} {
		return make(map[string]interface{}, 4)
	},
}

// diffMap 递归比较
func diffMap(fieldPrefix string, json1Map map[string]interface{}, json2Map map[string]interface{}) (result []DiffInfo) {
	result = make([]DiffInfo, 0)
	for json1Key, json1Value := range json1Map {
		json2Value, ok := json2Map[json1Key]
		if !ok {
			result = append(result, DiffInfo{Status: StatusError, Code: KeyNotExists, Field: fmt.Sprintf("%s.%s", fieldPrefix, json1Key), Message: fmt.Sprintf("%v\t%v", json1Value, nil)})
			continue
		}
		diffInterfaceResult := diffInterface(fmt.Sprintf("%s.%s", fieldPrefix, json1Key), json1Value, json2Value)
		result = append(result, diffInterfaceResult...)
	}
	return result
}

// diffInterface 比较对象
func diffInterface(fieldPrefix string, json1Value interface{}, json2Value interface{}) (result []DiffInfo) {
	result = make([]DiffInfo, 0)
	switch json1TypeValue := json1Value.(type) {
	case int, float64, string: //可以断言到
		if json1Value != json2Value {
			result = append(result, DiffInfo{Status: StatusError, Code: ValueNotEqual, Field: fieldPrefix, Message: fmt.Sprintf("%v\t%v", json1Value, json2Value)})
			return
		}
		return
	case bool:
		// fmt.Println(fieldPrefix, json1Value, json2Value)
		_, ok := json2Value.(bool)
		if !ok {
			result = append(result, DiffInfo{Status: StatusError, Code: ValueTypeNotEqual, Field: fieldPrefix, Message: fmt.Sprintf("%v\t%v", json1Value, json2Value)})
		}
		// fmt.Println(fieldPrefix, json1Value, json2Value)
		if json1Value != json2Value {
			//如果源数据为false但新数据为null
			if !json1TypeValue && json2Value == nil {
				result = append(result, DiffInfo{Status: StatusError, Code: BoolValueFalseOrNull, Field: fieldPrefix, Message: fmt.Sprintf("%v\t%v", json1Value, json2Value)})
				return
			}
			// fmt.Println(fieldPrefix, json1Value, json2Value)
			result = append(result, DiffInfo{Status: StatusError, Code: ValueNotEqual, Field: fieldPrefix, Message: fmt.Sprintf("%v\t%v", json1Value, json2Value)})
			return
		}
		return
	case map[string]interface{}: //可以断言到
		json2TypeValue, ok := json2Value.(map[string]interface{})
		if !ok {
			result = append(result, DiffInfo{Status: StatusError, Code: ValueTypeNotEqual, Field: fieldPrefix, Message: fmt.Sprintf("%v\t%v", json1Value, json2Value)})
			return
		}
		childResult := diffMap(fieldPrefix, json1TypeValue, json2TypeValue)
		if childResult != nil {
			result = append(result, childResult...)
		}
		return
	case interface{}:
		rt := reflect.TypeOf(json1Value)
		switch rt.Kind() {
		case reflect.Slice:
			json1ValueArray, _ := json1Value.([]interface{})
			json2ValueArray, ok := json2Value.([]interface{})
			if !ok {
				result = append(result, DiffInfo{Status: StatusError, Code: ValueTypeNotEqual, Field: fieldPrefix, Message: fmt.Sprintf("%v\t%v", json1Value, json2Value)})
				return
			}
			if len(json1ValueArray) != len(json2ValueArray) {
				result = append(result, DiffInfo{Status: StatusError, Code: ValueArrayLengthNotEqual, Field: fieldPrefix, Message: fmt.Sprintf("%v\t%v", json1Value, json2Value)})
				return
			}
			//需要判断此时数组内的元素类型
			for json1ValueArrayIndex, json1ValueArrayItem := range json1ValueArray {
				diffInterfaceResult := diffInterface(fmt.Sprintf("%s.[%d]", fieldPrefix, json1ValueArrayIndex), json1ValueArrayItem, json2ValueArray[json1ValueArrayIndex])
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
	// json1Map := mapPool.Get().(map[string]interface{})
	// json2Map := mapPool.Get().(map[string]interface{})
	json1Map := make(map[string]interface{}, 4)
	json2Map := make(map[string]interface{}, 4)

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
	// mapPool.Put(json1Map)
	// mapPool.Put(json2Map)
	return result, nil
}
