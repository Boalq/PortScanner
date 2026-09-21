package main

import (
	"flag"
	"fmt"
	"port-scanner/port"
)

func main() {
	var search_flag string;

	flag.StringVar(&search_flag, "search", "a", "Usage: ./program -search=... \n\t\t\t a = for all ports \n\t\t\t o = for open ports \n\t\t\t c = for closed ports")
	flag.Parse()

	fmt.Println("Boalq's Port Scanner in Go:");

	var hostname string

	fmt.Println("Enter your hostname: ")
	fmt.Scanln(&hostname)
	fmt.Printf("\nPorts:\n")

	ch := make(chan []port.ScanResult, 1)

	go port.InitialScan(hostname, search_flag, ch)
	defer close(ch)

	result := <- ch

	fmt.Println(result)
}