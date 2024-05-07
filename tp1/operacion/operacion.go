package operacion

import (
	"fmt"
	"math"
	"strconv"
	TDAPila "tdas/pila"
)

const raizCuadrada = "sqrt"
const suma = "+"
const resta = "-"
const multiplicacion = "*"
const division = "/"
const potencia = "^"
const logaritmo = "log"
const operadorTernario = "?"
const mensajeError = "ERROR"

const operandosMinimos = 2
const operandosMinimosSqrt = 1
const operandosMinimosOperadorTernario = 3
const resultadosMaximos = 1

func CalculadoraPolacaInversa(operaciones []string) (int64, error) {
	operandos := TDAPila.CrearPilaDinamica[int64]()
	cantidadOperandos := 0
	var resultado int64

	if len(operaciones) == 1 {
		return 0, fmt.Errorf(mensajeError)
	}

	for _, simbolo := range operaciones {
		if esOperador(simbolo) { // Al ser un operador, desapila y hace las cuentas
			if err := verificarOperandosMinimos(simbolo, cantidadOperandos); err != nil {
				return 0, err
			}

			if simbolo == raizCuadrada {
				op := operandos.Desapilar()
				if op < 0 {
					return 0, fmt.Errorf(mensajeError)
				}
				resultado = RaizCuadrada(op)

			} else {
				op2 := operandos.Desapilar()
				op1 := operandos.Desapilar()

				switch simbolo {
				case suma:
					resultado = Suma(op1, op2)
				case resta:
					resultado = Resta(op1, op2)
				case multiplicacion:
					resultado = Multiplicacion(op1, op2)
				case division:
					if op2 == 0 {
						return 0, fmt.Errorf(mensajeError)
					}
					resultado = Division(op1, op2)
				case potencia:
					if op2 < 0 {
						return 0, fmt.Errorf(mensajeError)
					}
					resultado = Potencia(op1, op2)
				case logaritmo:
					if op2 < 2 {
						return 0, fmt.Errorf(mensajeError)
					}
					resultado = Logaritmo(op1, op2)
				case operadorTernario:
					op3 := operandos.Desapilar()
					resultado = OperadorTernario(op1, op2, op3)
					cantidadOperandos--
				}
				cantidadOperandos--
			}
			operandos.Apilar(resultado)
		} else { // Al ser un número, lo convierte a int64 y lo apila
			numero, err := strconv.ParseInt(simbolo, 10, 64)
			if err != nil {
				return 0, fmt.Errorf(mensajeError)
			}
			operandos.Apilar(numero)
			cantidadOperandos++
		}
	}
	if cantidadOperandos != resultadosMaximos {
		return 0, fmt.Errorf(mensajeError)
	}
	return operandos.Desapilar(), nil
}

func verificarOperandosMinimos(simbolo string, cantidadOperandos int) error {
	switch simbolo {
	case suma, resta, multiplicacion, division, potencia, logaritmo:
		if cantidadOperandos < operandosMinimos {
			return fmt.Errorf(mensajeError)
		}
	case raizCuadrada:
		if cantidadOperandos < operandosMinimosSqrt {
			return fmt.Errorf(mensajeError)
		}
	case operadorTernario:
		if cantidadOperandos < operandosMinimosOperadorTernario {
			return fmt.Errorf(mensajeError)
		}
	}
	return nil
}

func RaizCuadrada(op int64) int64 { return int64(math.Sqrt(float64(op))) }

func Suma(op1, op2 int64) int64 { return op1 + op2 }

func Resta(op1, op2 int64) int64 { return op1 - op2 }

func Multiplicacion(op1, op2 int64) int64 { return op1 * op2 }

func Division(op1, op2 int64) int64 { return op1 / op2 }

func Potencia(op1, op2 int64) int64 { return int64(math.Pow(float64(op1), float64(op2))) }

func Logaritmo(op1, op2 int64) int64 { return int64(math.Log(float64(op1)) / math.Log(float64(op2))) }

func OperadorTernario(op1, op2, op3 int64) int64 {
	if op3 != 0 {
		return op1
	}
	return op2
}

func esOperador(simbolo string) bool {
	switch simbolo {
	case raizCuadrada:
		return true
	case suma:
		return true
	case resta:
		return true
	case multiplicacion:
		return true
	case division:
		return true
	case potencia:
		return true
	case logaritmo:
		return true
	case operadorTernario:
		return true
	default:
		return false
	}
}
