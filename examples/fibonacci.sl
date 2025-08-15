:: Exemplo: Sequência de Fibonacci

função fibonacci(n)
    se n <= 1
        retorne n
    senão
        retorne fibonacci(n - 1) + fibonacci(n - 2)
    fim
fim

imprimir "Sequência de Fibonacci:"

para numero i = 0 até 10
    numero resultado = fibonacci(i)
    imprimir "fibonacci(" + i + ") = " + resultado
fim
