package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os"
)

func main() {
	//charset.NewReader()
	data, err := ioutil.ReadFile("/home/sws/thunder/orm/boxset.ini")
	fmt.Println(string(data), err)
	//fmt.Println(charset.DetermineEncoding(data, "utf8"))
	enc := simplifiedchinese.GB18030.NewDecoder()
	var dst []byte = make([]byte, len(data)*3)
	fmt.Println(enc.Transform(dst, data, false))
	fmt.Println("dst:", string(dst))

	f, err := ini.Load("/home/sws/thunder/orm/boxset.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sect := f.Section("Boxtype")
	boxname := sect.Key("boxname").String()
	boxvalue := sect.Key("boxvalue").String()
	fmt.Println("boxname:", boxname)
	fmt.Println("boxvalue:", boxvalue)

}
