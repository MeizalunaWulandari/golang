package helper

import "fmt"

func SayHello(name string){
    fmt.Println("Hello", name)
}

/**
 * Package adalah tempat untuk mengorganisir kode program yang dibuat dengan di Go-Lang
 * Dengan menggunakan package kita bisa merapikan kode program yang buat
 * package sebenarnya hanya directory pada sistem operasi kita
 * untuk membuat suatu file golang berada dalam sebuah package cukup dengan menggunakan key `package`
 * nama function atau apapun itu yang berada pada package yang sama harus unik
 * Penamaan function pada package yang akan di ekport ke package lain harus dimulai dengan Uppercase
*/