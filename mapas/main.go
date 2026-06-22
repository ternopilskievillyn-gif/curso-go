package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" 
)

func main() {
	mapaCursos := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	curso := ""

	for curso != "q" {
		fmt.Print("Digite um curso ou digite q para sair: ")
		scanner.Scan()
		curso = scanner.Text()

		if curso != "q" {
			fmt.Print("Digite a carga horária do curso: ")
			scanner.Scan() 
			cargaStr := scanner.Text()
			
			cargaHoraria, _ := strconv.Atoi(cargaStr) 
			
			mapaCursos[curso] = cargaHoraria
		}
	}
	
	fmt.Println("\nCursos registrados:")
	for nomeCurso, cargaHoraria := range mapaCursos {
		fmt.Printf("Curso: %s, Carga Horária: %d\n", nomeCurso, cargaHoraria)
	}
}