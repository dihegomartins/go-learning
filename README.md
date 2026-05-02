# Go Learning & Algorithms 🚀

Este repositório é dedicado ao meu processo de aprendizado da linguagem **Go (Golang)**, focado em resolver desafios de lógica e implementar estruturas de dados clássicas.

O objetivo aqui é consolidar conceitos de performance, gerenciamento de memória e algoritmos eficientes.

## 📂 Estrutura do Repositorio

- `Aprenda_GO/`: Exercícios e notas do curso de fundamentos da linguagem.
- `basics/`: Implementações de lógica básica e sintaxe.
- `beecrowd/`: Desafios resolvidos na plataforma Beecrowd (antigo URI).
- `leetcode/`: Exercícios de algoritmos do LeetCode (foco em performance 0ms).
- `hackerrank/`: Problemas de lógica e estruturas de dados do HackerRank.

## 🛠️ Estruturas de Dados Implementadas

Até o momento, o foco principal tem sido em **Pilha (Stack)**, explorando diferentes formas de implementação:

- [x] Pilha baseada em Nós (Ponteiros/Dynamic Allocation)
- [x] Pilha baseada em Slices (Foco em Performance/Cache Locality)

## 🏆 Evolução de Performance

Um dos grandes diferenciais desse estudo é a otimização de código. Abaixo, um exemplo de evolução no LeetCode:

| Problema                | Abordagem       | Runtime | Performance |
| :---------------------- | :-------------- | :-----: | :---------: |
| #1047 Remove Duplicates | Stack com Nós   |  19ms   |   Top 15%   |
| #844 Backspace Compare  | Stack com Slice |   0ms   | Top 100% 🚀 |

## 🚀 Como Executar

Cada pasta de exercício contém um arquivo `main.go`. Para rodar qualquer um deles:

```bash
cd leetcode/20_ValidParentheses
go run main.go
```
