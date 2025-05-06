package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main(){
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide the port number you want to connect!")
		return
	}

	PORT:=":"+arguments[1]

	conn,err:=net.Dial("tcp","localhost"+PORT)
	if err!=nil{
		fmt.Println("error: ",err)
		return
	}
	defer conn.Close()

	response,err:=bufio.NewReader(conn).ReadString('\n')
	if err!=nil{
		fmt.Println("error: ",err)
		return
	}

	fmt.Println("Got Response: ",response)
}