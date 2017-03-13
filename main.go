package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"time"
)

var addr = flag.String("addr", "127.0.0.1:1818", "remote ip address + port")

func main() {
	flag.Parse()
	for {
		hostname, err := getRemoteHostNameThroughTCP(addr)

		if err != nil && err != io.EOF {
			fmt.Println(err)
		}

		fmt.Println(time.Now(), ":", hostname)
		time.Sleep(500 * time.Millisecond)
	}
}

func getRemoteHostNameThroughTCP(addr *string) (string, error) {
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	return bufio.NewReader(conn).ReadString('\n')
}

func testTCPConnection(addr *string) {
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		fmt.Println(err)
	}

	in := bufio.NewReader(conn)
	var line string

	for {
		line, err = in.ReadString('\n')
		fmt.Println(line)
		if err != nil {
			break
		}
	}

	err = conn.Close()
	if err != nil {
		fmt.Println(err)
	}
}
