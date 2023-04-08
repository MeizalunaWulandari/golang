package main

import (
    "fmt"
    "reflect"
)

type Sample struct{
    Name, Address string `required:"true" max:"10"`
}

func IsValid(data interface{})bool{
    t := reflect.TypeOf(data)
    for i := 0; i < t.NumField(); i++ {
        field :=  t.Field(i)
        if field.Tag.Get("required") == "true" {
            value := reflect.ValueOf(data).Field(i).Interface()
            if value == ""{
                return false
            }
        }
    }
    return true
}

func main(){
    sample := Sample{"Luna", "Banjarmasin"}
    sampleType :=  reflect.TypeOf(sample)

    fmt.Println(sampleType.NumField())
    fmt.Println(sampleType.Field(1).Name)

    fmt.Println(sampleType.Field(1).Tag.Get("required"))
    fmt.Println(sampleType.Field(1).Tag.Get("max"))

    fmt.Println(IsValid(sample))


}
/**
 * Package reflect singkatan dari reflection
 * reflection yaitu kemampuan membaca struktur kode saat apilikasinya berjalan
 */