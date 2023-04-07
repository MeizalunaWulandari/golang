package main

import "fmt"

type Biodata struct{
	Name, Address string
	Age 		  int
	Married 	  bool
}

 // Pemanggilan dengan cara biasa
func sayHello(bio Biodata, name string){
	fmt.Println("Hello", name, "My name is", bio.Name)
}


// Struct method

func (bio Biodata ) sayHello(name string){
	fmt.Println("Hello", name, "My name is", bio.Name)
}

func main(){
	var luna Biodata
	luna.Name = "Meizaluna Wulandari"
	luna.Address = "Banjarmasin"
	luna.Age = 16
	luna.Married = false
	sayHello(luna, "Andini")

	// Pemanggilan struct method
	luna.sayHello("Andini")
}

/**
 * Struct sama dingan type data lainnya, struct juga bisa gunakan sebagai paramter
 */