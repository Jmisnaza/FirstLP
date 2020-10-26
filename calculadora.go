package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	fmt.Println("Digite el primer numero")
	Fscanner := bufio.NewScanner(os.Stdin)
	Fscanner.Scan()
	fmt.Println("Digite un operador aritmetico como (+, -, /, *)")
	OAritmetico := bufio.NewScanner(os.Stdin)
	OAritmetico.Scan()
	fmt.Println("Digite el segundo numero")
	Sscanner := bufio.NewScanner(os.Stdin)
	Sscanner.Scan()
	firstNumber := Fscanner.Text()
	opArit := OAritmetico.Text()
	secondNumber := Sscanner.Text()

	fn, err1 := strconv.Atoi(firstNumber)
	sn, err2 := strconv.Atoi(secondNumber)
	if err1 != nil || err2 != nil{
		
		fmt.Println("Los datos ingresados son incorrectos intentelo de nuevo")
	
	}else {
		switch oa := opArit; oa {
		case "+":
			result := (fn + sn)
			fmt.Println("El resultado es ", result)
		case "-":
			result := (fn - sn)
			fmt.Println("El resultado es ", result)
		case "*":
			result := (fn * sn)
			fmt.Println("El resultado es ", result)
		case "/":
			result := (fn / sn)
			fmt.Println("El resultado es ", result)
		default:
			fmt.Println("Esta opcion no concuerda con ningun operador aritmetico")
		}
	}
}