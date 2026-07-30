package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var port int
	var addr string

	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.StringVar(&addr, "address", "localhost", "--address localhost")
	flag.Parse()

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", addr, port)) //tcp 주소 객체로 바꾸기
	if err != nil {
		log.Fatal("resolveTCPAddr failed:", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr) // 바꾼 주소 객체로 연결
	if err != nil {
		log.Fatal("Dial failed:", err)
	}

	go func() {
		scan := bufio.NewScanner(conn)
		scan.Split(bufio.ScanLines)
		for scan.Scan() {
			fmt.Println(scan.Text())
		}
		if err := scan.Err(); err != nil {
			log.Println("read from connection error:", err)
		}
	}()

	for {
		inputScan := bufio.NewScanner(os.Stdin)
		inputScan.Split(bufio.ScanLines)
		for inputScan.Scan() {
			if inputScan.Text() == "exit" {
				return
			}
			conn.Write([]byte(fmt.Sprintf("%s\n", inputScan.Text())))
		}
		if err := inputScan.Err(); err != nil {
			log.Println("read from stdin error:", err)
			return
		}
	}
}
