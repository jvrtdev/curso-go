package main

import "fmt"

func main() {
	//aritméticos
	soma := 1 + 1
	subtracao := 2 - 1
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//atribuicao

	var variavel1 string = "string"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)

	//operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 != 2)

	//operadores logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // E
	fmt.Println(verdadeiro || falso) // OU
	fmt.Println(!verdadeiro)         //inverte os valores
	fmt.Println(!falso)              //inverte os valores

	//operadores unarios
	numero := 10
	numero++
	numero += 15 //numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20

	numero *= 3  // numero = numero * 3
	numero /= 10 //numero = numero / 10
	numero %= 3  //numero = numero % 3

}
