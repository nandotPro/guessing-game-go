package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Jogo de Adivinhação")

	answer := rand.Intn(100)
	tentativas := 0

	fmt.Println("Escolha um número entre 0 e 100")

	for {
		var guess int
		fmt.Scan(&guess)
		tentativas++

		if guess == answer {
			fmt.Printf("Parabéns! Você acertou o número em %d tentativas!\n", tentativas)
			break
		} else if guess < answer {
			fmt.Println("O número é MAIOR que isso. Tente novamente!")
		} else {
			fmt.Println("O número é MENOR que isso. Tente novamente!")
		}
	}
}
