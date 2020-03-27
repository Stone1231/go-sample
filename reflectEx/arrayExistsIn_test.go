package reflectex

import (
	"reflect"
	"testing"
)

type Name string

const (
	Apple  Name = "apple"
	Banana Name = "banana"
)

func ExistsIn(item interface{}, arrayType interface{}) bool {
	arr := reflect.ValueOf(arrayType)
	switch arr.Kind() {
	case reflect.Array:
	case reflect.Slice:
	default:
		// panic("Invalid data-type")
		return false
	}
	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func TestExistsIn(t *testing.T) {
	type args struct {
		item      interface{}
		arrayType interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				item:      "www",
				arrayType: []string{"www", "ggg", "zzz"},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				item:      Apple,
				arrayType: []Name{Apple, Banana},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExistsIn(tt.args.item, tt.args.arrayType); got != tt.want {
				t.Errorf("ExistsIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
