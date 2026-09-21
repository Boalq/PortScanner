package port

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type ScanResult struct {
	Port  string
	State string
	Banner string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}

	adress := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, adress, 500*time.Millisecond)

	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()

	result.State = "Open"
	result.Banner = grabBanner(conn)
	return result
}

func grabBanner(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	buf := make([]byte, 256)
	n, _ := conn.Read(buf)
	return strings.TrimSpace(string(buf[:n]))
}

func InitialScan(hostname, flag string, ch chan<-[]ScanResult){
	var results []ScanResult
	defer close(ch)

	err := isValidHost(hostname)
	if err != nil{
		fmt.Println("Invalid Host: '", hostname, "'")
		return 
	}

	for i := 1; i <= 1024; i++ {
		result := ScanPort("tcp", hostname, i)

		if result.State == "Open" && flag == "o"{
			results = append(results, result)
		} else if result.State == "Closed" && flag == "c" {
			results = append(results, result)
		} else if flag == "a" {
			results = append(results, result)
		}
	}

	ch <- results
}

func isValidHost(hostname string) error {
	_, err := net.LookupHost(hostname)
	return err
}