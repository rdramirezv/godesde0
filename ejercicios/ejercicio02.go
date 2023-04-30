package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error
var tablatexto string

func TabladeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un número para calcular su tabla de multiplicar : ")
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El dato ingresado es incorrecto, " + err.Error())
				continue
			} else {
				break
			}
		}
	}
	fmt.Println("-------------------------------------")
	fmt.Println("* Imprimiento tabla del número ", numero1, " *")
	fmt.Println("-------------------------------------")

	for i := 0; i <= 10; i++ {
		fmt.Println(numero1, "x", i, "=", numero1*i) // Resultado por pantalla
		//tablatexto += fmt.Sprintln(numero1, "x", i, "=", numero1*i)
		tablatexto += fmt.Sprintf("%d x %d = %d \n", numero1, i, numero1*i)
	}

	return tablatexto
}
