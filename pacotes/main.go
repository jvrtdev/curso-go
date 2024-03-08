package main

import (
	"fmt"
	"modulo/auxiliar"
)
//aqui importamos um package que criamos dentro de uma pasta

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
}
//do pacote auxiliar, chamamos a função Escrever()