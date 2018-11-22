package main

import ("fmt")
func main() {
	fmt.Println("Introduzca un numero")
	var i int
	fmt.Scanf("%d", &i)
	if i % 2 == 0 { 
		fmt.Print("Es par")
	} else { 
		fmt.Print("Es impar")
	}
}