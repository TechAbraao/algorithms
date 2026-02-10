"""
    ## Descrição
    Você recebe uma lista de inteiros.
    Seu desafio é encontrar o primeiro número que não se repete na lista.

    ### Objetivos
    Exemplo um de entrada: numeros = [4, 5, 1, 2, 1, 4, 5, 3]
    Exemplo dois de entrada: numeros = [1, 2, 3, 2, 1]
    Exemplo três de entrada: numeros = [7, 7, 8, 8, 9, 9]

    Exemplo um de saída: 2
    Exemplo dois de saída: 3
    Exemplo três de saída: None
"""

def number_not_repeat(list_not_repeat: list) -> int | None:
    not_repeat = 0

    index = 0
    for n in list_not_repeat:
        current = n
        
        list_removed_current = list_not_repeat.copy()
        list_removed_current.remove(current)

        if current in list_removed_current:
            list_removed_current.insert(index, current)
            index += 1
        else:
            not_repeat = current
            break

        not_repeat = None

    return not_repeat


if __name__ == "__main__":
    value_one: any = number_not_repeat([4, 5, 1, 2, 1, 4, 5, 3])
    value_two: any = number_not_repeat([1, 2, 3, 2, 1])
    value_three: any = number_not_repeat([7, 7, 8, 8, 9, 9])
    print(value_one)
    print(value_two)
    print(value_three)