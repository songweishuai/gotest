package main

import (
	"fmt"
)

func main() {
	totalMoney := 10000.0
	baseMoney := 833.33
	mounthDay := 30.0
	rate := 0.000318
	money := 0.0
	nLoop := 0
	totalRate := 0.0
	for totalMoney > 1 {
		nLoop++
		money = totalMoney * rate * mounthDay
		totalMoney -= baseMoney
		totalRate += money
		fmt.Println("第", nLoop, "个月")
		fmt.Println("利率:", money)
		fmt.Println("剩余金额:", totalMoney)
		fmt.Println("本月应换:", baseMoney+money)
		//time.Sleep(time.Second)
	}

	fmt.Println(totalRate)
}
