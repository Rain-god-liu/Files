package main

import (
	"fmt"
	"time"
)

func main() {
	var m [1000] string
	var n [1000] int64
	localtime := "2006-01-02 15:04:05.000000001 +0800 CST"
	y := 0
	for i := 0; i < len(n); i++ {
		var x int64
		fmt.Scanf("%d", &x)
		if x == 0{
			break
		} else {
			n[i] = x
		}
		result := time.Unix(n[i], 0).Format(localtime)
		m[i] = result
		fmt.Println("input ok!")
		y++
	}
	fmt.Println("the result are:")
	for j := 0; j < y; j++ {
		fmt.Println(m[j])
	}
}
