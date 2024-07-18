package main

import (
	"fmt"
	"net/http"
)

func meuservidor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Seja bem-vindo ao meu servidor HTTP com Go")
}

func Servidor() {
	fmt.Print("Servidor rodando na porta 3000")

	roteador := http.NewServeMux()

	roteador.HandleFunc("/servidor/meu", meuservidor)

	erro := http.ListenAndServe(":3000", roteador)

	if erro != nil {
		fmt.Println("Erro ao acessar o servidor: ", erro)
	}
}

func main() {
	Servidor()
}
