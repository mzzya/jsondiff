package jsondiff

import (
	"encoding/json"
	"fmt"
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
		switch json1Val.(type) {
		case int, float64, string:
			if json1Val != json2Val {
				result = append(result, DiffInfo{status: "error", code: "val_not_equal", message: fmt.Sprintf("%s\t%v\t%v", json1Key, json1Val, json2Val)})
			}
			break
		case []string, []float64, []int:
			json1ValAry, ok := json1Val.([]interface{})
			if !ok {
				fmt.Println("json1Val.([]interface{})", ok)
			}
			fmt.Println(json1ValAry)
			break
		case map[string]interface{}:
			break
		case []map[string]interface{}:
			break
		}
	}

	for _, item := range result {
		fmt.Println(item.status, item.code, item.message)
	}
	return nil, nil
}
