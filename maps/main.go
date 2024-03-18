package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string] string{
		"nome": {
			"primeiro": "Joao",
			"ultimo": "Pedro",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["carro"] = map[string]string{
		"marca": "chevrolet",
	}
	fmt.Println(usuario2)

}