package container

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point   // 这里就是结构体的嵌入  把 point 结构体 嵌入到 circle 里面,以其中的一个成员形式存在
	Radius int
}

type CircleN struct {
	Point           //结构体的匿名成员  其实就是结构体的嵌入 省略了成员名
	Radius int
}

func Struct()  {
	var st1  Circle
	var st2  CircleN

	st1.Center.X = 3        /*区别*/
	st1.Center.Y = 2
	st1.Radius = 1

	st2.X = 3               /*区别*/
	st2.Y = 2
	st2.Radius = 1

	fmt.Printf("嵌入 Circle : %+v, 匿名 CircleN : %+v \n",st1,st2 )
}
