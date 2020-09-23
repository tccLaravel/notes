package base

import (
	"fmt"
	"strconv"
)

func StrToInt(s string)  {
	i,_ := strconv.Atoi(s)
	fmt.Printf("%s strconv.atoi is %+v \n",s,i)
	j,_ := strconv.ParseFloat(s,64)
	fmt.Printf("%s strconv.ParseFloat v is %+v \n",s,j)
	fmt.Printf("%s strconv.ParseFloat int(j) is %+v \n",s,int(j))
	fmt.Printf("%s strconv.ParseFloat f is %f \n",s,j)
	m,_:= strconv.ParseInt(s,10,64)
	fmt.Printf("%s strconv.ParseInt is %+v \n",s,m)
}
