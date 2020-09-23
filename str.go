package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 是否存在某个字符或子串
func IsExist(s,substr string)  {
	var c bool = true
	//子串substr在s中，返回true
	c = strings.Contains(s,substr)
	fmt.Printf("fun Contains :%s Contains :%s res is: %v \n",s,substr,c)

	substr = "tcc"
	//subst 中任何一个字符 Unicode 代码点在s中，返回true
	c = strings.ContainsAny(s,substr)
	fmt.Printf("fun ContainsAny :%s ContainsAny :%s res is: %v \n",s,substr,c)

	substr = "tcc"
	//
	r := []rune(substr)
	c = strings.ContainsRune(s,r[1])
	fmt.Printf("fun ContainsRune :%s ContainsRune :%s res is: %v \n",s,substr,c)
}
//子串出现次数(字符串匹配)
func SubCount(s,substr string)  {
	c := strings.Count(s,substr)
	fmt.Printf("fun Count :%s Count :%s res is: %v \n",s,substr,c)
}

//字符串切割
func SplitStr(s string)  {
	//用一个或多个连续的空格分隔字符串 s，返回子字符串的数组（slice）
	f := strings.Fields(s)
	fmt.Printf("fun Fields :%s, res is: %q\n",s,f)

	e := strings.FieldsFunc(s,unicode.IsNumber)
	fmt.Printf("fun FieldsFunc :%s, res is: %q\n",s,e)

	// Split 和 SplitAfter、 SplitN 和 SplitAfterN
	// Split 与 SplitAfter 的区别: 分割字符串
	// Split 去掉分割的字符串
	s = "唐m tccm 1236m 绵绵"
	a := strings.Split(s,"m")
	fmt.Printf("fun Split :%s, res is: %q\n",s,a)
	// SplitAfter 分割字符串会被放入被分割的字字符串后面
	b := strings.SplitAfter(s, "m")
	fmt.Printf("fun SplitAfter :%s, res is: %q\n",s,b)

	// SplitN 控制了返回的数据个数
	c := strings.SplitN(s,"m",3)
	fmt.Printf("fun SplitN :%s, res is: %q\n",s,c)
	d := strings.SplitAfterN(s, "m",3)
	fmt.Printf("fun SplitAfter :%s, res is: %q\n",s,d)
}

//前缀.后缀
func HasPreOrSuf(s string)  {
	a := strings.HasPrefix(s,"th")
	fmt.Printf("fun HasPrefix :%s, res is: %v\n",s,a)
	b := strings.HasSuffix(s,"889")
	fmt.Printf("fun HasSuffix :%s, res is: %v\n",s,b)
}

//字符或子串在字符串中出现的位置
func StrIndex(s,sub string)  {
	// 在 s 中查找 sep 的第一次出现，返回第一次出现的索引
	a := strings.Index(s,sub)
	fmt.Printf("fun Index :%s, index of sub :%q res is: %v\n",s,sub,a)
	// chars中任何一个Unicode代码点在s中首次出现的位置
	b := strings.IndexAny(s,sub)
	fmt.Printf("fun IndexAny :%s, IndexAny of sub :%q res is: %v\n",s,sub,b)

	// 查找字符 c 在 s 中第一次出现的位置，其中 c 满足 f(c) 返回 true
	c := strings.IndexFunc(s, func(r rune) bool {
		if r == 's' {
			return  true
		}
		return  false
	})
	fmt.Printf("fun IndexFunc :%s, IndexFunc of sub :%q res is: %v\n",s,'s',c)

	d := strings.IndexRune(s,'m')
	fmt.Printf("fun IndexRune :%s, IndexFunc of sub :%q res is: %v\n",s,'m',d)

	e := strings.LastIndex(s,sub)
	fmt.Printf("fun LastIndex :%s, LastIndex of sub :%q res is: %v\n",s,sub,e)

	f := strings.LastIndexAny(s,sub)
	fmt.Printf("fun LastIndexAny :%s, LastIndexAny of sub :%q res is: %v\n",s,sub,f)

	g := strings.LastIndexByte(s,'t')
	fmt.Printf("fun LastIndexByte :%s, LastIndexByte of sub :%q res is: %v\n",s,sub,g)

	h := strings.LastIndexFunc(s, func(r rune) bool {
		if r == '8' {
			return  true
		}
		return  false
	})
	fmt.Printf("fun LastIndexFunc :%s, LastIndexFunc of sub :%q res is: %v\n",s,'8',h)
}

func Str()  {
	s := "i am is tcc,this is a string 10 , used to test, str is here is work, 789"
	substr := "str"
	IsExist(s,substr)
	SubCount(s,substr)
	SplitStr(s)
	HasPreOrSuf(s)
	StrIndex(s,"is")
	//将字符串数组（或slice）连接起来
	a := strings.Join([]string{"name=xxx", "age=xx"}, "&")
	fmt.Printf("fun Join %s \n",a)
	//把一个字符串重复N次
	b := strings.Repeat(substr,3)
	fmt.Printf("fun Repeat :%s Repeat 3 time res is : %v \n",substr,b)
	//字符串子串替换 n=>替换个数 如果 n < 0，则不限制替换次数，即全部替换
	c := strings.Replace(s,substr,"new",-1)
	fmt.Printf("fun Replace :%s Replace  res is : %v \n",substr,c)
}
