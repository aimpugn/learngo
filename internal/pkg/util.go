package pkg

import (
	"reflect"
)

type util struct {
}

func getType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}
