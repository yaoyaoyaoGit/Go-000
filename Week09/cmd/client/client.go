package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

var port string

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
	c, err := net.Dial("tcp", ":"+port)
	if err != nil {
		return err
	}
	defer c.Close()
	for i := 0; i < 10; i++ {
		c.Write([]byte(fmt.Sprintf("message%d\n", i)))
		time.Sleep(1 * time.Second)
		// message, err := bufio.NewReader(c).ReadString('\n')
		// if err != nil {
		// 	return err
		// }
		// fmt.Print("->: " + message)
		// if strings.TrimSpace(string(text)) == "STOP" {
		// 	fmt.Println("TCP client exiting...")
		// 	return nil
		// }
	}
	for {
		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			return err
		}
		fmt.Print("->: " + message)
	}
	// if strings.TrimSpace(string(text)) == "STOP" {
	// 	fmt.Println("TCP client exiting...")
	// 	return nil
	// }
	return nil
}
