package xorm_util

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

//生成INSERT INTO ON DUPLICATE KEY UPDATE SQL语句和参数
func GenUpsertSQLAndParams(data interface{}, insertIgnoreFields []string, updateIgnoreFields []string) (string, []interface{}) {
	sqlFormat := "INSERT INTO %s(%s) VALUES(%s) ON DUPLICATE KEY UPDATE %s"
	insertIgnoreFieldMap := make(map[string]string)
	updateIgnoreFieldsMap := make(map[string]string)
	for _, f := range insertIgnoreFields {
		insertIgnoreFieldMap[f] = f
	}
	for _, f := range updateIgnoreFields {
		updateIgnoreFieldsMap[f] = f
	}

	var insertKeyStrList []string
	var insertValueStrList []string
	var insertValueList []interface{}
	var updateKeyValueStrList []string
	var updateValueList []interface{}
	keyValueMap := getKeyValueMap(data)
	for k, v := range keyValueMap {
		if _, ok := insertIgnoreFieldMap[k]; !ok {
			insertKeyStrList = append(insertKeyStrList, k)
			insertValueStrList = append(insertValueStrList, "?")
			insertValueList = append(insertValueList, v)
		}

		if _, ok := updateIgnoreFieldsMap[k]; !ok {
			updateKeyValueStrList = append(updateKeyValueStrList, fmt.Sprintf("%s=?", k))
			updateValueList = append(updateValueList, v)
		}
	}
	valueList := append(insertValueList, updateValueList...)
	insertKeyStr := strings.Join(insertKeyStrList, ",")
	insertValueStr := strings.Join(insertValueStrList, ",")
	updateKeyValueStr := strings.Join(updateKeyValueStrList, ",")
	return fmt.Sprintf(sqlFormat, getXormTableName(data), insertKeyStr, insertValueStr, updateKeyValueStr), valueList
}

func getKeyValueMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columnName := getXormColumnName(field.Tag)
		columnValue := v.Field(i).Interface()
		if columnName != "" {
			result[columnName] = columnValue
		}
	}
	return result
}

func getXormColumnName(tag reflect.StructTag) string {
	xormStr := tag.Get("xorm")
	arr := strings.Split(xormStr, " ")
	if len(arr) == 0 {
		return ""
	}
	return arr[0]
}

func getXormTableName(data interface{}) string {
	tableType := reflect.TypeOf(data).Name()
	return toSnakeCase(tableType)
}

func toSnakeCase(str string) string {
	var strBuf bytes.Buffer
	for _, ch := range []rune(str) {
		if ch >= 'A' && ch <= 'Z' {
			strBuf.WriteString("_")
			strBuf.WriteRune(ch + 32)
		}else{
			strBuf.WriteRune(ch)
		}
	}
	return strings.Trim(strBuf.String(), "_")
}