package database 

import "fmt"
var connection string

func init(){
	connection = "MySQL"
    fmt.Println(connection)
}

func GetDatabase()string{
	return connection
}

/**
 * Package initialization, biasa digunakan ketika ingin menjalankan suatu function pada packade tertentu secara otomatis 
 * ketika package tersebut di import
 * untuk membuatnya cukup dengan membuat function dengan nama init
 * func init()
*/