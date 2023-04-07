package helper

import "fmt"

var Application = "Belajar Golang"
var version = "1.0"

func inaccessible(name string){
    fmt.Println("Hello", name)
    // function ini tidak bisa diakses dari package lain karena diawali dengan huruf kecil
    
}

func Accessible(name string){
    fmt.Println("Hello", name)
    // function ini bisa diakses dari package lain karena diawali dengan huruf besar
}

/**
 * Access Modifier yaitu sebuah penentuan apakah suatu variable atau function dapat diakses 
 * diluar function tersebut?
 * Access modifier di Go-Lang cukup sederhana yaitu hanya dengan penulisan huruf pada awal penamaannya
 * jika suatu variable atau function diawali dengan huruf kapital artinya dia bisa diakses diluar package tersebut
 * contohnya var Hello, func Hello()
 * sedangkan variable atau function yang diawali dengan huruf kecil berarti dia hanya bisa diakes dari package tersebut
 * contoh var hello, func hello()
*/