package main

import (
	"fmt"
	"net"
	"test/checkErr"
)

func main() {
	//tcpAddr, err := net.ResolveTCPAddr("tcp", "192.168.1.97:4322")
	//tcpServerAddr, err := net.ResolveTCPAddr("tcp", "192.168.1.97:3342")
	tcpServerAddr, err := net.ResolveTCPAddr("tcp", "192.168.1.97:3342")
	checkErr.CheckErr(err, true)

	fmt.Println(tcpServerAddr)
	fmt.Printf("tcpAddr.Network() = %+v\n", tcpServerAddr.Network())
	fmt.Printf("tcpAddr.Zone = %v\n", tcpServerAddr.Zone)
	fmt.Printf("tcpAddr.String() = %v\n", tcpServerAddr.String())
	fmt.Println(net.LookupPort("tcp","127.0.0.1:1234"))

	connect, err := net.DialTCP("tcp", nil, tcpServerAddr)
	//connect, err := net.Dial("tcp", "192.168.1.97:3342")
	checkErr.CheckErr(err, true)
	defer connect.Close()

	fmt.Println("connect.LocalAddr()", connect.LocalAddr())

	//time.Sleep(time.Second*5)
	buf := []byte("100")
	fmt.Printf("buf is:%s\n", buf)
	connect.Write(buf)

	buf = []byte("song wei shuai\n")
	fmt.Printf("buf is:%s\n", buf)
	connect.Write(buf)

}
