package main

import"fmt"

var i,j int

func calculate(i int){
	for j = 2;j < i;j++{
		if i % j == 0{
			break
		}
	}
	if i == j{
		fmt.Println(i)
	}
}

func main(){
	for i = 1;i <= 10000;i++{
		go calculate(i)
	}
}