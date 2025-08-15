:: Exemplo que testará o sistema de bibliotecas
:: Este código deve falhar se a biblioteca smath não estiver carregada
:: inclua a biblioteca responsável por calculos matematicos!!

sovy smath include

numero a = 10
numero b = 3

imprimir "Testando operações matemáticas:"
imprimir "A = "
imprimir a
imprimir "B = "
imprimir b

:: Esta operação deve dar erro pedindo para instalar smath
numero resultado = a + b
imprimir "Resultado da divisão:"
imprimir resultado

numero resto = a % b
imprimir "Resto da divisão:"
imprimir resto
