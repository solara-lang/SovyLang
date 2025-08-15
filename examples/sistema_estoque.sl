:: Sistema simples de estoque
funcao verificar_estoque(produto, quantidade)
    imprimir "=== CONTROLE DE ESTOQUE ==="
    imprimir "Produto: "
    imprimir produto
    imprimir "Quantidade: "
    imprimir quantidade
    
    se quantidade > 50
        imprimir "Status: ESTOQUE ALTO"
    fim
    
    se quantidade >= 10 e quantidade <= 50
        imprimir "Status: ESTOQUE NORMAL"
    fim
    
    se quantidade < 10 e quantidade > 0
        imprimir "Status: ESTOQUE BAIXO - REABASTECER"
    fim
    
    se quantidade == 0
        imprimir "Status: SEM ESTOQUE - URGENTE"
    fim
fim

:: Testando o sistema
verificar_estoque("Notebook Dell", 5)
verificar_estoque("Mouse Wireless", 25)
verificar_estoque("Teclado Mecânico", 0)
