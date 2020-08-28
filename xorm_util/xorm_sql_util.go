package mysql_util

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
)

//生成INSERT INTO ON DUPLICATE KEY UPDATE SQL语句和参数
func GenUpsertSQLAndParams(data interface{}, insertIgnoreFieldSet *hashset.Set, updateIgnoreFieldSet *hashset.Set) (string, []interface{}) {
	return GenBatchUpsertSQLAndParams([]interface{}{data}, insertIgnoreFieldSet, updateIgnoreFieldSet)
}

//生成INSERT INTO ON DUPLICATE KEY UPDATE SQL语句和参数
func GenBatchUpsertSQLAndParams(datas []interface{}, insertIgnoreFieldSet *hashset.Set, updateIgnoreFieldSet *hashset.Set) (string, []interface{}) {
	if len(datas) == 0 {
		return "", nil
	}
	//1.定义sqlFormat
	sqlFormat := "INSERT INTO %s(%s) VALUES %s ON DUPLICATE KEY UPDATE %s"

	//2.获取表名
	tableName := getXormTableName(datas[0])
	keyValueMap := getKeyValueMap(datas[0])

	//3.获取插入的和更新的字段
	var insertFieldList []string
	var insertValueFmtList []string
	var updateFieldValueList []string
	for k := range keyValueMap {
		if !insertIgnoreFieldSet.Contains(k) {
			insertFieldList = append(insertFieldList, k)
			insertValueFmtList = append(insertValueFmtList, "?")
		}
		if !updateIgnoreFieldSet.Contains(k) {
			updateFieldValueList = append(updateFieldValueList, fmt.Sprintf("%s=VALUES(%s)", k, k))
		}
	}

	//4.获取单行插入值fmt,形如(?,?,?,?,?,?,?)
	insertValueFmtStr := fmt.Sprintf("(%s)", strings.Join(insertValueFmtList, ","))

	//5.获取插入值list
	var insertValueList []interface{}
	var insertValueFmtStrList []string
	for _, data := range datas {
		keyValueMap := getKeyValueMap(data)
		for _, insertField := range insertFieldList {
			if value, ok := keyValueMap[insertField]; ok {
				insertValueList = append(insertValueList, value)
			}
		}
		insertValueFmtStrList = append(insertValueFmtStrList, insertValueFmtStr)
	}

	insertFieldStr := strings.Join(insertFieldList, ",")
	updateFieldValueStr := strings.Join(updateFieldValueList, ",")
	insertValueFmtStrListStr := strings.Join(insertValueFmtStrList, ",")
	return fmt.Sprintf(sqlFormat, tableName, insertFieldStr, insertValueFmtStrListStr, updateFieldValueStr), insertValueList
}

func getKeyValueMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columnName := getXormColumnName(field.Tag)
		if columnName == "-" {
			continue
		}
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
	return strings.Trim(arr[len(arr)-1], "'")
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
		} else {
			strBuf.WriteRune(ch)
		}
	}
	return strings.Trim(strBuf.String(), "_")
}
