package main

import (
	"fmt"
	"unsafe"
)

func str(s string)  {
	k := *(*[]string)(unsafe.Pointer(&s))
	fmt.Printf("%s -- %T *(*[]string)(unsafe.Pointer(&s)) is %+v  %T \n",s,s,k,k)
}

/**
如果T2比T1大，并且两者共享等效内存布局，则该转换允许将一种类型的数据重新解释为另一类型的数据。
*/
type MyInt int

/**
底层数据类型一直  尽管 MyInt 与 int 是两种不同的类型  但是底层都是int类型,那么就可以通过unsafe.pointer转化
 */
func EqType()  {
	a := []MyInt{0, 1, 2}
	b := *(*[]int)(unsafe.Pointer(&a))
	b[2] = 9
	fmt.Printf("MyInt 转 int %+v \n",a)
}

type Tag struct {
	ID int
	Title string
	Pid int
}

type Cat struct {
	UUID int
	Name string
	PID int
	B float64
}

/**
 两个数据结构  字段按顺序( ID --> UUID   Tile-->Name Pid-->PID ) 底层结构的数据类型相同
 之后的字段  只能是 要么在 Tag 里面, 要么是在 Cat里面
 */
func EqStructLowType()  {
	tag := Tag{
		ID:    111,
		Title: "tag",
		Pid:   9,
	}
	cat := *(*Cat)(unsafe.Pointer(&tag))
	fmt.Printf("Tag 转 Cat 后  tag is: %+v, cat is: %+v \n",tag,cat)
	cat.PID = 123
	fmt.Printf("Tag 转 Cat 后 cat.PID = 123,tag is: %+v, cat is: %+v \n",tag,cat)
}

func Unsafe()  {
	EqType()
	EqStructLowType()
}
