package utils

import "reflect"

// try{}catch{}
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

// 通过反射实例化一个对象
func NewObject[T interface{}](s T) T {
	t := reflect.TypeOf(s).Elem()
	return reflect.New(t).Interface().(T)
}
