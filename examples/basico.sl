:: Exemplo básico da linguagem Solara
:: Demonstra variáveis, operações e estruturas condicionais

numero idade = 25
texto nome = "João"
booleano ativo = verdadeiro

imprimir "=== EXEMPLO BÁSICO SOLARA ==="
imprimir "Nome: "
imprimir nome
imprimir "Idade: "
imprimir idade
imprimir "Status ativo: "
imprimir ativo

:: Operações matemáticas
numero a = 10
numero b = 3

imprimir "=== OPERAÇÕES MATEMÁTICAS ==="
imprimir "Soma: "
imprimir a + b
imprimir "Subtração: "
imprimir a - b
imprimir "Multiplicação: "
imprimir a * b
imprimir "Divisão: "
imprimir a / b

:: Estruturas condicionais
imprimir "=== VERIFICAÇÕES ==="

se idade >= 18
    imprimir "É maior de idade"
fim

se idade < 18
    imprimir "É menor de idade"
fim

se nome == "João"
    imprimir "Olá, João!"
fim

imprimir "=== FIM DO EXEMPLO ==="
