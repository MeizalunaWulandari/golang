package main

import "fmt"

func main(){
	// counter := 1

	// for counter <= 20 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }	

	for counter := 1; counter <= 20; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Luna","Andini", "Rizka"}

	for i:=0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	// For Range untuk perulangan pada array, slice dan map
	// For Range pada slice
	for _, value := range slice {
		fmt.Println(value)
	}

	// For Range pada map
	person := map[string]string{
		"name" : "Meizaluna Wulandari",
		"address" : "Banjarmasin",
		"education" : "SMK/SMA Frater don Bosco",
	}

	for key, value := range person{
		fmt.Println(key, "=" ,value)
	}

}
/*
	gunakan _ (underscore) pada for range untuk mengatasi error karena valiable i tidak digunakan
*/