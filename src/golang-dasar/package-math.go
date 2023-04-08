package main

import (
    "fmt"
    "math"
)

func main(){
    fmt.Println(math.Round(1.7)) //Untuk membulatkan float kebilangan terdekat
    fmt.Println(math.Round(1.3)) //Untuk membulatkan float kebilangan terdekat

    fmt.Println(math.Floor(1.7)) //Untuk membulatkan ke bawah
    fmt.Println(math.Ceil(1.7)) //Untuk membulatkan ke Atas

    fmt.Println(math.Max(10, 20)) //Untuk Mencari nilai Tertinggi
    fmt.Println(math.Min(10, 20)) //Untuk Mencari nilai terendah
}
/**
 * Package math berisikan constant dan fungsi matematika
 */