package autoassign

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/yuewokeji/huifupay/utils"
	"reflect"
	"strings"
)

// 自动赋值
// 汇付部分json对象，转成json字符串，保存为string类型
// 该包用来支持json对象与字符串互转
//
// example：
// type Person struct {
//	 Name       string `json:"name"`
//	 NameObject Name   `json:"-" autoassign:"Name"`
//	 Age        int    `json:"age"`
// }
// Name是汇付字段，NameObject是自动赋值的字段

// 注意：
// 如果NameObject为空时，期望忽略Name字段：
// 1、Name的json tag为：`json:"name，omitempty"`
// 2、NameObject对象为空，或实现 EmptyValue 接口

// tag内容为对应string字段的字段名
const tagName = "autoassign"

var errPointerToStruct = errors.New("object must be a *struct")

// JSONStringToObject 将一个json字符串，转成对象
func JSONStringToObject(obj interface{}) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errPointerToStruct
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return jsonStringToObject(v)
}

func jsonStringToObject(v reflect.Value) error {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		fieldName := fieldType.Tag.Get(tagName)
		if fieldName == "" {
			if field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct) {
				err := jsonStringToObject(field)
				if err != nil {
					return err
				}
			}
			continue
		}
		if !field.IsValid() {
			return fmt.Errorf("autoassign field `%s` invalid", fieldType.Name)
		}
		if !field.CanSet() {
			return fmt.Errorf("autoassign field `%s` can't set", fieldType.Name)
		}

		found := v.FieldByName(fieldName)
		if !found.IsValid() {
			return fmt.Errorf("autoassign field `%s` invalid", fieldName)
		}
		if !field.CanSet() {
			return fmt.Errorf("autoassign field `%s` can't set", fieldName)
		}

		if found.Kind() == reflect.Ptr {
			if found.IsNil() {
				continue
			}
			found = found.Elem()
		}
		if found.Kind() != reflect.String {
			return fmt.Errorf("autoassign field `%s` must a string or *string", fieldName)
		}
		// 内容为空不处理
		if found.String() == "" {
			continue
		}

		instance := reflect.New(field.Type())
		ptr := instance.Interface()
		err := json.Unmarshal([]byte(found.String()), ptr)
		if err != nil {
			return fmt.Errorf("autoassign field `%s` unmarshal: %s", fieldName, err.Error())
		}
		field.Set(reflect.ValueOf(ptr).Elem())
	}

	return nil
}

// ObjectToJSONString 将一个对象，转成json字符串
func ObjectToJSONString(obj interface{}) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errPointerToStruct
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return objectToJSONString(v)
}

func objectToJSONString(v reflect.Value) error {
	t := v.Type()
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		fieldName := fieldType.Tag.Get(tagName)
		if fieldName == "" {
			if field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct) {
				err := objectToJSONString(field)
				if err != nil {
					return err
				}
			}
			continue
		}

		if !field.IsValid() {
			return fmt.Errorf("autoassign field `%s` invalid", fieldType.Name)
		}
		if !field.CanSet() {
			return fmt.Errorf("autoassign field `%s` can't set", fieldType.Name)
		}

		found := v.FieldByName(fieldName)
		if !found.IsValid() {
			return fmt.Errorf("autoassign field `%s` invalid", fieldName)
		}
		if !found.CanSet() {
			return fmt.Errorf("autoassign field `%s` can't set", fieldName)
		}

		if found.Kind() == reflect.Ptr {
			if found.IsNil() {
				continue
			}
			found = found.Elem()
		}
		if found.Kind() != reflect.String {
			return fmt.Errorf("autoassign field `%s` must a string or *string", fieldName)
		}

		// 汇付接口，部分字段为空时，如果传一个空值，可能返回98888888（系统错误）
		isEmpty := false
		if reflect.PtrTo(field.Type()).Implements(reflect.TypeOf((*EmptyValue)(nil)).Elem()) {
			isEmpty = field.Addr().Interface().(EmptyValue).IsEmpty()
		} else {
			isEmpty = utils.IsEmptyValue(field)
		}

		if isEmpty {
			foundType, _ := t.FieldByName(fieldName)
			_, attr, _ := strings.Cut(foundType.Tag.Get("json"), ",")
			if attr == "omitempty" {
				continue
			}
		}
		b, err := json.Marshal(field.Interface())
		if err != nil {
			return fmt.Errorf("autoassign field `%s` unmarshal: %s", fieldType.Name, err.Error())
		}
		found.SetString(string(b))
	}

	return nil
}

// EmptyValue 用来判断一个对象是否为空
type EmptyValue interface {
	IsEmpty() bool
}
