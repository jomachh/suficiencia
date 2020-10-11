def main():
    print("Bienvenido.")
    input("Presiona 'Enter' para continuar...")

    x = int(input("Ingresa un número por favor.\n"))
    y = int(input("Ingresa otro número por favor.\n"))

    if x%3 == 0 and y%3 == 0:
        triangleArea = (x * y) / 2
        print(f'El área del triángulo es: {triangleArea}')

        repeat = int(input("¿Desea realizar otra operación? (1 para continuar, 0 para salir)\n"))

        if repeat == 1:
            main()
    else:
        arithmeticOperations(x, y)

def arithmeticOperations (x, y):
    print("\nSeleciona el número de la operación aritmética a efectuar.")
    print("1. Suma")
    print("2. Resta")
    print("3. Multiplicación")
    print("4. División")
    print("5. Residuo")

    operation = int(input(""))

    operations(operation, x, y)
    print("\n\nGracias por utilizar nuestra herramienta, esperamos verte de nuevo pronto.")


def sum (x, y):
    return print(f'\nEl resultado de la suma es: {x + y}')

def sub (x, y):
    return print(f'\nEl resultado de la resta es: {x - y}')

def mul (x, y):
    return print(f'\nEl resultado de la multiplicación es: {x * y}')

def div (x, y):
    return print(f'\nEl resultado de la división es: {round(x / y, 2)}')

def res (x, y):
    return print(f'\nEl residuo de la división es: {x % y}')

def operations(arg, x, y):
    switcher = {
        1: sum,
        2: sub,
        3: mul,
        4: div,
        5: res
    }

    if arg > 5:
        print('Ingresa una opción válida.\n')
        operations(arg, x, y)
    else:
        func = switcher.get(arg, lambda: 'Ingresa una opción válida.')

    print (func(x, y))

if __name__ == "__main__":
    main()