package main

import (
    "fmt"
    "container/list"
)

func main(){
    data := list.New()


    data.PushBack("Puspa")
    data.PushBack("Meizaluna")
    data.PushBack("Wulandari")
    data.PushBack("Andini")

    fmt.Println(data.Front().Value) // Mengambil data paling depan
    fmt.Println(data.Front().Next().Value) // Mengambil Selanjutnya
    fmt.Println(data.Back().Value)  // Mengambil data paling belakang

    // Dari depan kebelakang
    for element := data.Front(); element != nil; element = element.Next(){
        fmt.Println(element.Value)
    }

    // Dari belakang ke depan
    for element2 := data.Back(); element2 != nil; element2 = element2.Prev(){
        fmt.Println(element2.Value)
    }

}
/**
 * Package Container/list adalah implementasi dari struktur data double linked list di golang 
 */