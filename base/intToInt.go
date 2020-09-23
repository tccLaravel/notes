package base

import (
	"fmt"
)


func IntToInt(i int)  {
	IntToInt8(i)
	intToInt16(i)
	intToInt32(i)
	intToInt64(i)
}

// 整型1 <=> 整型2   需要注意 长字节 转 短字节 时候的 越界问题
func IntToInt8(i int) (i8 int8) {
	i8 = int8(i)   //2的8次幂
	fmt.Printf("int %+v conversion to int8 %+v \n",i,i8)
	if i > 255{
		fmt.Println("注意越界问题")
	}
	return i8
}

func intToInt16(i int) (i16 int16) {
	i16 = int16(i)
	fmt.Printf("int %+v conversion to int16 %+v \n",i,i16)
	return i16
}

func intToInt32(i int) (i32 int32) {
	i32 = int32(i)
	fmt.Printf("int %+v conversion to int32 %+v \n",i,i32)
	return i32
}

func intToInt64(i int) (i64 int64) {
	i64 = int64(i)
	fmt.Printf("int %+v conversion to int64 %+v \n",i,i64)
	return i64
}
