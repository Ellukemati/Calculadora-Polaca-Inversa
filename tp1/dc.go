package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	Calculadora "tp1/operacion"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linea := scanner.Text()
		linea = strings.TrimSpace(linea)
		if linea == "" {
			continue
		}

		simbolos := strings.Fields(linea)
		resultado, err := Calculadora.CalculadoraPolacaInversa(simbolos)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(resultado)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error leyendo la entrada estándar:", err)
		os.Exit(1)
	}
}
