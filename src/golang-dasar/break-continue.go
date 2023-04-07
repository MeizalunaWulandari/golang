package main

import "fmt"

func main(){
	for i := 0; i < 10; i++ {

		if i == 5{
			break
		}

		fmt.Println("Perulangan ke",i)
	}

	for i := 0; i < 10; i++ {

		if i%2 == 0{
			continue
		}

		fmt.Println("Perulangan ke",i)
	}

}

/*
	Break & Continue adalah kata kunci yang bisa digunakan dalam perulangan
	Break digunakan untuk menghentikan seluruh perulangan
	continue digunakan untuk menghentikan perulangan yang berjalan, dan melanjutkan perulangan berikutnya
*/