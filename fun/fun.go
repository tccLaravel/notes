package fun

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func adder() func(int) int  {
	sum := 0 //自由变量
	return func(i int) int {
		// i 局部变量
		sum += i
		return  sum
	}
}


func fibonacci() intGen {
	a ,b := 0,1
	return func() int {
		a, b = b,a+b
		return b
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n",next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf(scanner.Text())
	}
}

func ClosureFun()  {
	add := adder()
	f := fibonacci()
	for i := 0; i< 10; i++ {
		fmt.Printf("adder 0 + 1 + ... +%d  = %d \n",i,add(i))
		//fmt.Printf("fibonacci i %d = %d \n",i,f())
	}
	printFileContents(f)
}
