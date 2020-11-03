package main

import (
	"bufio"
	"fmt"
	"log"
	"notes/container"
	"notes/fun"
	"notes/unsafe"
	"os"
)



func main() {
	unsafe.Unsafe()
	fmt.Println("Let`s go to learn go")
	fmt.Println("请选择你想学习的模块")
	fmt.Printf("base  container  fun \n")
	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		switch command {
		case "base":

		case "container":
			container.VarNewMake()
			container.PointerType()
		case "fun":
			fun.ClosureFun()
		}
		if command == "stop" {
			running = false
		}
		log.Println("command", command)
	}
}
