package main

import (
    "fmt"
    "regexp"
)

func main(){
    var regex *regexp.Regexp = regexp.MustCompile(`l[a-z]a`)
    fmt.Println(regex.MatchString("meizaluna"))
    fmt.Println(regex.MatchString("meiza"))
    fmt.Println(regex.MatchString("Meiza"))
    fmt.Println(regex.MatchString("lua"))
    fmt.Println(regex.MatchString("eko"))
    fmt.Println(regex.MatchString("beka"))

    search := regex.FindAllString("luna eko eka meiza meizaluna rizka miska", 1)
    fmt.Println(search)
}
/**
 * Package regexp adalah utilitas untuk melakukan pencarian regular expression
 */