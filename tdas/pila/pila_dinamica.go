package pila

/* Definición del struct pila proporcionado por la cátedra. */

const largoMinimo = 10
const factorRedimension = 2

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, largoMinimo),
		cantidad: 0,
	}
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == cap(pila.datos) {
		redimensionar(pila, cap(pila.datos)*factorRedimension) // Al llenarse extiende su capacidad
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	if pila.cantidad > largoMinimo && pila.cantidad*factorRedimension*factorRedimension <= cap(pila.datos) {
		redimensionar(pila, cap(pila.datos)/factorRedimension) // Al sobrar mucho espacio acorta su capacidad
	}
	elemento := pila.datos[pila.cantidad-1]
	pila.cantidad--
	return elemento
}

func redimensionar[T any](pila *pilaDinamica[T], nuevaCapacidad int) {
	nuevoArreglo := make([]T, nuevaCapacidad)
	copy(nuevoArreglo, pila.datos)
	pila.datos = nuevoArreglo
}
