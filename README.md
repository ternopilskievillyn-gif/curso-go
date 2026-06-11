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
