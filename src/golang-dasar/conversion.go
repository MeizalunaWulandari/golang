package main

import "fmt"

func main(){
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	var name = "Luna"
	var e = name[2]
	var eString = string(e)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}