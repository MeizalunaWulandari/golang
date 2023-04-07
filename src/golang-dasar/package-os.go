package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args
	fmt.Println("Argument")
	fmt.Println(args)
	
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", name)
	}else{
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)
}
/**
 * Package OS berisikan fungsinalitas untuk mengakses fitur sistem operasi secara independen
 * artisan tidak bergantung pada jenis sistem operasi tertentu dan dapat digunakan di semua sistem operasi
 * Saat mengambil argument pastikan untuk mengambil dari index 1 karena pada index 0 berisi lokasi file
 * Getenv biasanya digunakan untuk mengambil variable enveronment
 */