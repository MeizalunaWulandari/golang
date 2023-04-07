package main

import "fmt"

func random() interface{}{
	return 1
}

func main(){
	var result = random()
	// var resultString = result.(string)
	// fmt.Println(resultString)

	// Mengunakan switch expressiom
	switch value := result.(type){
	case string :
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Value", value, "unknown")
	}

}
/**
 * Type assertion merupakan kemampuan untuk mengubah type data menjadi type data yang kita inginkan
 * Fitur ini sering digunakan ketika bertemu dengan interface kosong
 * saat menggunakan assertion pastikan data yang dibuah benar 
 * agar lebih aman sebaiknya menggunakan switch expression saat menggunakan assertion 
 * dan menggunakan recover agar program tetap berjalan saat terjadi panic
*/