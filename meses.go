package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("Digite un numero del mes:")
	nScanner := bufio.NewScanner(os.Stdin)
	nScanner.Scan()
	numberMonth := nScanner.Text()
	switch nm := numberMonth; nm {
	case "1":
		fmt.Println("January")
	case "2":
		fmt.Println("Frebuary")
	case "3":
		fmt.Println("March")
	case "4":
		fmt.Println("April")
	case "5":
		fmt.Println("May")
	case "6":
		fmt.Println("June")
	case "7":
		fmt.Println("July")
	case "8":
		fmt.Println("August")
	case "9":
		fmt.Println("September")
	case "10":
		fmt.Println("October")
	case "11":
		fmt.Println("November")
	case "12":
		fmt.Println("December")
	default:
		fmt.Println("Esta opcion no concuerda con ningun mes")
	}
}