package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var prompt string
var path string

var infunc bool
var fname string

func main() {
	stdin := bufio.NewReader(os.Stdin)
	prompt = "% "
	for {
		fmt.Print(prompt)
		str, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		str = strings.Trim(str, "\n\r ")
		if str != "" {
			s, f := parser(str)
			commands(s, len(s), f)
		}
	}
}

func commands(s []string, slen int, f string) {
	cmd := s[0]
	slen = slen - 1
	args := s
	if !infunc {
		switch cmd {
		case "cd":
			chdir(args[1:], slen)
		case "version":
			version()
		case "exit":
			os.Exit(0)
		case "set":
			setvar(args[1:], slen)
		case "go":
			go commands(args[1:], slen,f)
		case "@":
			script(args[1:], slen)
		case "run":
			script(args[1:], slen)
		default:
			execute(cmd, args, slen)
		}
	}
}

func chdir(args []string, slen int) {
	err := os.Chdir(args[0])
	if err != nil {
		fmt.Println(err)
	}
}

func version() {
	fmt.Println("go userland shell.")
}

func setvar(s []string, slen int) {
	if slen < 2 {
		fmt.Println("set [parameter] [value]")
		return
	}
	para := s[0]
	val := s[1]
	switch para {
	case "prompt":
		prompt = val
	case "path":
		path = val
	default:
		fmt.Println("Unknown parameter:", para)
	}
}

func execute(cmd string, args []string, slen int) {
	var procattr os.ProcAttr
	procattr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	if path != "" {
		cmd = path + cmd
	}
	proc, err := os.StartProcess(cmd, args, &procattr)
	if err != nil {
		fmt.Println(err)
		return
	}
	proc.Wait()
	return
}
