package base

//var 变量名字 类型 = 表达式    var name type = expression
/**
var param string = ""          var name type = expression
var param1 string              var name type
var param2 = ""                var name  = expression
 param3 := ""                  name := expression
 p := new(int) //p *int 类型, 指向匿名的 int 变量
 */

//基础的数据类型
/**
//int uint 这是与cpu平台机器字大小一致
整型 (int uint uint8 uint16 uint32 uint64 int16 int32  int64)

float (float32 float64)

boolean

//一个字符串是一个不可改变的字节序列,字符串可以包含任意的数据，包括byte值0
文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列  type rune int32
string

//(boolean string 整型)
const
//复数
complex
*/

/**
基础数据类型的基础转换method
strconv.ParseInt
strconv.ParseUint
strconv.ParseBool
strconv.ParseFloat
 */

//复杂数据类型
//数组 array 定长 定类型
var a1 [3]int    // [0,0,0]
var a2 = [3]int{3,5,7}
var a3 = [...]int{3,5,7}

//引用类型 (ptr slice map fun  channel)
//slice 切片  变长 定类型
/**
它的底层其实是一个数组,slice本身只是对底层数组的引用; 由  指针prt , 长度len, 容量cap 三部分组成
指针ptr就是数组的引用
长度len就是本身所拥有的数据个数
容量cap 一般是从slice的开始位置到底层数据的结尾位置 ,若果存储的数据超过当前cap,那么go会以 2*cap 扩容,重新分配内存空间;
append copy
 */
var s1 = make([]int,3) //type  len  此时 len = cap = 3  在底层，make创建了一个匿名的数组变量
var s2 = make([]int,3,5) // type len cap 此时 len = 3 cap = 5
var s3 []int

//map
/**
delete 通过key删除某个元素
 */
var m = make(map[interface{}]interface{})
var m1 = make(map[string]int)
var m2 = map[string]int{}
var m3 = map[string]int{"age":18,"sex":1}

//struct 结构体
/**
	需要注意的就是 结构体嵌入和匿名成员
 */
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       []string
	Position  string
	Salary    int
	ManagerID int
}
//注意一点  所有的变量定义都遵循 var name type 这种格式,  在这里  Employee  就是一种type
var tcc Employee
//结构体嵌入和匿名成员
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

//json
/**
json.Marshal   //json 格式化 为[]byte类型
json.Unmarshal()  //json 解码
 */

//特殊类型
/**
[]byte  uint8类型
rune    UTF8编码的Unicode码点
 */

/**
reflect 包
实际开发过程中,有很多参数是不知道传入的类型的,这样就需要先判定参数类型,然后在转化成我们需要的类型

unsafe包
 */

func BaseType()  {
	IntToInt(159)
	IntToFloat32(123)
	IntToFloat64(123)
	Float32ToInt(321.369)
	Float64ToInt(321.369)
	StrToInt("288")
}