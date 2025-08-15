:: Exemplo: Trabalhando com listas (arrays)

lista frutas = ["maçã", "banana", "laranja"]
lista numeros = [1, 2, 3, 4, 5]

imprimir "Frutas:"
imprimir frutas

imprimir "Primeira fruta: " + frutas[0]
imprimir "Última fruta: " + frutas[2]

imprimir "Números:"
para numero i = 0 até tamanho(numeros) - 1
    imprimir "numeros[" + i + "] = " + numeros[i]
fim

:: Adicionar elemento
lista novaLista = adicionar(frutas, "uva")
imprimir "Lista com uva:"
imprimir novaLista
