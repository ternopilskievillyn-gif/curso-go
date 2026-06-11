package main 
import ("fmt")

func main() {
	var amigos [5]string
	fmt.Println("Digite o nome de 5 amigos:")
	for i := 0; i < len(amigos); i++ {
		fmt.Printf("Amigo %d: ", i+1)
		fmt.Scan(&amigos[i])
	}
	fmt.Println("Seus amigos são:")
	for i := 0; i < len(amigos); i++ {
		fmt.Printf("%d. %s\n", i+1, amigos[i])
	}
}