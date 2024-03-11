package main

import "fmt"

func somar(num1 int8, num2 int8) int8 {
	return num1 + num2
}
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 30)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)
	
	//ignorar um retorno de funcao

	resultadodasoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadodasoma)
}
