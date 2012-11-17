package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func script(s []string, slen int) {
	scr, err := os.Open(s[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	r := bufio.NewReader(scr)
	for {
		str, err := r.ReadString('\n')
		if err != nil {
			return
		}
		str = strings.Trim(str, "\n\r ")
		if str != "" {
			s,s1 := parser(str)
			commands(s, len(s),s1)
		}
	}
}
