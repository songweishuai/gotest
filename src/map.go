package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//func main() {
//	var countryCapitionalMap map[string]string = make(map[string]string)
//
//	countryCapitionalMap["France"] = "paris"
//	countryCapitionalMap["Italy"] = "Rome"
//	countryCapitionalMap["Japan"] = "Tokyo"
//	countryCapitionalMap["India"] = "new Delhi"
//
//	fmt.Println(countryCapitionalMap["asd"])
//
//	for country := range countryCapitionalMap {
//		fmt.Println("Capital of", country, "is", countryCapitionalMap[country])
//	}
//
//	captial, ok := countryCapitionalMap["china"]
//	if ok {
//		fmt.Println("Captial of china is ", captial)
//	} else {
//		fmt.Println("Captial of China is not present")
//	}
//
//	delete(countryCapitionalMap, "France")
//
//	fmt.Println("*************")
//	for country := range countryCapitionalMap {
//		fmt.Println("Capital of", country, "is", countryCapitionalMap[country])
//	}
//}

func main() {
	var m = map[string]string{
		"song": "weishuai",
	}

	info := make([]map[string]string, 1)

	info = append(info, m)

	data, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
