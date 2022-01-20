/*
 * @Author: Jiwei
 * @Date: 2022-01-14 11:46:09
 * @LastEditTime: 2022-01-19 15:51:47
 * @LastEditors: Jiwei
 * @Description:
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch02\main.go
 *
 */
package main

import (
	"fmt"
	"jvmgo/ch02/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {

	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Cound not find or load main class")
		return
	}
	fmt.Printf("class data:%v\n", classData)

}
