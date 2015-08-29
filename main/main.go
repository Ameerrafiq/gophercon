package main

import (
	"degreespkg"
	"runtime"
)

func main() {
	//TO maximize the utilization of cpu process
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS((numCPU * 3) / 4)
	//Starting server
	degreespkg.StartServer()
}
