package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("/home/sws/thunder/bin/dhcp/Config.ini")
	if err != nil {
		os.Exit(1)
	}

	ts := simplifiedchinese.GB18030.NewDecoder()
	dataUtf8 := make([]byte, len(data)*2)
	_, _, err = ts.Transform(dataUtf8, data, false)
	if err != nil {
		os.Exit(2)
	}

	info := make([]map[string]string, 20)
LOOP:
	for true {
		n, data, _ := bufio.ScanLines(dataUtf8, false)
		if n <= 0 {
			break LOOP
		}

		if string(data) != "[ITEM]" {
			dataUtf8 = dataUtf8[n:]
			continue
		}

		m := make(map[string]string)
		dataUtf8 = dataUtf8[n:]
		for true {
			n, data, _ := bufio.ScanLines(dataUtf8, false)
			if n <= 0 {
				break LOOP
			}

			if string(data) == "[ITEM]" {
				break
			}

			if !bytes.ContainsAny(data, "=") {
				m["select"] = string(data)
			} else {
				d := bytes.Split(data, []byte("="))
				if len(d) == 2 {
					m[string(d[0])] = string(d[1])
				}
			}
			info = append(info, m)
			dataUtf8 = dataUtf8[n:]
		}
	}

	dataJS,_:= json.Marshal(info)
	fmt.Println(string(dataJS))
}
