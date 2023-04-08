package main

import (
    "fmt"
    "container/ring"
    "strconv"
)

func main(){
    var data *ring.Ring = ring.New(5)

    for i := 0; i < data.Len(); i++ {
        data.Value = "Data " + strconv.FormatInt(int64(i), 10)
        data = data.Next()
    }
    data.Do(func(value interface{}){
        fmt.Println(value)
    })
}
/**
 * Package Container/ring adalah implementasi dari struktur data circural list
 * circular list adalah struktur data ring, dimana akhir elemen akan kembali keawal
 */