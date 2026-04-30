// i need a function that can transform a string into a number. What ways of achiving this do you know?
// note: Dont worry, all inputs will be be strings, and very string is a perfect valid representation of an integral number.
package main
package kata

import (
	"fmt"
	"bufio"
	"string"
	"strconv"
	"os"
)

func main() {
	scanner := bufio.Newscanner(os.Stdin)
	fmt.Print("Ingresa un número: ")
	scanner.Scan()
	input := scanner.Text()
	//Metodo 1: usando strconv.Atoi
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error al convertir el número: ", err)
	} else {
		fmt.Println("Número convertido a string: ", num)
	}

	//Método 2: Usando fmt.sscanf
	var num2 int
	_, err2 := fmt.sscanf(input, "%d", &num2)
	if err2 != nil {
		fmt.Println("Error al convertir el número: ", err2)
	} else {
		fmt.Println("Numero convertido a string: ", num2)
	}
}