package main

import "fmt"

// Manual tanpa recursive function
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

func factorialRecursive(value int)int{
	if value == 1{
		return 1
	}else{
		return value * factorialRecursive(value-1)
	}
}

func main(){
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(factorialRecursive(5))

}

/*
	Recursive function yaitu function yang memanggil dirinya sendiri
	untuk membuat recursive function pastikan ada kondisi berhenti tuntuk mencegah terjadinya infinte loop
*/