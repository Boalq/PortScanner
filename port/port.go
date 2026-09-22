package port

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ScanResult struct {
	Port  string
	State string
	Banner string
}

// Learning: In go the type have the * instead of the variable!
func ScanPort(protocol, hostname string, port int, ch chan<- ScanResult, wg *sync.WaitGroup){
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}

	defer wg.Done()

	adress := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, adress, 500*time.Millisecond)

	if err != nil {
		result.State = "Closed"
		// Mistake: So basically I didn't returned a closed result with ch <- result therefore:
		// The Closed ports got lost
		ch <- result
		return
	}
	defer conn.Close()

	result.State = "Open"
	result.Banner = grabBanner(conn)
	ch <- result
}

func grabBanner(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	buf := make([]byte, 256)
	n, _ := conn.Read(buf)
	return strings.TrimSpace(string(buf[:n]))
}

func InitialScan(hostname, flag string) []ScanResult{
	var results []ScanResult
	var wg sync.WaitGroup

	ch := make(chan ScanResult)

	err := isValidHost(hostname)
	if err != nil{
		fmt.Println("Invalid Host: '", hostname, "'")
		return nil
	}

	for i := 1; i <= 2000; i++ {
		wg.Add(1)
		go ScanPort("tcp", hostname, i, ch, &wg)
		time.Sleep(2 * time.Millisecond)	// Needed buffer for not overasking the server
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		//fmt.Printf("State: '%s' | Flag: '%s'\n", result.State, flag)
		if result.State == "Open" && flag == "o"{
			results = append(results, result)
		} else if result.State == "Closed" && flag == "c" {
			results = append(results, result)
		} else if flag == "a" {
			results = append(results, result)
		}
	}

	return results
}

func isValidHost(hostname string) error {
	_, err := net.LookupHost(hostname)
	if err != nil {
		return err
	}
	return nil
}