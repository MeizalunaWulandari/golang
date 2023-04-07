package main

import "fmt"

type HasName interface{
	getName() string
}

func sayHello(name HasName){
	fmt.Println("Hello", name.getName())
}

// IMPLEMENTASI INTERFACE
type Person struct{
	Name string
}

func (person Person) getName()string{
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName()string{
	return animal.Name	
}

func main(){
	luna := Person{
		Name : "Meizaluna Wulandari",
	}
	sayHello(luna)

	kucing := Animal{
		Name : "Kucing",
	}
	sayHello(kucing)
}


/**
 * Interface yang dimaksud bukan UI interface
 * Interface adalah type data yang abstrak dia tidak memiliki emplimentasi secara langsung
 * biasanya interface digunakan sebagai kontrak
 * ==== IMPLENTASI INTERFACE ===
 * Setiap type data yang sesuai dengan interface, secara otomatis dianggap sebagai interface tersebut
 * 
 */