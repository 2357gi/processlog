package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	logDir, commandName, args := ParseArgs()
	stdout, stderr, err := initout(logDir, commandName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant create file: %v\n", err)
		os.Exit(1)
	}

	startTime := time.Now()
	ps, err := execution(commandName, args, stdout, stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant exec command %s", commandName)
		os.Exit(1)

	}

	exitTime := time.Now()
	fmt.Fprintln(stdout, ps.String())
	fmt.Fprintf(stdout, "Wall clock time=%v system time=%v user time=%v\n",
		exitTime.Sub(startTime), ps.SystemTime(), ps.UserTime())

}
