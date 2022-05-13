package main

import "fmt"

func operacoes() {

	var a, b int
	fmt.Println("Digite o primeiro número: ")
	fmt.Scan(&a)
	fmt.Println("Digite o segundo número: ")
	fmt.Scan(&b)
	fmt.Println(a, "+", b, "=", a+b)
	fmt.Println(a, "-", b, "=", a-b)
	fmt.Println(a, "*", b, "=", a*b)
	fmt.Println(a, "/", b, "=", a/b)

}
