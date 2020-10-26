package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	var nMayor []int
	fmt.Println("Digite el primer numero")
	Fscanner := bufio.NewScanner(os.Stdin)
	Fscanner.Scan()
	firstNumber := Fscanner.Text()
	fn, err1 := strconv.Atoi(firstNumber)	
	nMayor = append(nMayor, fn)
	fmt.Println("Digite el segundo numero")
	Sscanner := bufio.NewScanner(os.Stdin)
	Sscanner.Scan()
	secondNumber := Sscanner.Text()
	sn, err2 := strconv.Atoi(secondNumber)
	nMayor = append(nMayor, sn)
	fmt.Println("Digite el tercer numero")
	Tscanner := bufio.NewScanner(os.Stdin)
	Tscanner.Scan()
	thirdNumber := Tscanner.Text()
	tn, err3 := strconv.Atoi(thirdNumber)
	nMayor = append(nMayor, tn)
	numeroMayor := nMayor[0]
	for _, numero := range nMayor {
		if numero > numeroMayor{
			numeroMayor = numero
		}
	}
	if err1 != nil || err2 != nil || err3 != nil {

		fmt.Println("Los datos ingresados son incorrectos intentelo de nuevo")

	}else {
		fmt.Println("El numero mayor es ", numeroMayor)
	}	
	
}