package main 

import "fmt"

func main() {
	var numero1, numero2 int
	var operacao string

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&numero1) 

	fmt.Print("Digite a operação (+, -, *, /, $): ")
	fmt.Scan(&operacao)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&numero2) 

	if operacao == "+" {
		fmt.Println("Resultado:", somar(numero1, numero2))
	} else if operacao == "-" {
		fmt.Println("Resultado:", subtrair(numero1, numero2))
	} else if operacao == "*" {
		fmt.Println("Resultado:", multiplicar(numero1, numero2))
	} else if operacao == "/" {
		if numero2 == 0 {
			fmt.Println("Não é possível dividir por zero!")
		} else {
			fmt.Println("Resultado:", dividir(numero1, numero2))
		}
	}
} 

func somar(numero1 int, numero2 int) int {
	return numero1 + numero2
}

func subtrair(numero1 int, numero2 int) int {
	return numero1 - numero2
}

func multiplicar(numero1 int, numero2 int) (resultado int) {
	resultado = numero1 * numero2
	return
}

func dividir(numero1 int, numero2 int) float64 {
	return float64(numero1) / float64(numero2)
}