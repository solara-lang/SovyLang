:: Demo completo da linguagem Solara
:: Mostra todos os recursos funcionando

numero x = 10
texto nome = "Solara"
booleano ativo = verdadeiro

imprimir "=== DEMO DA LINGUAGEM SOLARA ==="
imprimir ""
imprimir "1. Variáveis:"
imprimir nome
imprimir x
imprimir ativo

imprimir ""
imprimir "2. Operações matemáticas:"
numero resultado = x + 5
imprimir resultado

imprimir ""
imprimir "3. Condicionais:"
se x > 5
    imprimir "X é maior que 5"
fim

se x < 5
    imprimir "X é menor que 5"
senao
    imprimir "X não é menor que 5"
fim

imprimir ""
imprimir "4. Funções:"
funcao cumprimentar(pessoa)
    imprimir "Olá "
    imprimir pessoa
    imprimir "!"
fim

cumprimentar("Mundo")

imprimir ""
imprimir "5. Loops:"
para numero i = 1 ate 3
    imprimir "Loop: "
    imprimir i
fim

imprimir ""
imprimir "=== FIM DO DEMO ==="
