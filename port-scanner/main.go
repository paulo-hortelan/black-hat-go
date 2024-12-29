package main

import (
	"fmt"
	"net"
	"time"
)

func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("app.aonet.com.br:%d", p)
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	const numPorts = 65000
	ports := make(chan int, 1000)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= numPorts; i++ {
			ports <- i
		}
		close(ports)
	}()

	for i := 0; i < numPorts; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	fmt.Println("Open ports:", openports)

	close(results)

	// for i := 1; i <= 15; i++ { // Adjust range based on the number of ports scanned
	// 	port := <-ports
	// 	fmt.Printf("Port %d is open\n", port)
	// }

	// go func() {
	// 	for i := 1; i <= 10; i++ {
	// 		ports <- i
	// 	}
	// 	close(ports)
	// }()

	// for i := 1; i <= 10; i++ { // Adjust range based on the number of ports scanned
	// 	result := <-results
	// 	fmt.Printf("Port %d is open\n", result)
	// }

	// for i := 0; i < 100; i++ {
	// 	port := <-results
	// 	if port != 0 {
	// 		openports = append(openports, port)
	// 	}
	// }

	// close(ports)
	// close(results)
	// sort.Ints(openports)
	// for _, port := range openports {
	// 	fmt.Printf("%d open\n", port)
	// }
}
