package auxiliar

import "fmt"

func escrever2() {
	fmt.Println("escrevendo pela funcao escrever 2")
}
//funções declaradas com letra minuscula só podem ser importadas dentro dos seus pacotes
//se tentarmos usar a func escrever2 no arquivo main.go, não vai ser possível
//apenas funções com letra maiuscula permitem ser usadas fora de seu pacote