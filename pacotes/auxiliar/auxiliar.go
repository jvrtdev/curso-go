package auxiliar

import "fmt"

func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
//funções declaradas com letra minuscula só podem ser importadas dentro dos seus pacotes
//se tentarmos usar a func escrever2 no arquivo main.go, não vai ser possível
//apenas funções com letra maiuscula permitem ser usadas fora de seu pacote