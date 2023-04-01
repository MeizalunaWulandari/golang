package main

import "fmt"

func main(){
	type NoTelpon string
	type married bool

	var noTelponLuna NoTelpon = "081294107600"
	var marriedStatus married = false
	fmt.Println(noTelponLuna)
	fmt.Println(marriedStatus)
}

/*
	Type declaration adalah kemampuan membuat type baru dari type yang sudah ada
*/