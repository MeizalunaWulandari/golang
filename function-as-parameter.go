package main

import "fmt"

type Filter func(string)string

// func sayHelloWithFilter(name string, filter func(string)string){
func sayHelloWithFilter(name string, filter Filter){
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string)string{
	if name == "anjing" {
		return "..."
	}else{
		return name
	}
}

func main(){
	sayHelloWithFilter("Luna", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
	
}

/*
	Function tidak hanya bisa disimpan disebuah variable, melainkan juga bisa gunakan sebagai parameter dari function lain
	jika function terlalu panjang untuk mempersingkat dapat menggunakan function type declaration 

*/