package main

import (
	"PortScanner/scanner"
	"PortScanner/util"
	"fmt"
	"os"
	"runtime"
)

func main() {
	if len(os.Args) == 3 {
		ipList := os.Args[1]
		portList := os.Args[2]

		ips, err := util.GetIpList(ipList)
		ports, err := util.GetPorts(portList)
		_ = err
		//fmt.Println(ports)

		tasks, _ := scanner.GenerateTask(ips, ports)
		scanner.RunTask(tasks)
		scanner.PrintResult()
	} else {
		fmt.Printf("%v iplist port\n", os.Args[0])
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}