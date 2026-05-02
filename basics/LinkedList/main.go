package main

import "fmt"

func main() {
	// 1. Criamos o primeiro nó manualmente (o nosso ponto de partida)
	head := &ListNode{Val: 10}

	// 2. Agora usamos sua função para inserir o 20 depois do 10
	// O 'head' atual é o nosso "anterior"
	inserirDepois(head, 20)

	// 3. Vamos inserir o 30 depois do 20?
	// Para isso, precisamos chegar no 20 primeiro.
	// Como sabemos que o 20 é o head.Next:
	inserirDepois(head.Next, 30)

	fmt.Println("Lista após inserções:")
	percorrerLista(head) // Deve mostrar 10, 20, 30

	// 4. Vamos testar sua função de remover?
	// Vamos remover quem está DEPOIS do head (ou seja, remover o 20)
	fmt.Println("\nRemovendo o nó após o 10...")
	removerProximo(head)

	fmt.Println("Lista final:")
	percorrerLista(head) // Deve mostrar 10, 30
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func percorrerLista(head *ListNode) {
	curr := head

	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func removerProximo(noAnterior *ListNode) {
	if noAnterior == nil || noAnterior.Next == nil {
		return
	}

	noAnterior.Next = noAnterior.Next.Next
}

func inserirDepois(anterior *ListNode, valor int) {
	if anterior == nil {
		return
	}

	// 1. Criamos o novo vagão
	novoNo := &ListNode{Val: valor}

	// 2. O novo vagão segura a mão de quem o 'anterior' estava segurando
	novoNo.Next = anterior.Next

	// 3. O anterior solta a mão do antigo e segura a mão do novo
	anterior.Next = novoNo
}
