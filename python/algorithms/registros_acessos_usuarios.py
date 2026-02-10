"""

    ## Descrição
        Você recebe uma lista de registros de acesso de usuários.

    Cada registro é um dicionário com:
        - "user_id" (int)
        - "timestamp" (string no formato "YYYY-MM-DD HH:MM:SS")

    Seu desafio é encontrar o **user_id que fez o primeiro acesso do dia**.

    ### Regras
        - Considere apenas acessos do mesmo dia.
        - A lista pode estar desordenada.
        - Se a lista estiver vazia, retorne None.
        - Se houver empate de horário, retorne o menor user_id.

    ### Exemplos de entrada

    #### Cenário 1
        acessos = [
            {"user_id": 3, "timestamp": "2026-01-30 09:15:00"},
            {"user_id": 1, "timestamp": "2026-01-30 08:00:00"},
            {"user_id": 2, "timestamp": "2026-01-30 08:00:00"},
            {"user_id": 4, "timestamp": "2026-01-30 10:30:00"}
        ]

    #### Cenário 2
        acessos = []


    #### Cenário 3
        acessos = [
            {"user_id": 5, "timestamp": "2026-01-30 11:00:00"},
            {"user_id": 2, "timestamp": "2026-01-30 07:45:00"},
            {"user_id": 9, "timestamp": "2026-01-30 09:10:00"},
        ]

    #### Cenário 4
        acessos = [
            {"user_id": 4, "timestamp": "2026-01-30 12:30:00"},
            {"user_id": 1, "timestamp": "2026-01-30 06:00:00"},
            {"user_id": 6, "timestamp": "2026-01-30 09:45:00"},
        ]


    ### Exemplos de saída
    #### Cenário 1
        1

    #### Cenário 2
        None

    #### Cenário 3:
        2

    #### Cenário 4:
        1

"""

def user_access_records(history_access):
    
    if not history_access:
        return None, None

    date = "2026-01-30"
    time = "23:59:59"

    first_access_date = f"{date} {time}"
    first_access_id = 0
    for dicts in history_access:
        for key, value in dicts.items():
            if key == "timestamp":
                if value < first_access_date:
                    first_access_date = value

                if value == first_access_date:
                    break

    finding_id = filter(lambda f: f["timestamp"] == first_access_date, history_access)
    first_access_id = list(finding_id)[0].get("user_id")

    return first_access_date, first_access_id
                

if __name__ == "__main__":

    history_access_one = [
        {"user_id": 3, "timestamp": "2026-01-30 09:15:00"},
        {"user_id": 1, "timestamp": "2026-01-30 08:00:00"},
        {"user_id": 2, "timestamp": "2026-01-30 08:00:00"},
        {"user_id": 4, "timestamp": "2026-01-30 10:30:00"}
    ]

    history_access_two = [
        {"user_id": 5, "timestamp": "2026-01-30 11:00:00"},
        {"user_id": 2, "timestamp": "2026-01-30 07:45:00"},
        {"user_id": 9, "timestamp": "2026-01-30 09:10:00"},
    ]

    history_access_three = [
        {"user_id": 4, "timestamp": "2026-01-30 12:30:00"},
        {"user_id": 1, "timestamp": "2026-01-30 06:00:00"},
        {"user_id": 6, "timestamp": "2026-01-30 09:45:00"},
    ]



    x, y = user_access_records(history_access_one)
    n, d = user_access_records([])
    v, n = user_access_records(history_access_two)
    i, u = user_access_records(history_access_three)

    print(f"{y}")
    print(f"{d}")
    print(f"{n}")
    print(f"{u}")
