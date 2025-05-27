package main

import (
	"fmt"
	"strings"
)

var tempKelvin = 273.15

func main() {
	for {
		var opcao string
		clearScreen()
		header()
		fmt.Print("Digite a unidade de origem: ")
		fmt.Scanln(&opcao)
		opcao = strings.ToLower(opcao)

		if opcao == "s" {
			clearScreen()
			fmt.Printf("\nObrigado por usar o Conversor de Escalas Termométricas!\n")
			line("*")
			fmt.Println("Saindo do programa.")
			fmt.Println("\n")
			break
		}

		switch opcao {
		case "celsius", "c":
			var celsius float64
			fmt.Print("Digite a temperatura em Celsius: ")
			fmt.Scanln(&celsius)
			fahrenheit := (celsius * 9 / 5) + 32
			kelvin := celsius + tempKelvin
			fmt.Printf("%.2f°C = %.2f°F\n", celsius, fahrenheit)
			fmt.Printf("%.2f°C = %.2fK\n", celsius, kelvin)
			pause()
		case "fahrenheit", "f":
			var fahrenheit float64
			fmt.Print("Digite a temperatura em Fahrenheit: ")
			fmt.Scanln(&fahrenheit)
			celsius := (fahrenheit - 32) * 5 / 9
			kelvin := celsius + tempKelvin
			fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, celsius)
			fmt.Printf("%.2f°F = %.2fK\n", fahrenheit, kelvin)
			pause()
		case "kelvin", "k":
			var kelvin float64
			fmt.Print("Digite a temperatura em Kelvin: ")
			fmt.Scanln(&kelvin)
			celsius := kelvin - tempKelvin
			fahrenheit := (celsius * 9 / 5) + 32
			fmt.Printf("%.2fK = %.2f°C\n", kelvin, celsius)
			fmt.Printf("%.2fK = %.2f°F\n", kelvin, fahrenheit)
			pause()
		default:
			line("")
			fmt.Println("Unidade inválida. Tente Novamente ;)")
			line("")
		}
	}
}

func line(caracter string) {
	if len(caracter) == 0 {
		caracter = "-"
	}
	fmt.Println(strings.Repeat(caracter, 76))
}

func header() {
	line("=")
	fmt.Println(" Conversor de Escalas Termométricas")
	line("-")
	fmt.Println("|   (c) Celsius   |  (f) Fahrenheit  |    (k) Kelvin    |     (s) Sair     |")
	line("=")
}

func pause() {
	fmt.Printf("\nObrigado por usar o Conversor de Escalas Termométricas!\n")
	line("-")
	fmt.Println("\nPressione Enter para continuar...")
	fmt.Scanln()
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
