NAMES = ["Abraão", "Santos", "Ana", "Carlos", "Fernanda", "Ricardo", "Juliana", "Marcos", "Patrícia", "Eduardo", "Camila", "Rafael", "Beatriz", "Lucas", "Mariana", "Felipe", "Larissa", "Gustavo", "Vanessa", "Thiago", "Letícia", "Bruno", "Isabela", "Diego", "Aline", "Rodrigo", "Priscila", "Leandro", "Tatiana", "Fábio", "Renata", "Anderson", "Cláudia", "Sandro", "Michele", "Wagner", "Simone", "Cleber", "Daiane", "Murilo", "Viviane", "Henrique", "Natália", "Alexandre", "Débora", "Willian", "Cristiane", "Matheus", "Érica", "Vinícius"]


# 1. Qual tamanho do nome? mostrar [tamanho do nome - nome]
def loops_names(names):
    # Varrer cada nome na lista de nomes
    for name in names:
        length = len(name)
        start_with = name[0]
        message = f"[LENGTH: {length} | NAME: {name}] | START WITH: {start_with}"
        
        print(message)
        

if __name__ == "__main__":
    loops_names(NAMES)