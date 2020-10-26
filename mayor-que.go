package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	fmt.Println("Digite la edad de Juan")
	Fscanner := bufio.NewScanner(os.Stdin)
	Fscanner.Scan()
	fmt.Println("Digite la edad de Maria")
	Sscanner := bufio.NewScanner(os.Stdin)
	Sscanner.Scan()
	fNumber := Fscanner.Text()
	sNumber := Sscanner.Text()

	if fNumber > sNumber{
		fmt.Printf("Juan es mayor que Maria con %v años", fNumber)
	}else {
		fmt.Printf("Maria es mayor que Juan con %v años", sNumber)
	}

}