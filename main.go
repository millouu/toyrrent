package main

import (
	"fmt"
	"log"
	"net"
	"flag"
)

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	packet := make([]byte, 4096)
	defer c.Close()
	packet = []byte("Hello there\n")
	num, _ := c.Write(packet)
	fmt.Printf("Wrote back %d bytes, the payload is %s\n", num, string(packet))
}

func main() {
	RECEIVE_CONN_PORT := flag.String("port","8080","Port to start the server")
	CONNECT_TO_PORT := flag.String("connect-port","8081","Port to connect to other server")
	flag.Parse()

	*RECEIVE_CONN_PORT = ":" + *RECEIVE_CONN_PORT
	*CONNECT_TO_PORT = ":" + *CONNECT_TO_PORT

	l, err := net.Listen("tcp4", *RECEIVE_CONN_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	
	fmt.Println("Server Started on Port ",*RECEIVE_CONN_PORT)
	
	// Try connecting to a server
	conn,err:=net.Dial("tcp","localhost" + *CONNECT_TO_PORT)
	if err!=nil{
		fmt.Println("Failed to connect to external server: ",err)
		return
	}

	defer conn.Close()

	for {

		// Download Pipe
        fmt.Println("Awaiting Connection")
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		handleConnection(c)
        fmt.Println("Served Connection")
	}
}
