/*
 * @Author: Jiwei
 * @Date: 2021-07-19 09:22:46
 * @LastEditTime: 2022-01-10 09:56:06
 * @LastEditors: Jiwei
 * @Description:
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch01\cmd.go
 *
 */
package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version message")
	flag.StringVar(&cmd.cpOption, "classpath", "", "input classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "input classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args)
}
