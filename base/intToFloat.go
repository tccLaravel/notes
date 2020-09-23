package base

import (
	"fmt"
	"reflect"
	"unsafe"
)

func IntAndFloat(i interface{}) {
	fmt.Printf("i is %+v \n",i)
	fmt.Printf("i is %+v \n",*(*int)(unsafe.Pointer(&i)))
	switch reflect.ValueOf(i).Kind() {
	case reflect.Int:
		IntToFloat32(*(*int)(unsafe.Pointer(&i)))
		IntToFloat64(*(*int)(unsafe.Pointer(&i)))
	case reflect.Float32, reflect.Float64:
		Float32ToInt(*(*float32)(unsafe.Pointer(&i)))
		Float64ToInt(*(*float64)(unsafe.Pointer(&i)))
	default:
		fmt.Printf("%s \n","请输入正确的参数")
	}
}

func IntToFloat32(i int) (f float32) {
	f = float32(i)
	fmt.Printf("int %+v conversion to float32 %+v \n",i,f)
	return f
}

func IntToFloat64(i int) (f float64) {
	f = float64(i)
	fmt.Printf("int %+v conversion to float64 %+v \n",i,f)
	return f
}

func Float32ToInt(f float32) (i int) {
	i = int(f)
	fmt.Printf("float32 %+v conversion to int %+v \n",f,i)
	return i
}

func Float64ToInt(f float64) (i int) {
	i = int(f)
	fmt.Printf("float64 %+v conversion to int %+v \n",f,i)
	return i
}
