package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "David"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"rua dos bobos", 0}

	u2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{nome: "Davi"}
	fmt.Println(u3)
}
