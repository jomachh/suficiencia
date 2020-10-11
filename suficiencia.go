
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var x, y int
	var repeat bool

	fmt.Println("\n\nBienvenido.")
	fmt.Println("Presiona 'Enter' para continuar...")
	bufio.NewReader(os.Stdin).ReadBytes('\n') 

	fmt.Println("Hola, ingresa un número por favor.")
	fmt.Scanf("%d", &x)
	fmt.Println("Ingresa otro número por favor.")
	fmt.Scanf("%d", &y)

	if x%3 == 0 && y%3 == 0 {
		var triangleArea = (x * y) / 2
		fmt.Printf("El área del triángulo es: %v\n\n", triangleArea)
		fmt.Println("¿Desea realizar otra operación? (1 para continuar, 0 para salir)")
		fmt.Scanf("%t", &repeat)

		if repeat { main() }

	} else {
		arithmeticOperations(x, y)
	}
}

func arithmeticOperations (x int, y int) {
	var operation int

		fmt.Println("\nSeleciona el número de la operación aritmética a efectuar.")
		fmt.Println("1. Suma")
		fmt.Println("2. Resta")
		fmt.Println("3. Multiplicación")
		fmt.Println("4. División")
		fmt.Println("5. Residuo")
		fmt.Scanf("%d", &operation)

		switch operation {
		case 1:
			result := x + y
			fmt.Printf("El resultado de la suma es: %d", result)
		case 2:
			result := x - y
			fmt.Printf("El resultado de la resta es: %d", result)
		case 3:
			result := x * y
			fmt.Printf("El resultado de la multiplicación es: %d", result)
		case 4:
			result := float64(x) / float64(y)
			fmt.Printf("El resultado de la división es: %.2f", result)
		case 5:
			result := x % y
			fmt.Printf("El residuo de la división es: %d", result)
		default:
			fmt.Println("\n\nSelecciona una opción válida")
			arithmeticOperations(x, y)
		}

		fmt.Println("\n\nGracias por utilizar nuestra herramienta, esperamos verte de nuevo pronto.")
}

