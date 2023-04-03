package main

import "fmt"

func sumAll(number ...int) int {
	total := 0

	for _, value := range number {
		total += value
	}
	return total
}

func main(){
	total := sumAll(20,25)
	fmt.Println(total)

	slice := []int{20,30}
	total = sumAll(slice...)
	fmt.Println(total)


}

/*
	Parameter yang berada paling kanan(diakhir) memiliki kemampuan untuk dijadikan varargs
	varargs artinya datanya bisa memiliki lebih dari satu input, atau anggap saja semacam array
	jika telah memiliki variable slice tetap bisa dimasukkan ke dalam variadic function dengan cara mengambahkan ... pada akhir
*/