package main

import "fmt"

func getFullName() (firstName, lastName string) {
	firstName = "Meizaluna"
	lastName = "Wulandari"
	return
}

func main(){
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}

/*
	
*/