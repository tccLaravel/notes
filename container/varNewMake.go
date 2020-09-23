package container

import (
	"fmt"
)

type Pointer struct {
	ID int
	Tag []string
	Cat map[string]string
	C chan int
}
/**
slice  map chan  都是引用类型,所以初始化pointer后,其实都只是申明了变量,还没有为变量申请到指针对应的内存空间
必须对其 make() 后才能使用,
 */
func PointerType()  {
	p := Pointer{}
	fmt.Printf("Pointer的零值是: %+v \n",p)
	if p.Tag == nil {
		fmt.Printf("p.Tag is nil \n")
	}
	if p.Cat == nil {
		fmt.Printf("p.Cat is nil \n")
	}
	if p.C == nil {
		fmt.Printf("p.C is nil \n")
	}
	p.Tag = append(p.Tag,"tcc")
	fmt.Printf("slice 本身增加元素的时候会检索底层素组内存空间,如果没有会 make() \n")
	fmt.Printf("其实 通过 append 操作后, p.Tag 已经是另一个变量了")
	fmt.Printf("Tag 增加 `tcc` : %+v \n",p)
	p.Cat = make(map[string]string)
	fmt.Printf("p.Cat 如果没有这里make() 申请内存空间 会 deallock \n")
	p.Cat["fruit"] = "apple"
	fmt.Printf("Cat 增加 `apple` : %+v \n",p)

	p.C = make(chan int)
	fmt.Printf("p.C 如果没有这里make() 申请内存空间 会 deallock \n")
	go func() {
		for  {
			r := <- p.C
			fmt.Printf("received from p.C is : %+v \n",r)
		}
	}()
	p.C <- 10
}

func VarNewMake()  {
	fmt.Printf("变量的默认值 都是该变量的类型的 '零值' \n")
	var param1 string
	fmt.Printf("string 类型的 零值是 %+v \n",param1)
	var param2 int
	fmt.Printf("int 类型的 零值是 %+v \n",param2)
	var param3 *int
	fmt.Printf("*int 类型的 零值是,param3 =  %+v \n ",param3)
	//fmt.Printf("*int 类型的 零值是,*param3 =  %+v \n ",*param3)
	fmt.Printf("使用 *param3 = 10 这里会直接报 内存 指针类型错误 invalid memory address or nil pointer dereference \n")
	var param4 *int
	param4 = new(int) //通过new给分配内存
	fmt.Printf("new *int 类型的  零值是,param4 = %+v \n, ",param4)
	fmt.Printf("new *int 类型的  零值是,*param4 = %+v \n, ",*param4)
	fmt.Printf("new 方法,会生成type类型的 '零值', 它返回的永远是类型的指针，指向分配类型的内存地址。\n")
	fmt.Printf("如果 type 是复杂类型的,那么 new 会把生成复杂类型的基本类型的 '零值' ")

	*param4 = 10
	fmt.Printf("但是 这是一个引用类型 ,我们不紧要声明它，还要为它分配内容空间,不然这个变量是不能被使用的 \n")

	fmt.Printf("基本类型(int,float,string,bool)的变量的默认值 都是该变量的类型的 '零值'(0,0,'',false) \n")
	fmt.Printf("引用类型(*int,*string ...)变量的默认值的'零值'都是 nil,如果没有被新分配内存 是无法被使用的 \n")
	fmt.Printf("make 方法 只用于 chan,map,slice 的初始化,由于这三个类型本身就是引用类型,就直接返回类型本身了,没必要返回指针类型 \n")
	param7 := make([]int,0)
	fmt.Printf("make([]int) is: %+v \n ",param7)

	var param8 []int
	a := append(param8,3)
	fmt.Printf("var param8 []int is: %v \n ",param8)
	fmt.Printf("a is: %v \n ",a)
}
