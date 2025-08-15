:: Sistema de avaliação acadêmica
numero nota_matematica = 87
numero nota_portugues = 92
numero nota_ciencias = 78

numero media = (nota_matematica + nota_portugues + nota_ciencias) / 3

imprimir "=== BOLETIM ESCOLAR ==="
imprimir "Matemática: "
imprimir nota_matematica
imprimir "Português: "
imprimir nota_portugues
imprimir "Ciências: "
imprimir nota_ciencias
imprimir "Média: "
imprimir media

se media >= 90
    imprimir "Classificação: EXCELENTE"
fim

se media >= 80 e media < 90
    imprimir "Classificação: MUITO BOM"
fim

se media >= 70 e media < 80
    imprimir "Classificação: BOM"
fim

se media >= 60 e media < 70
    imprimir "Classificação: REGULAR"
fim

se media < 60
    imprimir "Classificação: INSUFICIENTE"
fim
