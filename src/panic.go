package main

import "fmt"

func endApp(){
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main(){
	runApp(true)
}

/*
	panic adalah yang bisa digunakan untuk menggentikan sebuah program
	panic biasanya dipanggil ketika error saat program kita berjalan
	saat function panic dipanggil, program akan berhenti namun defer tetap dieksekusi
*/