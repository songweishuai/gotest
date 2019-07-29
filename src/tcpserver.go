package main

import (
	"fmt"
	"net"
	"test/checkErr"
	"time"
)

func main() {
	tcpServerAddr, err := net.ResolveTCPAddr("tcp", "192.168.1.97:3342")
	checkErr.CheckErr(err, true)

	listenter, err := net.ListenTCP("tcp", tcpServerAddr)
	checkErr.CheckErr(err, true)

	fmt.Printf("listenter.Addr() = %v\n", listenter.Addr())

	for {
		con, err := listenter.AcceptTCP()
		checkErr.CheckErr(err, true)

		go HandleFun(con)
	}
}

func HandleFun(con *net.TCPConn) {
	if con == nil {
		return
	}
	defer con.Close()

	fmt.Printf("con.LocalAddr() = %v\n", con.LocalAddr())
	fmt.Printf("con.RemoteAddr() = %v\n", con.RemoteAddr())

	con.SetDeadline(time.Now().Add(time.Second))

	for {
		buf := make([]byte, 64)
		len, err := con.Read(buf)
		if err != nil {
			checkErr.CheckErr(err, false)
			break
		}
		fmt.Printf("len = %v\n", len)
		fmt.Printf("recv buf is:%s", buf)
		//time.Sleep(time.Second * 2)
	}
}
