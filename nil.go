package main

import "fmt"

func newMap(name string) map[string]string{
	if name == ""{
		return nil
	}else{
		return map[string]string{
			"name" : name,
		}
	}
}

func main(){
	person := newMap("Meizaluna")
	fmt.Println(person)
}

/**
 * Nil sama seperti null pada bahasa PHP
 * untuk melakukan pengecekan data yang defaultnya kosong disarankan menggunakan nil
 * nil sendiri hanya bisa digunakan  dibeberapa type data seperti, interface, function, map, slice, pointer dan channel
 */