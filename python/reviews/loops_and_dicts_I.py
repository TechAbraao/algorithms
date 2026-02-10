

def first_account_access(list_of_access: list[dict]) -> int:
    if list_of_access is None:
            print(" [!] Lista está vazia. ")
            return None

    # Maior data, utilizada para comparação
    first_data = None
    # ID relacionado, para posteriormente expor
    first_id = None

    # Loop para iterar na lista de dicionários
    for access in list_of_access:

        # Nessa lógica, ele pega na primeira iteração os dados, isto é, o timestamp do primeiro dicionário
        # torna-se o objeto de comparação, pois assume que ele é o maior.
        if first_data is None:
            first_id = access.get("user_id", "N/A")
            first_data = access.get("timestamp", "N/A")
            continue

        # Agora, se a data da próxima iteração for menor que a data salva como objeto de comparação
        # o algoritmo deverá trocar, afim de buscar a menor data possível.
        if access.get("timestamp", "N/A") < first_data:
            first_id = access.get("user_id", "N/A")
            first_data = access.get("timestamp", "N/A")
    
    print(f" [!] User ID com primeiro acesso: {first_id}")
    return first_id


if __name__ == "__main__":
    list_one = [
            {"user_id": 3, "timestamp": "2026-01-30 09:15:00"},
            {"user_id": 1, "timestamp": "2026-01-30 08:00:00"},
            {"user_id": 2, "timestamp": "2026-01-30 08:00:00"},
            {"user_id": 4, "timestamp": "2026-01-30 10:30:00"}
        ]
    list_two = []
    list_three = [
            {"user_id": 5, "timestamp": "2026-01-30 11:00:00"},
            {"user_id": 2, "timestamp": "2026-01-30 07:45:00"},
            {"user_id": 9, "timestamp": "2026-01-30 09:10:00"},
        ]
    list_four = [
            {"user_id": 4, "timestamp": "2026-01-30 12:30:00"},
            {"user_id": 1, "timestamp": "2026-01-30 06:00:00"},
            {"user_id": 6, "timestamp": "2026-01-30 09:45:00"},
        ]

    first_account_access(list_one)
    first_account_access(list_two)
    first_account_access(list_three)
    first_account_access(list_four)