package case_util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"fmt"
)

func TestTransToLowerCamelKeyMap(t *testing.T){
	mapStr := `{"id":727141,"is_deleted":0,"dispath_time":-62135596800,"department_id":null,"discover_time":"","feature":{"test_test":"abc"},"grade":"三年级","import_time":1567495648,"importer":"1","note":"dfadsf","phone":"13115016000","recover_reason":"","source_id":10005,"source_name":"sfdasdf","root_source_id":103,"root_source_name":"测试广告","status":"waitDispath","call_status":"waitDialed"}`

	result := make(map[string]interface{})
	json.Unmarshal([]byte(mapStr), &result)

	transfered := TransToLowerCamelKeyMap(result)

	fmt.Printf("transfered=%+v\n", transfered)
}

func TestUpperCamelCase(t *testing.T) {
	data := map[string]string{
		"":                      "",
		"f":                     "F",
		"foo":                   "Foo",
		" foo_bar\n":            "FooBar",
		" foo-bar\t":            "FooBar",
		" foo bar\r":            "FooBar",
		"HTTP_status_code":      "HttpStatusCode",
		"skip   many spaces":    "SkipManySpaces",
		"skip---many-dashes":    "SkipManyDashes",
		"skip___many_underline": "SkipManyUnderline",
	}

	for in, out := range data {
		converted := UpperCamelCase(in)
		assert.Equal(t, out, converted)
	}
}

func TestLowerCamelCase(t *testing.T) {
	data := map[string]string{
		"":                      "",
		"F":                     "f",
		"foo":                   "foo",
		" foo_bar\n":            "fooBar",
		" foo-bar\t":            "fooBar",
		" foo bar\r":            "fooBar",
		"HTTP_status_code":      "httpStatusCode",
		"skip   many spaces":    "skipManySpaces",
		"skip---many-dashes":    "skipManyDashes",
		"skip___many_underline": "skipManyUnderline",
	}

	for in, out := range data {
		converted := LowerCamelCase(in)
		assert.Equal(t, out, converted)
	}
}