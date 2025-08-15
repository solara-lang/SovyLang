:: Calculadora de Índice de Massa Corporal
numero peso = 70
numero altura = 1.75

numero imc = peso / (altura * altura)

imprimir "=== CALCULADORA DE IMC ==="
imprimir "Peso: "
imprimir peso
imprimir "Altura: "
imprimir altura
imprimir "IMC: "
imprimir imc

se imc < 18.5
    imprimir "Classificação: Abaixo do peso"
fim

se imc >= 18.5 e imc < 25
    imprimir "Classificação: Peso normal"
fim

se imc >= 25 e imc < 30
    imprimir "Classificação: Sobrepeso"
fim

se imc >= 30
    imprimir "Classificação: Obesidade"
fim
