package main

import (
    "fmt"
    "time"
)

func main(){
    now := time.Now()

    fmt.Println(now)
    fmt.Println("Tahun : ",now.Year())
    fmt.Println("Bulan : ",now.Month())
    fmt.Println("Tanggal : ",now.Day())
    fmt.Println("Jam : ",now.Hour())
    fmt.Println("Menit : ",now.Minute())
    fmt.Println("Detik : ",now.Second())
    fmt.Println("Milidetik : ",now.Nanosecond())

    utc := time.Date(2023, 12, 31, 12, 59, 10,11, time.UTC)
                    //(tahun, bulan, tanggal, jam, menit, detik, milidetik, lokasi) 
    fmt.Println(utc)

    // layout := time.RFC3339
    layout := "2006-01-02" //Standar golang, bukan tanggal yang dibuat sendiri
    parse, _ := time.Parse(layout, "2023-04-08")
    fmt.Println(parse)
}
/**
 * Package Time adalah package yang berisikan function untuk management waktu
 */