package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 10000000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias - apelido
	//int32 = rune

	var numero3 rune = 123456
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var real1 float32 = 123.45
	fmt.Println(real1)

	var real2 float64 = 123333333333.56
	fmt.Println(real2)

	var str string = "topdemais"
	fmt.Println(str)

	var boleano1 bool
	fmt.Println(boleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
