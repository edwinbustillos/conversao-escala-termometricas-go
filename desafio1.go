package main

import (
	"fmt"
	"strings"
)

func main() {
	var opcao string
	fmt.Print("Digite a unidade de origem (Celsius = c, Fahrenheit = f ou Kelvin = k): ")
	fmt.Scanln(&opcao)
	opcao = strings.ToLower(opcao)

	switch opcao {
	case "celsius", "c":
		var celsius float64
		fmt.Print("Digite a temperatura em Celsius: ")
		fmt.Scanln(&celsius)
		fahrenheit := (celsius * 9 / 5) + 32
		kelvin := celsius + 273.15
		fmt.Printf("%.2f°C = %.2f°F\n", celsius, fahrenheit)
		fmt.Printf("%.2f°C = %.2fK\n", celsius, kelvin)
	case "fahrenheit", "f":
		var fahrenheit float64
		fmt.Print("Digite a temperatura em Fahrenheit: ")
		fmt.Scanln(&fahrenheit)
		celsius := (fahrenheit - 32) * 5 / 9
		kelvin := celsius + 273.15
		fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, celsius)
		fmt.Printf("%.2f°F = %.2fK\n", fahrenheit, kelvin)
	case "kelvin", "k":
		var kelvin float64
		fmt.Print("Digite a temperatura em Kelvin: ")
		fmt.Scanln(&kelvin)
		celsius := kelvin - 273.15
		fahrenheit := (celsius * 9 / 5) + 32
		fmt.Printf("%.2fK = %.2f°C\n", kelvin, celsius)
		fmt.Printf("%.2fK = %.2f°F\n", kelvin, fahrenheit)
	default:
		fmt.Println("Unidade inválida. Use (c) para Celsius, (f) para Fahrenheit ou (k) para Kelvin.")
	}
}
