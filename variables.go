package main 

import "fmt"

func main () {

	// se puede declarar variables indicando VAR y el tipo

	var numero int
	numero=3
	fmt.Println(numero)

	// usando := no hace falta que aclare el TIPO de variable, una vez que se declaran asi, no se le puede cambiar el valor mas adelante

	numero2:=9
	nombre:="Miguel"

	fmt.Println(numero2)
	fmt.Println(nombre)

	// declarar variables en una misma linea

	edad, ciudad := 23, "Avellaneda"
	fmt.Println(edad, ciudad)

}