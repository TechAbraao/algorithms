"""
## Descrição
    Você recebe uma lista de tuplas representando partidas de um jogo.
    Cada tupla contém (nome_do_jogador, pontos_na_partida).

    Seu desafio é calcular o ranking final dos jogadores com base na soma total de pontos.

### Regras
    - Some todos os pontos de cada jogador.
    - Ordene do maior para o menor total de pontos.
    - Se houver empate na pontuação total:
        - O jogador que tiver a maior pontuação em uma única partida fica na frente.
    - Se ainda houver empate:
        - Mantenha a ordem alfabética pelo nome.
    - Se a lista estiver vazia, retorne uma lista vazia.

### Exemplos de entrada
    partidas = [
        ("Ana", 30),
        ("Bruno", 50),
        ("Ana", 20),
        ("Carlos", 40),
        ("Bruno", 10),
        ("Carlos", 15),
    ]

    partidas = []

    partidas = [
        ("Lucas", 40),
        ("Marcos", 40),
        ("Lucas", 10),
        ("Marcos", 10),
    ]

### Exemplos de saída
    [
        "1º lugar - Bruno: 60 pontos",
        "2º lugar - Carlos: 55 pontos",
        "3º lugar - Ana: 50 pontos"
    ]

    []

    [
        "1º lugar - Lucas: 50 pontos",
        "2º lugar - Marcos: 50 pontos"
    ]
"""

from typing import Tuple

def ranking(players: list[Tuple[str, int]]) -> None:
    
    for p in players:
        print(p)


if __name__ == "__main__":

    partidas = [
        ("Ana", 30),
        ("Bruno", 50),
        ("Ana", 20),
        ("Carlos", 40),
        ("Bruno", 10),
        ("Carlos", 15),
    ]

    ranking(partidas)
    