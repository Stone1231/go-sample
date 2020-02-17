package reflectex

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"
)

func Test_setReflect(t *testing.T) {
	bs := &Somestruct{}
	val := reflect.ValueOf(bs).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.FieldByName(field.Name)
		id := field.Tag.Get("id")
		if id == "-" || id == "" {
			continue
		}
		if value.IsValid() {
			switch value.Kind() {
			case reflect.Int:
				value.SetInt(math.MaxInt8)
			case reflect.Int32:
				value.SetInt(math.MaxInt32)
			case reflect.Int64:
				value.SetInt(math.MaxInt64)
			case reflect.Uint64:
				value.SetUint(math.MaxUint64)
			case reflect.String:
				value.SetString(id)
			default:
				fmt.Errorf("field: %s reflect not support (1)", id)
			}
		} else {
			fmt.Errorf("field: %s ,reflect not support (2)", id)
		}
		// fmt.Println(fmt.Sprintf("%v", val))
	}
	fmt.Println(bs.FieldStr)
	fmt.Println(bs.FieldInt)
	fmt.Println(bs.FieldInt64)
}

func Test_getReflect(t *testing.T) {
	vs := Somestruct{}
	vs.FieldStr = "str"
	vs.FieldInt = 2000 //math.MaxInt8
	vs.FieldInt64 = math.MaxInt64
	val := reflect.ValueOf(vs)
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.FieldByName(field.Name)
		id := field.Tag.Get("id")
		if id == "-" || id == "" {
			continue
		}
		if value.IsValid() {
			var tmp string
			switch value.Kind() {
			case reflect.Int:
				tmp = strconv.Itoa(int(value.Int()))
			case reflect.Int64:
				tmp = strconv.Itoa(int(value.Int()))
			case reflect.Uint64:
				tmp = strconv.Itoa(int(value.Uint()))
			case reflect.String:
				tmp = value.String()
			default:
				fmt.Errorf("field: %s reflect not support (1)", id)
			}
			fmt.Println(tmp)
		} else {
			fmt.Errorf("field: %s reflect not support (2)", id)
		}
	}
}
