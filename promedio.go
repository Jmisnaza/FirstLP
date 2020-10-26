package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	var cantNumber float64 = 2
	var suma float64
	fmt.Println("Digite el primer numero")
	Fscanner := bufio.NewScanner(os.Stdin)
	Fscanner.Scan()
	fmt.Println("Digite el segundo numero")
	Sscanner := bufio.NewScanner(os.Stdin)
	Sscanner.Scan()
	firstNumber := Fscanner.Text()
	secondNumber := Sscanner.Text()
	fn, err1 := strconv.ParseFloat(firstNumber, 64)
	fs, err2 := strconv.ParseFloat(secondNumber, 64)
	if err1 != nil || err2 != nil{
		fmt.Println("Los datos ingresados son incorrectos intentelo de nuevo")
	} else {
		suma = fn + fs
		promedio := (suma / cantNumber)
		fmt.Printf("El promedio entre %v y %v es %v", fn, fs, promedio)
	}
	

}