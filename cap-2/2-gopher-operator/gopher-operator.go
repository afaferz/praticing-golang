package main

import "fmt"

// Package Level Escope - Escopo do pacote
var z1 = "Another simple string"

func main() {
	// Tipagem automática (declaração)
	// Gopher Operator tem escopo de bloco
	x := 10
	y := true
	z := "A simple string"

	// Tipos diferentes não são comparados
	// simpleExpression := x == y

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T", z, z)

	// Atribuição
	x = 20
	fmt.Printf("x: %v, %T\n", x, x)

	// Atribuição de x e declaração de x1
	x, x1 := 15, 30
	fmt.Printf("x: %v, %T x1: %v, %T\n", x, x, x1, x1)

	simpleExpression := x == x1

	fmt.Printf("simple expression: %v, %T\n", simpleExpression, simpleExpression)

	// Variável z1 é impressa sem problemas :D
	fmt.Printf("z1: %v, %T\n", z1, z1)

	// Palavras reservadas (Keywords) - Não podem ser atribuidas à nomes de variáveis

	// break        default      func         interface    select
	// case         defer        go           map          struct
	// chan         else         goto         package      switch
	// const        fallthrough  if           range        type
	// continue     for          import       return       var
}
