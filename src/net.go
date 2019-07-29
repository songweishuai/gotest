package main

import (
	"fmt"
	"net"
	"os"
	"test/checkErr"
)

func main() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

	addrs1, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//fmt.Println(addrs1)
	for index, address := range addrs1 {
		fmt.Println(index, address)
	}


	str, err := net.LookupAddr("127.0.1.1")
	checkErr.CheckErr(err, true)
	fmt.Println("str is:", str)

	host, err := net.LookupHost("www.baidu.com")
	checkErr.CheckErr(err, true)
	fmt.Println("host is:", host)

	myHostName, err := net.LookupHost("127.0.0.1")
	checkErr.CheckErr(err, true)
	fmt.Println("myHostName is:", myHostName)

	ip, err := net.LookupIP("www.baidu.com")
	checkErr.CheckErr(err, true)
	fmt.Println("ip is:", ip)

	canName,err:=net.LookupCNAME("www.baidu.com")
	checkErr.CheckErr(err,true)
	fmt.Println("baidu's cname is:",canName)

	nsName,err:=net.LookupNS("www.baidu.com")
	checkErr.CheckErr(err,false)
	fmt.Println("baidu's ns is:",nsName)

	mxName,err:=net.LookupMX("www.baidu.com")
	checkErr.CheckErr(err,false)
	fmt.Println("mxName is:",mxName)
}
