package main 

import "fmt"

func main() {
	var amigos []string 
	nome := ""
	
	for nome != "q" {
		fmt.Print("Digite o nome do amigo ou 'q' para sair: ")
		
		fmt.Scan(&nome) 
		
		if nome != "q" {
			amigos = append(amigos, nome)
		}
	}
	
	fmt.Println("Amigos:", amigos)
}