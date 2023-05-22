package main

import (
	"runtime"
	"user-svc/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
