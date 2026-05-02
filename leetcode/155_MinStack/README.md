# LeetCode 155 - Min Stack 🏔️

## 📝 Descrição do Desafio

O objetivo é projetar uma pilha que suporte as operações clássicas (`push`, `pop`, `top`) e também uma operação extra (`getMin`) que retorna o menor elemento da pilha em **tempo constante O(1)**.

---

## 💡 Como resolvi

Para atingir o tempo constante de busca do valor mínimo sem percorrer a pilha, utilizei uma abordagem de **Nós Encadeados (Linked List)** onde cada nó é responsável por carregar o seu próprio estado de "mínimo histórico".

### A Estratégia:

- **Estrutura de Nó**: Cada nó possui um campo `Min` que guarda o menor valor encontrado na pilha até o momento em que aquele nó foi criado.
- **Lógica do Push**: Ao inserir um novo valor, o nó compara o seu próprio valor com o `Min` do nó anterior (o antigo topo).
- **Persistência**: O menor entre eles é salvo no campo `Min` do novo nó.
- **Resultado**: O método `getMin` apenas retorna o valor pré-calculado que está no topo atual, eliminando a necessidade de qualquer loop ou busca, garantindo complexidade $O(1)$.

---

## 🚀 Performance

- **Runtime**: 0 ms (Beats 100.00% of users with Go).
- **Complexidade de Tempo**: $O(1)$ para todas as operações.

---

### 🔗 Link do Desafio

[LeetCode 155 - Min Stack](https://leetcode.com/problems/min-stack/)
