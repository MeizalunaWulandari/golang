package main

import "golang-dasar/helper"

func main(){
    helper.SayHello("Luna")
}

/**
 * golang hanya hanya dapat mengakses file golang yang berada pada package yang sama
 * untuk itu jika ingin menggunakan package lain di package yg berbeda kita harus melakukan import
 * untuk melakukan import cukup dengan `import ("dirname/packageName")`
 */