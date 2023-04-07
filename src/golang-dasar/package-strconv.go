package main

import (
    "fmt"
    "strconv"
)

func main(){
    boolean, err := strconv.ParseBool("true")
    if err == nil {
        fmt.Println(boolean)
    }else{
        fmt.Println("Error :", err.Error())
    }

    number, err := strconv.ParseInt("1000000", 10, 64)
    if err == nil {
        fmt.Println(number)
    }else{
        fmt.Println("Error :", err.Error())
    }

    value := strconv.FormatInt(10000, 10)
    fmt.Println(value)

}
/**
 * Package strconv itu singkatan dari String Convertion
 * berfungsi untuk merubah suatu type data menjadi string
 * function pada package strconv biasanya diawali dengan kata Parse dan Format
 * function yang berawalan Parse digunakan untuk merubah string menjadi type data lain ParseTypeData()
 * sedangkan function yang berawalan Format digunakan untuk merubah type data lain menjadi string
 */