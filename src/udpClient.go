package main

import (
	"fmt"
	"net"
	"test/checkErr"
	"time"
)

func main() {
	udpClient, err := net.ResolveUDPAddr("udp", "192.168.1.97:3343")
	checkErr.CheckErr(err, true)
	fmt.Println("udpClient is:", udpClient)

	udpConn, err := net.DialUDP("udp", nil, udpClient)
	checkErr.CheckErr(err, true)
	defer udpConn.Close()

	for i := 0; i < 1; i++ {
		var buf = []byte("songweishuai")
		len, err := udpConn.Write(buf)
		fmt.Println(len, err)
		time.Sleep(time.Second * 3)
	}

}
