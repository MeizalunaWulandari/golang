package main

import "fmt"

func getFullName() (string, string, string){
	return "Andini", "Puspa", "Wulandari"
}

func main(){
	firstName, _, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}

/*
	Go-Lang bisa mengembalikan multiple value
	untuk mengembalikan multiple value type data harus didefinisikan satu persatu pada function dipisahkan dengan koma func namaFunction() (typeData, typeData) {}
	semua data multiple wajib digunakan, namun jika ingin mengunakan hanya beberapa saja bisa diabalikan dengan _ (underscore)
*/