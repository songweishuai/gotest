package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//os.Getpid()
	ProcessId := os.Getppid()
	process := fmt.Sprintf("/proc/%d/stat", ProcessId)

	info, err := ioutil.ReadFile(process)
	if err != nil {
		fmt.Println(err.Error())
	}

	InfoString := string(info)
	fmt.Println(InfoString)

	InfoStringSplit := strings.Split(InfoString, " ")

	for Index, v := range InfoStringSplit {
		fmt.Println("index:", Index)
		fmt.Println("v:", v)
	}
}
