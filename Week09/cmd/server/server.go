package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var port string

var count int

func main() {

	port = "8080"

	args := os.Args
	if len(args) == 2 {
		port = args[1]
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

}

func run() error {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	defer l.Close()
	fmt.Println("Listening on: ", port)

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		fmt.Printf("Received message from: %s", conn.RemoteAddr())

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		// 测试粘包问题
		// data := make([]byte, 20)
		// n, err := bufio.NewReader(conn).Read(data)

		//用换行符解决粘包的问题，依然无法读取所有的包，如果发送非常快的话
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		// 为什么这种方法没法读取所有的包，
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("CLOSE CONN")
			return
		}
		// fmt.Println("-> ", string(data), n)
		fmt.Println("-> ", string(data))
		count++
		conn.Write([]byte(fmt.Sprintf("This is %d msg\n", count)))
	}
}
