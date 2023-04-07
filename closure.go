package main

import "fmt"

func main(){
	name := "Luna"
	counter := 0

	increment := func(){
		name := "Meizaluna"
		fmt.Println("Increment")
		fmt.Println(name)
		counter++
	}
	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
/*
	Closure yaitu kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
	Harap berhati-hati saat mengguunakan closure, karena dapat merubah variable diatasnya,
	closure dapat mengakses variable diatasnya namun variable yang dibuatnya hanya dapat digunakan pada function tersebut
*/