package main 

import ("fmt")

//Hello word + formas de usar o print
// func main(){
// 	fmt.Println("Hello World GO")
// 	fmt.Println("1 + 1 =", 1+1)
// 	fmt.Println("1.1 + 2.2=", 1.1 + 2.2)
// 	fmt.Println(true)
// }

// como declarar variáveis
func main(){
	fmt.Println("INTEIROS SEM SINAL")
	var numero uint8 = 255; // byte e uint8 são a mesma coisa
	var numero2 uint16 = 33
	fmt.Println("O valor da variável é", numero)
	fmt.Println("O valor da variável é", numero2)

	// criando constantes
	const pi = 3.14
	const e = 2.71	

	const (
		pi2 = 3.14
		e2 = 2.71
	)
	fmt.Println("O valor de pi é", pi)
	fmt.Println("O valor de e é", e)
}