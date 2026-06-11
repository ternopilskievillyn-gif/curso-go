# Curso Go

Este repositório contém meus estudos iniciais sobre a linguagem Go.
Aqui estou praticando conceitos básicos como variáveis, conversões de tipo, operadores, condicionais e estruturas de repetição.

## Conteúdos estudados

* Hello World
* Variáveis
* Conversões de tipo
* Operadores aritméticos
* Condicionais
* Estruturas de repetição

---

## Hello World

Primeiro exemplo básico em Go:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

---

## Variáveis na linguagem Go

Em Go, podemos declarar variáveis de diferentes formas. Uma das formas é utilizando a palavra-chave `var`.

Também é possível declarar mais de uma variável usando apenas um bloco `var`, como no exemplo abaixo:

```go
var (
    nome      string = "Evillyn"
    sobrenome string = "Naiane"
)
```

Também podemos declarar variáveis informando o tipo do dado:

```go
var idade int = 21
var altura float64 = 1.73
var ativo bool = true
```

Principais tipos utilizados:

* `string`: textos
* `int`: números inteiros
* `float64`: números com casas decimais
* `bool`: valores verdadeiro ou falso

Exemplo:

```go
package main

import "fmt"

func main() {
    var nome string = "Evillyn"
    var idade int = 21
    var altura float64 = 1.73

    fmt.Println(nome)
    fmt.Println(idade)
    fmt.Println(altura)
}
```

---

## Conversões de tipo

Em Go, quando queremos transformar um valor de um tipo para outro, precisamos fazer a conversão de forma explícita.

Exemplo:

```go
package main

import "fmt"

func main() {
    var numeroInteiro int = 10
    var numeroDecimal float64 = float64(numeroInteiro)

    fmt.Println(numeroInteiro)
    fmt.Println(numeroDecimal)
}
```

---

## Operadores aritméticos

Os operadores aritméticos são usados para realizar cálculos matemáticos.

Principais operadores:

* `+` soma
* `-` subtração
* `*` multiplicação
* `/` divisão
* `%` resto da divisão

Exemplo:

```go
package main

import "fmt"

func main() {
    a := 10
    b := 3

    fmt.Println("Soma:", a+b)
    fmt.Println("Subtração:", a-b)
    fmt.Println("Multiplicação:", a*b)
    fmt.Println("Divisão:", a/b)
    fmt.Println("Resto:", a%b)
}
```

---

## Condicionais

As condicionais permitem executar diferentes blocos de código dependendo de uma condição.

Exemplo com `if` e `else`:

```go
package main

import "fmt"

func main() {
    idade := 18

    if idade >= 18 {
        fmt.Println("Maior de idade")
    } else {
        fmt.Println("Menor de idade")
    }
}
```

---

## Estruturas de repetição

Em Go, usamos o `for` para criar repetições.

Exemplo:

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Número:", i)
    }
}
```

---

## Organização do repositório

```text
curso-go/
├── condicionais/
├── conversoes_tipo/
├── hello_word/
├── operadores_aritmeticos/
├── repeticao/
└── README.md
```

---

## Objetivo

O objetivo deste repositório é registrar minha evolução no aprendizado da linguagem Go, praticando desde os conceitos básicos até estruturas mais avançadas.
