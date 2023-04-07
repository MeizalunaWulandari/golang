package main

import (
    "fmt"
    "strings"
)

func main(){
    fmt.Println(strings.Contains("Meizaluna Wulandari", "luna")) // harusnya true
    fmt.Println(strings.Contains("Meizaluna Wulandari", "Andini")) // harusnya false

    fmt.Println(strings.Split("Meizaluna Wulandari", " ")) // untuk memisahkan menjadi slice dengan separator
    fmt.Println(strings.ToLower("Meizaluna Wulandari")) // Untuk merubah semua huruf menjadi lowercase
    fmt.Println(strings.ToUpper("Meizaluna Wulandari")) // Untuk merubah semua huruf menjadi uppercase

    fmt.Println(strings.Trim("    Meizaluna Wulandari    ", " ")) // untuk menghilangkan spasi dikiri dan dikanan
    fmt.Println(strings.ReplaceAll("luna Wulandari luna", "luna", "Andini")) // untuk menganti semua kata dengan kata lainnya
}
/**
 * Package string adalah package yang berisi function-function untuk memanipulasi type data string
 * fmt.Println(strings.Contains("Meizaluna Wulandari", "luna"))
 * untuk cek apakah ada string luna pada string Meizaluna Wulandari. Harusnya true
 */