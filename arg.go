// 引数の処理
package main

import (
	"flag"
	"fmt"
	"os"
)

func ParseArgs() (string, string, []string) {
	if os.Args[0] == "" {
		u := `
[usage]  You need exec command.
e.g. $ processlog echo "hoge"
`
		fmt.Fprintf(os.Stderr, u)
	}
	logDir := flag.String(
		"logdir", "", "Log output directory (default=stderr)")
	flag.Parse()
	return *logDir, flag.Arg(0), flag.Args()[1:]
}
