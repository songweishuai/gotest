package main

import (
	"flag"
	"os"
	"test/checkErr"
	"fmt"
)

var name = flag.String("n", "", "please input filename")

func main() {
	flag.Parse()
	if *name == "" {
		os.Exit(0)
	}

	fileInfo, err := os.Lstat(*name)
	checkErr.CheckErr(err,true)

	switch mode:=fileInfo.Mode() ;{
	case mode.IsDir():
		fmt.Printf("%+v\n",mode)
		fmt.Println(*name,"is dir")
	case mode.IsRegular():
		fmt.Printf("%+v\n",mode)
		fmt.Println(*name,"is regular file")
	default:
		fmt.Printf("%+v\n",mode)
	}
}
