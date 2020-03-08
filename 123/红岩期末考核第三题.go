//红岩期末考核第三问
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type character struct {
	blood int
	lanliang int
	attack float32
}

var x  int
var A character
var B character

func A_taowa(A *character,B *character) {
	for v:=0;v<=100 ;v++ {
		//time.Sleep(time.Duration(1)*time.Second)
		rand.Seed(time.Now().UnixNano())
		x = rand.Intn(1-0+1)+0 //生成0-1随机整数
		if x > 0 {
			B.blood -= 10
			A.lanliang += 10
			if B.blood<=0 {
				goto Loop
			}
			fmt.Printf("A获得了连击特效，当前B的血量为%d/n",B.blood)
			A_taowa(A, B)
		} else{
			goto Loop
		}
	}
	Loop:
}

func A_attack(A *character,B *character){
	fmt.Println("现在是A的回合")
	for i:=0;i<=4;i++ {
		B.blood-=10
		A.lanliang+=10
	    A_taowa(A,B)
		if B.blood<=0 {
			fmt.Println("B的血量为0了")
			break
		}
		fmt.Println("B的血量 B的攻击 A的血量 A的攻击 A的蓝量")
		fmt.Println(B.blood,B.attack,A.blood,A.attack,A.lanliang)
	}
}

func B_attack(A *character,B *character){
	fmt.Println("现在是B的回合")
	for i:=0;i<=4;i++{
		A.blood-=10
		B.lanliang+=10
		if A.blood<=0 {
			fmt.Println("A的血量为0了")
			break
		}
		fmt.Println("B的血量 B的攻击 A的血量 A的攻击 A的蓝量")
		fmt.Println(B.blood,B.attack,A.blood,A.attack,A.lanliang)
	}
	if B.lanliang>=50 {
		A.attack=A.attack*(0.9)
	}
}



func main(){
	var count float64
	count=0
	var countA float64
	var countB float64
	countA=0
	countB=0
	var shenglv_A  float64
	var shenglv_B  float64
	shenglv_A=0
	shenglv_B=0
	for s:=0;s<=100 ;s++ {
		A = character{
			blood:    100,
			lanliang: 0,
			attack:   10,
		}

		B = character{
			blood:    300,
			lanliang: 0,
			attack:   20,
		}
		for i := 0; i <= 9; i++ {
			A_attack(&A, &B)
			if B.blood <= 0 {
				countA++
				break
			}
			B_attack(&A, &B)
			if A.blood <= 0 {
				countB++
				break
			}
			count++
			fmt.Printf("这是第%d次战斗", count)
			fmt.Println(",接下来他们要开始第二个回合了")
		}
	}
	if count!=0 {
		shenglv_A = countA / count
		shenglv_B = countB / count
	}
	if shenglv_A>shenglv_B {
		fmt.Printf("A的胜率更高，是%f",shenglv_A)
	} else{
		fmt.Printf("B的胜率更高，是%f",shenglv_B)
	}
}

