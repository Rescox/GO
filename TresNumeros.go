package main

import ("fmt")

func main() { 
	var num[100] int
	max := 0
	min := 0
	var tamano int
	fmt.Println("Introduzca los tres numeros")
	fmt.Scanln(&tamano)

	fmt.Println("Pon un numero")
	fmt.Scanln(&num[0])
	min = num[0]
	for i := 1; i < tamano; i++ {
		fmt.Scanln(&num[i])
		if num[i] > max { 
			max = num[i]
		}
		if min > num[i] { 
			min = num[i]
		}
	}
	
	fmt.Println("El numero mas grande es :%d y el minimo: %d", max, min )
}