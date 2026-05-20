


# Quanto maior o array, mais o for percorre até encontrar o target.
# Precisamos assumir sempre o pior cenário, ou seja, o target está na
# última posição do array.
def simple_search(arr, target):
    for item in arr:
        if item == target:
            return True
        return False

## Desafio 1
# Esse é um exemplo de busca já indexada, ou seja, independente do tamanho da lista (entrada), a busca
# trará sempre o primeiro item da lista. A complexidade é CONSTANTE O(1)
# Isso significa tmb que o numero de operações nao cresce com o tamanho dos dados
def algoritmo_1(lista):
    return lista[0]


## Isso no contexto de Sets no Python, que removem duplicata e trazem apenas 1 unico item.
def algoritmo_11(set_1, item):
    return item in set_1
    ## Retorna True se tiver, e False se não tiver.

## Desafio 2
# O truque é, se tem 1 Loop, é basicamente ctz que é de complexidade linear O(n)]
# O problema disso é que, se dobra os dados de entrada, dobra o tempo de execução.
def algoritmo_2(lista):
    for item in lista:
        print(item)
    

## Desafio 3
# SÃO dois FOR em 1 lista, isso é complexidade QUADRATICA
# Isso significa que o tempo de execução cresce com o quadrado do tamanho
# da entrada, se dobrar os dados, o tempo quadriplica.
def algoritmo_3(lista):
    for i in lista: # Loop Externo: roda n vezes
        for j in lista: # Loop Interno: roda n vezes pra cada i
            print(i, j)
    
    
    
if __name__ == "__main__":
    pass