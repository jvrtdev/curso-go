package main

import (
	"fmt"
)

func main(){
	
	var array1[5] string
	array1[0] = "Posicao 1"
	fmt.Println(array1)

	array2 := [5]string{
		"posicao1", "posicao2",
		"posicao3", "posicao4", "posicao5"}
	fmt.Println(array2)
	
	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)

	slice := []int{10,11,12,13,14,15,16,17}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)


	slice2 := array2[1:3]
	fmt.Println(slice2)

	// arrays internos 
	fmt.Println("- - - - - - - - - ")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println("Tamanho",len(slice3))
	fmt.Println("Capacidade",cap(slice3))
	
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}