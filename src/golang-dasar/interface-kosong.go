package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	}else if i == 2 {
		return true
	}else {
		return "Ups"
	}
}

func main(){
	var data = Ups(3)
	fmt.Println(data)
}

/**
 * Interface kosong adalah interface yang tidak memiliki deklarasi method satupun
 * hal ini membuat semua secara otomatis semua type data akan menjadi implementasinya
*/