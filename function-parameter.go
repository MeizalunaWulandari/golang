package main

import "fmt"

func sayHelloTo(firstName string, lastName string){
	fmt.Println("Halo", firstName, lastName)
}

func main(){
	sayHelloTo("Meizaluna", "Wulandari")
	sayHelloTo("Andani", "Wulandari")
}

/*
	Parameter digunakan ketika membutuhkan data dari luar function tersebut
	parameter bersifat unique pada function tersebut
	membuat parameter cukup dengan func namaFantction(namaParameter typeData)
	untuk membuat berberapa parameter dalam satu function cukup pisahkan dengan koma
*/