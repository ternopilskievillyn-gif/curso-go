package main
import "fmt"

// for 
// func main() {
// 		var numero int 
// 		fmt.Print("Digite um número: ")
// 		fmt.Scan(&numero)
// 		for i := 1; i <= 10; i++ {
// 			fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
// 		}
// 	}

	// while teoricamente é uma outra forma de usar o for
func main(){
	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)
	i := 1
	for i <= 10 {
		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
		i++
	}
}