/******
** @创建时间 : 2022/6/4 12:23
** @作者 : MUGUAGAI
******/
package util

import (
	"encoding/json"
	"strconv"
)

func String2Int(strArr []string) []int64 {
	res := make([]int64, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.ParseInt(val, 10, 64)
	}

	return res
}

func Int2String(strArr []int64) []string {
	res := make([]string, len(strArr))

	for index, val := range strArr {
		res[index] = strconv.FormatInt(val, 10)
	}

	return res
}

func MapToJson(param map[string]interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}
