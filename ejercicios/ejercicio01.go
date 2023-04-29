package ejercicios

import (
	"fmt"
	"strconv"
)

func DosValores(numtexto string) (int, string) {
	var mensaje string

	// numtexto, _ := strconv.Atoi(numero)
	numero, errordata := strconv.Atoi(numtexto)

	if errordata != nil {
		fmt.Println("Error", errordata.Error())
		// return 0, "Hubo un error" + err.Error()
	}

	if numero > 100 {
		mensaje = "Es mayor a 100"
	} else {
		mensaje = "Es menor o igual a 100"
	}

	return numero, mensaje
}
