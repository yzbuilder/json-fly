package jsonutil

import (
	"fmt"
	"reflect"
)

type JsonUtil struct {
	Translation *Translation
	MapConfig   map[string]interface{}
}

func baseJson() *JsonUtil {
	j := &JsonUtil{
		Translation: &Translation{
			StringRule: baseStringRule,
			IntRule:    baseIntRule,
			FloatRule:  baseFloatRule,
			BoolRule:   baseBoolRule,
			MapRule:    baseMapRule,
			StructRule: baseStructRule,
			SliceRule:  baseSliceRule,
		},
		MapConfig: defaultTestMapConfig,
	}
	return j
}

func NewJson(mapConfig map[string]interface{}) *JsonUtil {
	j := baseJson()
	if mapConfig != nil {
		j.MapConfig = mapConfig
	}
	return j
}

func (j *JsonUtil) SetTranslation(translation *Translation) *JsonUtil {
	j.Translation = translation
	return j
}

func (j *JsonUtil) SetMapConfig(mapConfig map[string]interface{}) *JsonUtil {
	j.MapConfig = mapConfig
	return j
}

var index int = 0

func (j *JsonUtil) JsonSetUp(v reflect.Value, name ...string) {
	index += 1
	fmt.Printf("第 %v 次递归, 名称: %v, 类型: %v, 具体类型: %v \n", index, name, v.Kind(), v.Type())
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", v.Elem())
	case reflect.Slice, reflect.Array:
		if len(name) != 0 {
			if j.Translation.SliceRule(name[0], v, j.MapConfig) {
				return
			}
		}
		if v.Len() == 0 {
			// Slice 和 Array 类型 强行给一个元素 进入下一递归
			newVal := reflect.New(v.Type().Elem())
			v.Set(reflect.Append(v, reflect.ValueOf(newVal.Interface()).Elem()))
			j.JsonSetUp(v.Index(0), name...)
		} else {
			for i := 0; i < v.Len(); i++ {
				j.JsonSetUp(v.Index(i), name...)
			}
		}
	case reflect.Struct:
		if len(name) != 0 {
			if j.Translation.StructRule(name[0], v, j.MapConfig) {
				return
			}
		}
		for i := 0; i < v.NumField(); i++ {
			j.JsonSetUp(v.Field(i), v.Type().Field(i).Name)
		}
	case reflect.Map:
		j.Translation.MapRule(name[0], v, j.MapConfig)
	case reflect.Ptr:
		if v.IsNil() {
			// 空指针 需要 按照指针的类型定义一个新的地址 使得源数据的地址是可被寻址的
			v.Set(reflect.ValueOf(reflect.New(v.Type().Elem()).Interface()))
			// fmt.Println("空指针赋值完成 准备解引用指针", v.Type().Elem().Kind())
			j.JsonSetUp(reflect.Indirect(v), name...)
		} else {
			j.JsonSetUp(v.Elem(), name...)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", v.Elem())
		} else {
			j.JsonSetUp(v.Elem(), name...)
		}
	case reflect.String:
		j.Translation.StringRule(name[0], v, j.MapConfig)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		j.Translation.IntRule(name[0], v, j.MapConfig)
	case reflect.Float32, reflect.Float64:
		j.Translation.FloatRule(name[0], v, j.MapConfig)
	case reflect.Bool:
		j.Translation.BoolRule(name[0], v, j.MapConfig)
	default: // basic types, channels, funcs
		return
	}
}
