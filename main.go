package main

import (
	"fmt"
	"port-scanner/port"
	"os"
)

func main() {
	flag := ""

	if(len(os.Args) != 2){
		fmt.Printf("Usage:\n program... \n\t -a = Scan all ports \n\t -o = Scan all open Ports \n\t -c = Scan all closed Ports")
		return
	}

	flag = os.Args[1]

	fmt.Println("Boalq's Port Scanner in Go:");

	var hostname string

	fmt.Println("Enter your hostname: ")
	fmt.Scanln(&hostname)
	fmt.Printf("\nPorts:\n")

	results := port.InitialScan(hostname, flag)
	fmt.Println(results)
}