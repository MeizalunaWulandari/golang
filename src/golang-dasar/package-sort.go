package main
import (
    "fmt"
    "sort"
)

type User struct {
    Name string
    Age int
}

type UserSlice []User

func (value UserSlice) Len() int {
    return len(value)
}
func (value UserSlice) Less(i, j int) bool {
    return value[i].Age < value[j].Age
}
func (value UserSlice) Swap(i, j int){
    value[i], value[j] = value[j], value[i]
}

func main(){
    users :=  []User{
        {"Luna", 16},
        {"Andini", 36},
        {"Abdi", 12},
        {"Rizka", 16},
        {"Eko", 30},
    }
    sort.Sort(UserSlice(users))
    fmt.Println(users)
}
/**
 * Package sort adalah utilitas atau helper untuk proses pengurutan data
 * agar data bisa diurutkan, kita harus mengimplementasikan kontrak di interface.
 * sort.Interface
 */