"""
    ## Descrição
        Você recebe uma lista de palavras (strings).
        Seu desafio é encontrar a **palavra mais longa** da lista.

    ### Regras
        - Se houver empate no tamanho, retorne a que aparece primeiro.
        - Se a lista estiver vazia, retorne None.
        - Ignore espaços extras no começo e no fim das palavras.

    ### Exemplos de entrada
        palavras = ["python", "java", "golang", "js"]
        palavras = ["  casa", "apartamento  ", "sol "]
        palavras = []
        palavras = ["sol", "lua", "céu"]
        palavras = ["   novidade", "              bebe", "   bom"]

    ### Exemplos de saída
        "golang"
        "apartamento"
        None
        "sol"
"""

def string_size(str_list: list) -> str | None:
    
    if str_list == []:
        return None

    previus_word = ""
    for word in str_list:
        word_strip = word.strip()

        if len(word_strip) == len(previus_word):
            previus_word = word
            break

        if len(word_strip) > len(previus_word):
            previus_word = word

    return previus_word


if __name__ == "__main__":
    one = string_size(["python", "java", "golang", "js"])
    two = string_size(["  casa", "apartamento  ", "sol "])
    three = string_size([])
    four = string_size(["sol", "lua", "céu"])
    five = string_size(["   novidade", "bebe", "bom"])
    print(one)
    print(two)
    print(three)
    print(four)
    print(five)