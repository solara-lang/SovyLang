:: Exemplo usando sistema de bibliotecas
:: Primeiro importa a biblioteca smath

sovy smath include

numero a = 16
numero b = 4

imprimir "=== EXEMPLO COM BIBLIOTECA SMATH ==="
imprimir "A = "
imprimir a
imprimir "B = "
imprimir b

:: Operações básicas que agora funcionam com smath
numero divisao = a / b
imprimir "Divisão:"
imprimir divisao

numero resto = a % b
imprimir "Resto:"
imprimir resto

:: Usando funções avançadas da biblioteca smath
numero potencia_resultado = potencia(a, 2)
imprimir "A ao quadrado:"
imprimir potencia_resultado

numero raiz_resultado = raiz(a)
imprimir "Raiz quadrada de A:"
imprimir raiz_resultado

numero pi_valor = pi()
imprimir "Valor de PI:"
imprimir pi_valor

numero maximo = max(a, b)
imprimir "Máximo entre A e B:"
imprimir maximo

imprimir "=== TESTE CONCLUÍDO ==="
