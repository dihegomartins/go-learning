package main // SEMPRE main para arquivos executáveis

import "fmt"

func main() {
    var idade int = 27
    var preco float64 = 49.49
    estaLogado := true
    mensagem := "Aprendendo Go"
    
    fmt.Printf("Tipo do preço: %T\n", preco)
    fmt.Printf("Dados: %v, %v, %v, %v\n", idade, preco, estaLogado, mensagem)
}