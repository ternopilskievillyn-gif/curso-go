package main
import ("fmt")

// if e else

// func main(){
// 	var idade int
// 	fmt.Print("Digite a sua idade: ")
// 	fmt.Scanf("%d", &idade)
// 	if idade >= 18 {
// 		fmt.Println("Você é maior de idade.")
// 	} else{
// 		fmt.Println("Você é menor de idade.")}
// }


// switch case

func main(){
	var mes int
	fmt.Print("Digite o número do mês (1-12): ")
	fmt.Scanf("%d", &mes)
	switch mes {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fevereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Maio")
	case 6:	
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	default:
		fmt.Print("Valor inválido. Digite um número entre 1 e 7.")

}
	switch mes {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("O mês tem 31 dias.")
	case 4, 6, 9, 11:
		fmt.Println("O mês tem 30 dias.")
	case 2:
		fmt.Println("O mês tem 28 ou 29 dias.")
	default:
		fmt.Print("Valor inválido. Digite um número entre 1 e 12.")
	}
}
