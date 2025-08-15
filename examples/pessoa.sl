:: Exemplo: Trabalhando com mapas (objetos)

mapa pessoa = {
    nome: "Maria",
    idade: 30,
    ativo: verdadeiro
}

imprimir "Dados da pessoa:"
imprimir "Nome: " + pessoa["nome"]
imprimir "Idade: " + pessoa["idade"]

se pessoa["ativo"]
    imprimir pessoa["nome"] + " está ativa no sistema"
senão
    imprimir pessoa["nome"] + " não está ativa"
fim

:: Função que trabalha com mapas
função descreverPessoa(p)
    imprimir "Descrição:"
    imprimir "- Nome: " + p["nome"]
    imprimir "- Idade: " + p["idade"]
    
    se p["idade"] >= 18
        imprimir "- Status: Adulto"
    senão
        imprimir "- Status: Menor"
    fim
fim

descreverPessoa(pessoa)
