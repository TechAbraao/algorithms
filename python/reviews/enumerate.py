def func_enumerate(list_names):
    for index, value in enumerate(list_names):
        print(f"Índice: {index} - Valor: {value}")

if __name__ == "__main__":
    func_enumerate(["Abraão", "José", "Messias"])