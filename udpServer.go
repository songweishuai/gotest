package main

import (
	"fmt"
	"net"
	"test/checkErr"
	)

func main() {
	udpServer, err := net.ResolveUDPAddr("udp", "192.168.1.97:3343")
	checkErr.CheckErr(err, true)

	udpConn, err := net.ListenUDP("udp", udpServer)
	checkErr.CheckErr(err, true)
	defer udpConn.Close()

	fmt.Println("udpConn.LocalAddr():", udpConn.LocalAddr())
	fmt.Println("udpConn.RemoteAddr():", udpConn.RemoteAddr())

	//time out set
	//udpConn.SetDeadline(time.Now().Add(time.Second))

	var message string
	for {
		var buf []byte = []byte(" ")
		//len, addr, err := udpConn.ReadFrom(buf)
		//fmt.Println("add is:",addr)
		//checkErr.CheckErr(err, false)
		//fmt.Println("add is:", addr, ",len is:", len,"buf is:",buf)
		len, err := udpConn.Read(buf)
		checkErr.CheckErr(err, false)
		fmt.Println(len, err)
		fmt.Printf("%s\n", buf)
		message=message+string(buf)
		fmt.Println(message)
	}
}
