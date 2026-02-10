
"""

## Descrição
    Você recebe uma lista de inteiros.
    Seu desafio é contar quantas vezes cada número aparece na lista e retornar essa informação em uma estrutura de dados adequada.

### Objetivos
    Exemplo de entrada: numeros = [1, 2, 2, 3, 1, 4, 2, 3]

    Exemplo de saída: 
        {
            1: 2,
            2: 3,
            3: 2,
            4: 1
        }

"""

def frequency_count_one(list_to_count: list):
    structure = {}

    for i in list_to_count:
        current_number = i
        rep = 0

        if current_number in structure:
            continue

        current_number = i
        for j in list_to_count:
            if current_number == j:
                rep += 1
        
        structure[current_number] = rep

    return structure
    
def frequency_count_two(list_to_count: list):
    structure = {}

    for number in list_to_count:
        if number in structure:
            structure[number] += 1
        else:
            structure[number] = 1

    return structure


if __name__== "__main__":
    struc_one = frequency_count_one([1, 2, 2, 3, 1, 4, 2, 3])
    struc_two = frequency_count_two([1, 2, 2, 3, 1, 4, 2, 3])
    print(struc_one)
    print(struc_two)