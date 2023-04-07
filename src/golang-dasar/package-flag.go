package main

import(
	"fmt"
	"flag"
)

func main(){
	var host = flag.String("host","localhost","Put your database host")
	var user = flag.String("user","user","Put your database user")
	var pass = flag.String("pass","root","Put your database password")

	flag.Parse()
	fmt.Println("Host :", *host)
	fmt.Println("User :", *user)
	fmt.Println("Passwords :", *pass)
}
/**
 * Package flag erat kaitannyan dengan package os
 * untuk membuat flag cukup dengan
 * flag.String("flagName","defaultValue","helper or description")
 * untuk mengunakannya pada terminal cukup dengan
 * go run golang.go -flagName="newValue"
 * flag akan menghasilkan pointer jadi untuk menampilkannya gunakan operator *
 * fmt.Println(*varibaleName)
 * untuk parsing flag gunakan function flag.Parse()
 */