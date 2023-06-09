package main

import "fmt"

type Blacklist func(string)bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You are blocked", name)
	}else{
		fmt.Println("Welcome back", name)
	}
}

func main(){
	blacklist := func(name string)bool{
		return name == "admin"
	}
	registerUser("Luna", blacklist)
	registerUser("admin", blacklist)
	registerUser("admin", func(name string)bool{
		return name == "admin"
	})
}

/*

*/