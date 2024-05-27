package utils

import (
	"reflect"
	"sort"
	"strings"
	"unicode"
)

// StructToMapJSONTag 根据json tag将struct转map
func StructToMapJSONTag(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		name := v.Field(i).Name
		if unicode.IsLower(rune(name[0])) {
			// 不可导出
			continue
		}

		tag := v.Field(i).Tag.Get("json")
		if tag != "" && tag != "-" {
			tagName, option, _ := strings.Cut(tag, ",")
			if option == "omitempty" && IsEmptyValue(reflectValue.Field(i)) {
				continue
			}

			field := reflectValue.Field(i).Interface()

			switch v.Field(i).Type.Kind() {
			case reflect.Interface:
				t := reflect.TypeOf(field)
				if t.Kind() == reflect.Ptr {
					t = t.Elem()
				}
				if t.Kind() == reflect.Struct {
					res[tagName] = StructToMapJSONTag(field)
				} else {
					res[tagName] = field
				}
			case reflect.Struct:
				res[tagName] = StructToMapJSONTag(field)
			default:
				res[tagName] = field
			}
		}
	}
	return res
}

// SortByKeys 递归地对map的键进行排序
func SortByKeys(m map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	keys := make([]string, 0, len(m))

	// 收集所有的键
	for k := range m {
		keys = append(keys, k)
	}
	// 对键进行排序
	sort.Strings(keys)

	// 按照排序后的键顺序重新构建 map
	for _, k := range keys {
		if v, ok := m[k].(map[string]interface{}); ok {
			// 递归处理嵌套的 map
			result[k] = SortByKeys(v)
		} else {
			result[k] = m[k]
		}
	}

	return result
}
