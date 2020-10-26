package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	var cantNumber float64 = 3
	fmt.Println("Digite nota 1")
	Fscanner := bufio.NewScanner(os.Stdin)
	Fscanner.Scan()
	firstNumber := Fscanner.Text()
	fn, err1 := strconv.ParseFloat(firstNumber, 64)	
	fmt.Println("Digite nota 2")
	Sscanner := bufio.NewScanner(os.Stdin)
	Sscanner.Scan()
	secondNumber := Sscanner.Text()
	sn, err2 := strconv.ParseFloat(secondNumber, 64)
	fmt.Println("Digite nota 3")
	Tscanner := bufio.NewScanner(os.Stdin)
	Tscanner.Scan()
	thirdNumber := Tscanner.Text()
	tn, err3 := strconv.ParseFloat(thirdNumber, 64)
	
	if err1 != nil || err2 != nil || err3 != nil {

		fmt.Println("Los datos ingresados son incorrectos intentelo de nuevo")		

	}else {

		sumaNotas := fn+sn+tn
		promedioNotas := sumaNotas/cantNumber
		fmt.Println("El promedio de las notas es ", promedioNotas)

	}		 
}