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
		fmt.Print("Digite seu palpite: ")

		// Verificando se a entrada é um número válido
		_, err := fmt.Scan(&guess)
		if err != nil {
			// Limpar o buffer de entrada para evitar loop infinito
			var discard string
			fmt.Scanln(&discard)

			fmt.Println("Erro: Por favor, digite apenas números!")
			continue
		}

		// Verificando se o número está no intervalo válido
		if guess < 0 || guess > 100 {
			fmt.Println("Por favor, digite um número entre 0 e 100!")
			continue
		}

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
