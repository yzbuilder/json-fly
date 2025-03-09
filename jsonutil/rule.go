package jsonutil

import (
	"fmt"
	"reflect"
)

type Translation struct {
	StringRule func(name string, v reflect.Value, m map[string]interface{})
	IntRule    func(name string, v reflect.Value, m map[string]interface{})
	FloatRule  func(name string, v reflect.Value, m map[string]interface{})
	BoolRule   func(name string, v reflect.Value, m map[string]interface{})
	MapRule    func(name string, v reflect.Value, m map[string]interface{})
	StructRule func(name string, v reflect.Value, m map[string]interface{}) bool
	SliceRule  func(name string, v reflect.Value, m map[string]interface{}) bool
}

func baseStringRule(name string, v reflect.Value, m map[string]interface{}) {
	for k, val := range m {
		value := reflect.ValueOf(val)
		if k == name {
			if value.Kind() != reflect.String {
				fmt.Printf("类型匹配错误, expected %v but get a %v\n", v.Type(), value.Type())
				return
			}
			v.SetString(value.String())
			return
		}
	}
	v.SetString("this is a default string")
}

func baseIntRule(name string, v reflect.Value, m map[string]interface{}) {
	for k, val := range m {
		value := reflect.ValueOf(val)
		if k == name {
			switch value.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				v.SetInt(value.Int())
				return
			default:
				fmt.Printf("类型匹配错误, expected %v but get a %v\n", v.Type(), value.Type())
				return
			}
		}
	}
	v.SetInt(11111)
}

func baseFloatRule(name string, v reflect.Value, m map[string]interface{}) {
	for k, val := range m {
		value := reflect.ValueOf(val)
		if k == name {
			switch value.Kind() {
			case reflect.Float32, reflect.Float64:
				v.SetFloat(value.Float())
				return
			default:
				fmt.Printf("类型匹配错误, expected %v but get a %v\n", v.Type(), value.Type())
				return
			}
		}
	}
	v.SetFloat(0.0)
}

func baseBoolRule(name string, v reflect.Value, m map[string]interface{}) {
	for k, val := range m {
		value := reflect.ValueOf(val)
		if k == name {
			v.SetBool(value.Bool())
			return
		}
	}
}

func baseMapRule(name string, v reflect.Value, m map[string]interface{}) {
	for _, key := range v.MapKeys() { // 保留原始数据
		v.SetMapIndex(key, v.MapIndex(key))
	}
	for k, val := range m {
		value := reflect.ValueOf(val)
		if k == name {
			if v.Type() != value.Type() {
				fmt.Printf("类型匹配错误, expected %v but get a %v\n", v.Type(), value.Type())
				return
			}
			if v.IsNil() {
				v.Set(value)
				return
			}
			for _, valueKey := range value.MapKeys() {
				v.SetMapIndex(valueKey, value.MapIndex(valueKey))
			}
			return
		}
	}
}

func baseStructRule(name string, v reflect.Value, m map[string]interface{}) bool {
	for k, val := range m {
		value := reflect.ValueOf(val)
		if k == name {
			if v.Type() == value.Type() {
				v.Set(value)
				return true
			}
			fmt.Printf("类型匹配错误, expected %v but get a %v\n", v.Type(), value.Type())
			return false
		}
	}
	return false
}

func baseSliceRule(name string, v reflect.Value, m map[string]interface{}) bool {
	for k, val := range m {
		value := reflect.ValueOf(val)
		if k == name {
			if v.Type() == value.Type() {
				v.Set(value)
				return true
			}
			fmt.Printf("类型匹配错误, expected %v but get a %v\n", v.Type(), value.Type())
			return false
		}
	}
	return false
}
