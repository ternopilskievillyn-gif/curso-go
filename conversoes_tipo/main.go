package main 
import ("fmt"
"strconv"
)

func main (){
	var texto string
	fmt.Print("Digite um número:")
	fmt.Scanf("%s", &texto)
	var numero int 
	//o _ simboliza que está ignorando o erro
	numero, _ = strconv.Atoi(texto)
	fmt.Printf("O número digitado foi: %d", numero)

}