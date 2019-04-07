package main

import (
	"fmt"
	"strconv"
)

func main () {  

	edad := "2"

	//aca convierto un STRING en un INT
	// Atoi devuelve dos valores
	edad_int, err := strconv.Atoi(edad)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(edad_int+10)

	edad2 :="18"
	fmt.Println("la edad es: ",edad2)

}