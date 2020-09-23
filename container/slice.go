package container

import (
	"fmt"
)

func Seasons()  {
	months := [...]string{1: "January",2:"February",3:"March",4:"April",5:"May",6:"June",7:"July",8:"August",9:"September",10:"October",11:"November", 12: "December"}
	one := months[1:4]
	two := months[4:7]
	three := months[7:10]
	four := months[10:13]
	fmt.Println("s[i:j] 是从第 i 个元素开始到 j-1个元素的子序列")
	fmt.Printf("four seasons is : %+v \n " +
		"one months[1:4] : %+v ,len is: %+v, cap is: %+v \n " +
		"two months[4:7] : %+v ,len is: %+v, cap is: %+v \n " +
		"three months[7:10] : %+v ,len is: %+v, cap is: %+v \n " +
		"four months[10:13] : %+v ,len is: %+v, cap is: %+v \n",
		months,one,len(one),cap(one),two,len(two),cap(two),three,len(three),cap(three),four,len(four),cap(four))
	fmt.Println("容量一般是从slice的开始位置到底层数据的结尾位置,看上面的四个季度的cap值就能发现这一点")
}

func AppendBaseArray()  {
	a := [3]int{}
	fmt.Printf("a array init [3]int{} is %+v",a)
	b := a[:2]
	b1 := append(b,1,2,3)
	fmt.Printf(", b = a[:2] is: %+v, b1 = append(b,1,2,3) is : %+v , last a is %+v \n",b,b1,a)
	fmt.Println("特别注意这里的 a 的值 没有变更,因为当b 追加了 3个元素后,len=5  > 数组的 len = 3 了,已经超出了数组a本身的长度,这样在底层会重新创建一个新的数组,在被引用")
	fmt.Printf("如果上面改为 b1 = append(b,1) 那么长度没变 这样a 的值 就是 %+v \n",append(a[:2],1))
}

func AppendBaseSlice()  {
	var s = make([]int,3)
	s1 := append(s,1,2,3)
	fmt.Printf("s slice init make([]int,3) is %+v s1 = append(s,1,2,3) is: %+v \n",s,s1)
}

func Append()  {
	fmt.Println("AppendBaseArray")
	AppendBaseArray()
	fmt.Println("AppendBaseSlice")
	AppendBaseSlice()
}

func Copy()  {
	s1 := []int{3,6,9}
	s2 := []int{1,4}
	copy(s1,s2)
	fmt.Printf("s1 is %+v \n",s1)
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
}

func SliceInit()  {
	var s1 = make([]int,0)
	var s3 []int
	fmt.Printf("var s1 = make([]int,0) is: %+v, len(s1) is: %+v ,cap(s1) is %+v ,s1 == nil is:  %+v \n",s1,len(s1),cap(s1),(s1==nil))
	fmt.Printf("var s3 []int is: %+v, len(s3) is: %+v ,cap(s3) is %+v ,s3 == nil is: %+v\n",s3,len(s3),cap(s3),(s3 == nil))
	fmt.Println(append(s1,1),len(s1),cap(s1))
	fmt.Println(append(s3,3),len(s3),cap(s3))
}

func remove(slice []int, i int) []int {
	fmt.Printf("slice[i:] is: %+v, slice[i+1:] is %+v \n",slice[i:],slice[i+1:])
	copy(slice[i:], slice[i+1:])
	fmt.Printf("slice is: %+v, \n",slice)
	return slice[:len(slice)-1]
}

