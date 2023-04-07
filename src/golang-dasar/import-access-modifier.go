package main

import (
	"golang-dasar/helper"
	"fmt"
)

func main(){
	helper.Accessible("Luna") 
	// helper.inaccessible("Luna") // akan terjadi undefined karena function ini tidak bisa diakses disini
	fmt.Println(helper.Application)

}