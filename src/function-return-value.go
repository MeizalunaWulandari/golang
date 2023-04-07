package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello bro"
	}else{
		return "Hello " + name
	}
}

func main(){
	result := getHello("Luna")
	fmt.Println(result)
	fmt.Println(getHello(""))
}

/*
	return berfunngsi untuk mengmbalikan data
	untuk mengembalikan data dengan return pembuatan function harus disertai dengan typeData func namaFunction(namaParamter typeData) type Data {}
	setelah return kode program dibawahnya tidak akan dieksekusi
*/