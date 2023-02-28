package reflectex

import (
	"errors"
	"reflect"
)

var (
	// ErrTargetNotPtr Target不是指针
	ErrTargetNotPtr = errors.New("target is not point")
)

func Inject(target, impl interface{}) (err error) {
	targetRt := reflect.TypeOf(target)
	if targetRt.Kind() != reflect.Ptr {
		err = ErrTargetNotPtr
		return
	}
	
	return
}
