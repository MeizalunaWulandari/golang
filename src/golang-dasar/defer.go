package main

import "fmt"

func logging(){
	fmt.Println("Selesai memanggil Function")
}

func runApplication(value int){
	defer logging()
	fmt.Println("Application is runnng")
	result := 10 / value
	fmt.Println("Result", result)
}

func main(){
	runApplication(0)
}

/*
	defer adalah function yang bisa dijadwalkan setelah suatu function selesai di eksekusi
	defer selalu dieksekusi walau ada error sekalipun
	biasakan menggunakan defer diatas, agar bisa dieksekusi walau program dibawahnya error
*/