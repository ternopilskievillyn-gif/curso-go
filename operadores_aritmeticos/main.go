package main
import ("fmt")

//Operadores Aritméticos
// +, -, *, /, % (resto da divisão), ++ (incremento), -- (decremento)

func main(){

	var numero1  int
	var numero2  int
	fmt.Print("Digite o primeiro número:")
	fmt.Scanln(&numero1)
	fmt.Print("Digite o primeiro número:")
	fmt.Scanln(&numero2)

	fmt.Printf("%d + %d = %d \n",numero1, numero2, numero1 + numero2)
	fmt.Printf("%d - %d = %d \n",numero1, numero2, numero1 - numero2)
	fmt.Printf("%d * %d = %d \n",numero1, numero2, numero1 * numero2)
	fmt.Printf("%d / %d = %d \n",numero1, numero2, numero1 / numero2)
	fmt.Printf("%d %% %d = %d \n",numero1, numero2, numero1 % numero2)
	incremento:= numero1
	incremento += numero2
	fmt.Printf("O incremento de %d com %d é %d",numero1, numero2, incremento )
	decremento:= numero1
	decremento -= numero2
	fmt.Printf("O decremento de %d com %d é %d",numero1, numero2, decremento )
	fmt.Print("Finalizados")
}