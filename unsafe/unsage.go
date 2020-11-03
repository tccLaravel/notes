package unsafe

import (
	"fmt"
	"unsafe"
)


type V struct {
	b string
	i int32
	j int64
}

func (v V) GetI() {
	fmt.Printf("i=%+v\n", v.i)
}
func (v V) GetJ() {
	fmt.Printf("j=%+v\n", v.j)
}

type Person struct {
	name   string
	age    int
	gender bool
}

func OffsetOf()  {
	var v *V = new(V)
	fmt.Printf("v size is = %d \n",unsafe.Sizeof(*v))
	p := unsafe.Pointer(v)
	ptr := unsafe.Pointer(uintptr(p))
	fmt.Println("v指针地址：", ptr)
	b := (*string)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(v.b)))
	fmt.Println("b指针地址：", b)
	//var i *int32 = (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + 4*unsafe.Sizeof(byte(0))))
	i := (*int32)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(v.i)))
	fmt.Println("i指针地址：", i)
	// 根据v的基准地址加上偏移量进行指针运算，运算后的值为j的地址，使用unsafe.Pointer转为指针
	j := (*int32)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(v.j)))
	fmt.Println("j指针地址：", j)
	*j = 123
	*i = 987
	*b = "tcc"
	fmt.Printf("v is %+v \n",v)

}

type Part1 struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}

func AlignOf()  {
	fmt.Println("******  AlignOf   ****** ")
	v := Part1{}
	fmt.Printf("part1 的对齐倍数是 %+v, 大小是 %+v \n",unsafe.Alignof(v),unsafe.Sizeof(v))
	fmt.Printf("a 的对齐倍数是 %+v, 偏移量是 %+v , 大小是 %+v \n",unsafe.Alignof(v.a),unsafe.Offsetof(v.a),unsafe.Sizeof(v.a))
	fmt.Printf("b 的对齐倍数是 %+v, 偏移量是 %+v , 大小是 %+v \n",unsafe.Alignof(v.b),unsafe.Offsetof(v.b),unsafe.Sizeof(v.b))
	fmt.Printf("c 的对齐倍数是 %+v, 偏移量是 %+v , 大小是 %+v \n",unsafe.Alignof(v.c),unsafe.Offsetof(v.c),unsafe.Sizeof(v.c))
	fmt.Printf("d 的对齐倍数是 %+v, 偏移量是 %+v , 大小是 %+v \n",unsafe.Alignof(v.d),unsafe.Offsetof(v.d),unsafe.Sizeof(v.d))
	fmt.Printf("e 的对齐倍数是 %+v, 偏移量是 %+v , 大小是 %+v \n",unsafe.Alignof(v.e),unsafe.Offsetof(v.e),unsafe.Sizeof(v.e))
	p := unsafe.Pointer(&v)
	ptr := unsafe.Pointer(uintptr(p))
	fmt.Println("v指针地址：", ptr)
	fmt.Printf("v is %+v \n",v)
}

func Unsafe()  {
	OffsetOf()
	//AlignOf()
}