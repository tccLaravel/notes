package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigStruct struct {
	Host              string   `json:"host"`
	Port              int      `json:"port"`
	AnalyticsFile     string   `json:"analytics_file"`
	StaticFileVersion int      `json:"static_file_version"`
	StaticDir         string   `json:"static_dir"`
	TemplatesDir      string   `json:"templates_dir"`
	SerTcpSocketHost  string   `json:"serTcpSocketHost"`
	SerTcpSocketPort  int      `json:"serTcpSocketPort"`
	Fruits            []string `json:"fruits"`
}

type Other struct {
	SerTcpSocketHost string   `json:"serTcpSocketHost"`
	SerTcpSocketPort int      `json:"serTcpSocketPort"`
	Fruits           []string `json:"fruits"`
}

func Json()  {
	jsonStr := `{"host": "localhost:8080","port": 8080,"analytics_file": "go.log","static_file_version": 1,"static_dir": "E:/go/notes/src/","templates_dir": "E:/go/notes/src/templates/","serTcpSocketHost": ":159","serTcpSocketPort": 357,"fruits": ["apple", "banana"]}`

	//json str 转map
	var mp map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &mp); err == nil {
		fmt.Println("********** json str 转map json.Unmarshal([]byte(jsonStr), &map[string]interface{}) *********")
		fmt.Printf("mp  is: %+v \n,  and dat[`host`] is : %s \n",mp,mp["host"])
	}

	//json str 转struct
	var config ConfigStruct
	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
		fmt.Println("********** json str 转 struct json.Unmarshal([]byte(jsonStr), &ConfigStruct) *********")
		fmt.Printf("config  is: %+v \n,  and config.Host is : %s \n",config,config.Host)
	}

	//json str 转struct(部份字段)
	var part Other
	if err := json.Unmarshal([]byte(jsonStr), &part); err == nil {
		fmt.Println("********** json str 转 部分 struct json.Unmarshal([]byte(jsonStr), &Other) *********")
		fmt.Printf("part is: %+v \n,  and part.SerTcpSocketPort is : %+v \n",part,part.SerTcpSocketPort)
	}

	//struct 到json str
	if b, err := json.Marshal(config); err == nil {
		fmt.Println("********** struct 到 json str  json.Marshal(config)  *********")
		fmt.Printf("config struct is: %+v \n ,  and string(b) is : %s \n",config,string(b))
	}

	//map 到json str
	fmt.Println("********** map 到json str  json.NewEncoder(os.Stdout)  *********")
	enc := json.NewEncoder(os.Stdout)
	_ = enc.Encode(mp)

	//array 到 json str
	arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	lang, err := json.Marshal(arr)
	if err == nil {
		fmt.Println("************** array 到 json str json.Marshal(arr) ***************")
		fmt.Println(string(lang))
	}

	//json 到 []string
	var wo []string
	if err := json.Unmarshal(lang, &wo); err == nil {
		fmt.Println("*********** json 到 []string json.Unmarshal(lang, &wo) **************")
		fmt.Println(wo)
	}
}
