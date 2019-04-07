package main

import "fmt"

func Sumar(num1 int , num2 int) int {
	return num1 + num2
}

func Restar(num1 int , num2 int) int {
	return num1 - num2
}

func Dividir(num1 int , num2 int) int {
	return num1/num2
}

func Multiplicar(num1 int , num2 int) int {
	return num1*num2
}


func main() {

	numero1 := 1 
	numero2 := 2

	fmt.Println(Sumar(numero1, numero2))
	fmt.Println(Restar(numero1, numero2))
	fmt.Println(Dividir(numero1, numero2))
	fmt.Println(Multiplicar(numero1, numero2))
}