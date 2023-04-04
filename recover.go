package main

import "fmt"

func endApp(){
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message: ", message)
	}
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	
	fmt.Println("Aplikasi Berjalan ")
}

func main(){
	runApp(false)
	fmt.Println("Luna")
}

/*
	recover adalah function yang dapat digunakan untuk menangkap data panic
	dengan recover proses panic akan berhenti sehingga program akan tetap berjalan
	recover yang benar disimpan di function defer
*/